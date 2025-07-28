# GolangSSG

A lightweight Static Site Generator written in Go that converts Markdown files to HTML.

## Features

✅ **Currently Supported:**

- Headers (H1-H6) with `#` syntax
- Text formatting: **bold**, *italic*, ~~strikethrough~~
- `Inline code` blocks
- Unordered lists with `-`
- Ordered lists with `1.`
- Automatic paragraph wrapping

🚧 **Coming Soon:**

- Pre Compile Regex (Patterns micro optimization)
- Fenced code blocks with syntax highlighting
- Blockquotes
- Links and images
- Horizontal rules
- Front matter support
- Template system
- Multiple file processing


## Quick Start

```bash
git clone https://github.com/ECHIDNATHEG/GolangSSG.git
cd GolangSSG
go run main.go
```

## Example

**Input (Markdown):**
```markdown
# My Blog Post

This is **bold** and *italic* text.

## Todo List
- Write code
- Test features
- Deploy site
```

**Output (HTML):**
```html
<h1>My Blog Post</h1>

<p>This is <strong>bold</strong> and <em>italic</em> text.</p>

<h2>Todo List</h2>
<li>
  <ul>Write code</ul>
  <ul>Test features</ul>
  <ul>Deploy site</ul>
</li>
```

## Usage

Currently processes a single test file. Edit `TESTFILE` constant in `main.go` to point to your Markdown file:

```go
const TESTFILE = "/path/to/your/file.md"
```

## Architecture

- **Scanner-based parsing**: Processes Markdown line by line
- **Regex pattern matching**: Converts Markdown syntax to HTML
- **State tracking**: Handles multi-line elements like lists
- **Modular functions**: Separate handlers for headers, text formatting, and lists

## Contributing

This is an active learning project! Feel free to:
- Report issues
- Suggest features
- Submit pull requests
- Fork and experiment

## Roadmap

1. **Phase 1** (Current): Basic Markdown parsing
2. **Phase 2**: File system operations and multiple file support
3. **Phase 3**: Template system and themes
4. **Phase 4**: Advanced features (plugins, optimization)

## License

MIT License - feel free to use and modify!

---

*Built