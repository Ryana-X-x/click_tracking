# Session Model

## Purpose

The Session model represents the beginning of a tracking lifecycle.

Each tracking link click creates exactly one Session.
A Session may generate 20–30 Events over time.

This model is intentionally minimal to keep write performance high.

---

## Field Breakdown

### ID (UUID)

Primary key of the session.

Why UUID instead of auto-increment?

- System is horizontally scalable
- Multiple API instances generate sessions
- Avoids DB sequence contention
- No coordination required between services

UUIDs allow stateless API servers.

---

### CampaignID (UUID)

Represents the campaign responsible for this session.

Why indexed?

- Most reporting filters by campaign
- Campaign-level aggregation is common
- Reduces scan time on large datasets

---

### CreatedAt (time.Time)

Timestamp when session was created.

Why not use gorm.Model?

- Avoids unnecessary fields (UpdatedAt, DeletedAt)
- Prevents soft-delete overhead
- Gives us strict control over schema
- Tracking systems are append-only; soft deletes are unnecessary

---

## Design Philosophy

Sessions are lightweight.

They exist only to:
- Group related events
- Provide campaign attribution anchor

We avoid adding:
- User data
- Device data
- IP address
- Any heavy metadata

Why?

Because sessions must be created extremely fast.
This is a write-heavy system.

---

## What Would Break at Scale Without This Design?

If we used auto-increment IDs:
- High write contention on DB sequence
- Harder horizontal scaling

If we added unnecessary columns:
- Increased storage cost
- Slower inserts
- More index overhead

If CampaignID was not indexed:
- Campaign aggregation would become slow as data grows

---

## Production Considerations

- Sessions table is append-only.
- No updates expected after creation.
- Can later be partitioned by time if scale increases.
- Should have connection pooling tuned at DB layer.

This model is intentionally minimal to support high throughput.
