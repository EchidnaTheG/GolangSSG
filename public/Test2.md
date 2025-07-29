# Table Testing Document

This document tests various table formats and edge cases.

## Basic Table

| Name     | Age | City        |
|----------|-----|-------------|
| Alice    | 25  | New York    |
| Bob      | 30  | Los Angeles |
| Charlie  | 35  | Chicago     |

## Table with Formatting

| **Product** | *Price* | ~~Old Price~~ | `Code` |
|-------------|---------|---------------|--------|
| **iPhone**  | $999    | ~~$1099~~     | `IP14` |
| *MacBook*   | $1299   | ~~$1499~~     | `MB16` |
| iPad        | $599    | ~~$699~~      | `ID11` |

## Table with Links and Images

| Item | Image | Link | Description |
|------|-------|------|-------------|
| Logo | ![Logo](logo.png) | [Website](https://example.com) | Our company logo |
| Banner | ![Banner](banner.jpg) | [Download](files/banner.zip) | Marketing banner |
| Icon | ![Icon](icon.svg) | [Source](https://icons.com) | App icon design |

## Complex Table with Mixed Content

| Feature | **Status** | *Priority* | Notes |
|---------|------------|------------|-------|
| Authentication | ✅ **Done** | *High* | Uses `JWT` tokens |
| Database | 🔄 *In Progress* | *Medium* | ~~MySQL~~ → PostgreSQL |
| API | ❌ **Pending** | *Low* | RESTful with `JSON` |
| Testing | ✅ **Done** | *High* | `Jest` + `Cypress` |

## Table with Code Blocks

Before the table, here's some code:

```javascript
function parseTable(data) {
    return data.split('|');
}
```

| Function | Language | Purpose |
|----------|----------|---------|
| `parseTable()` | **JavaScript** | Parse table data |
| `formatCell()` | *Python* | Format cell content |
| `renderHTML()` | **Go** | Generate HTML output |

After the table, more content.

## Alignment Table (Advanced)

| Left Align | Center Align | Right Align |
|:-----------|:------------:|------------:|
| Left       | Center       | Right       |
| Data       | More Data    | End Data    |

## Table with Special Characters

| Symbol | Name | Unicode | HTML Entity |
|--------|------|---------|-------------|
| © | Copyright | U+00A9 | `&copy;` |
| ™ | Trademark | U+2122 | `&trade;` |
| ® | Registered | U+00AE | `&reg;` |
| € | Euro | U+20AC | `&euro;` |

## Minimal Table

| A | B |
|---|---|
| 1 | 2 |

## Wide Table

| Col1 | Col2 | Col3 | Col4 | Col5 | Col6 |
|------|------|------|------|------|------|
| Data1 | Data2 | Data3 | Data4 | Data5 | Data6 |
| More1 | More2 | More3 | More4 | More5 | More6 |

## Table with Empty Cells

| Name | Email | Phone | Notes |
|------|-------|-------|-------|
| John | john@example.com | 555-1234 | Regular customer |
| Jane |  | 555-5678 | New customer |
| Bob | bob@test.com |  |  |

## Mixed Content Around Tables

Here's a paragraph before the table with **bold** and *italic* text.

| Item | Quantity | Price |
|------|----------|-------|
| Apples | 10 | $5.00 |
| Bananas | 5 | $2.50 |

And here's content after the table with `inline code` and [a link](https://example.com).

## Final Test

This tests table parsing with surrounding content to ensure proper boundaries.

- List item before table
- Another list item

| Final | Test |
|-------|------|
| Last | Row |

1. Numbered list after table
2. Another numbered item

**The end!**