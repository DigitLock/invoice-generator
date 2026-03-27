# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## Project Overview

Invoice Generator is a self-hosted invoicing tool for freelancers and small businesses. It has two modes:
- **Guest mode**: Browser-only invoice form with client-side PDF generation (no backend)
- **Authorized mode**: Full CRUD with dashboard, company/client management, invoice history, and status tracking

The project shares JWT authentication with a sibling "Expense Tracker" application (shared `SECRET_KEY`).

## Architecture

### Frontend (Stage 1+3)
- **Vue.js 3** with TypeScript, **Pinia** for state management, **Tailwind CSS** for styling
- Guest invoice form reuses the same component as authorized mode
- Client-side PDF generation for guest mode (jsPDF or pdf-lib)
- Dev server runs on port 5173

### Backend (Stage 2)
- **Go** microservice with **Chi** router on port 8081
- **PostgreSQL 14+** with **sqlc** (type-safe SQL) and **pgx/v5** driver
- JWT auth via `golang-jwt` (shared SECRET_KEY with Expense Tracker)
- Server-side PDF generation (gofpdf or similar)
- Data isolation by `family_id` from JWT claims (`user_id` + `family_id`)

### Mobile (Stage 4)
- Flutter app (iOS + Android), read-only for company/bank settings

## Key Domain Concepts

- **Invoice statuses**: Draft -> Sent -> Partially Paid -> Paid | Cancelled (terminal). Backward transitions not allowed.
- **isOverdue flag**: Boolean, independent of status, applicable to all statuses except Draft. Tracks client reliability.
- **Invoice number format**: `INV-DDMMYYYY-NNN` (sequential per user per date, editable)
- **Currencies**: EUR and RSD predefined, plus custom ISO 4217 codes. One currency per invoice.
- **Monetary values**: Stored as `DECIMAL(15,2)`, transmitted as JSON strings to avoid floating-point issues.
- **Soft deletes**: Companies, clients, bank accounts, and invoices use `deleted_at` timestamps.
- **Business rule**: Invoices cannot be created for inactive clients.

## API Structure

REST API at `/api/v1/` with resources: companies, clients, bank-accounts (nested under companies), invoices. All authorized endpoints require `Authorization: Bearer <JWT>` header. PDF generation at `GET /api/v1/invoices/{id}/pdf`.

## Deployment

Self-hosted via Docker container behind Cloudflare Tunnel at `invoice.digitlock.systems`.

## Staging Plan

The project is built incrementally:
1. **Stage 1** (current): Vue.js guest form + client-side PDF (no backend)
2. **Stage 2**: Go backend + PostgreSQL + auth + CRUD APIs
3. **Stage 3**: Authorized dashboard UI in same Vue.js app
4. **Stage 4**: Flutter mobile app
