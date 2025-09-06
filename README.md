# Advanced ISBN Validation Library

This project challenges you to build a **comprehensive ISBN validation system** supporting both ISBN-10 and ISBN-13 formats with advanced features including batch processing, custom validation rules, and performance optimization.

**🕒 Target completion time: 30 minutes** *(experienced developer with 5+ years)*

---

## 📚 Background

ISBNs (International Standard Book Numbers) uniquely identify books:
- **ISBN-10**: 10-digit format (legacy, pre-2007)
- **ISBN-13**: 13-digit format (current standard, EAN-13 compatible)

Both formats use checksums to detect errors, but with different algorithms.

---

## 📖 Validation Rules

### ISBN-10 Checksum
- First 9 digits multiplied by position (1-9), sum mod 11
- Checksum `10` represented as `X`

### ISBN-13 Checksum  
- Alternating weights (1,3,1,3...) for first 12 digits
- Checksum = `(10 - (sum mod 10)) mod 10`
- Must start with `978` or `979` prefix

---

## 🎯 Requirements

### Core Features *(Must implement all)*

1. **Multi-Format Validation**
   - Support both ISBN-10 and ISBN-13
   - Auto-detect format from input
   - Proper checksum validation for each type

2. **Input Normalization** 
   - Handle various separators: hyphens, spaces, dots
   - Case-insensitive 'X' handling
   - Validate ISBN-13 prefixes (978/979 only)

3. **Batch Processing**
   - Process arrays/slices of ISBNs efficiently  
   - Return structured results with error details
   - Concurrent processing for batches >50 items

4. **Advanced Error Handling**
   - Specific error types: `InvalidFormat`, `InvalidChecksum`, `InvalidLength`, `UnsupportedPrefix`
   - Detailed error reporting with position information
   - Graceful handling of edge cases

5. **Configuration System**
   - Configurable validation strictness levels
   - Enable/disable specific ISBN formats
   - Performance tuning options (worker pools, timeouts)

### Performance Requirements
- **Throughput**: >500 validations/second for single items
- **Batch processing**: >2000 items/second with concurrency  
- **Memory efficiency**: <1MB heap growth per 1000 validations

---

## 🏗️ Architecture Requirements

### Project Structure
```
cmd/
├── cli/           # Command-line interface
└── server/        # HTTP API (bonus)
pkg/
├── isbn/          # Core validation logic
├── batch/         # Batch processing
└── config/        # Configuration management
internal/
└── errors/        # Custom error types
```

### Key Interfaces
```go
type Validator interface {
    Validate(isbn string) (*Result, error)
    ValidateBatch(isbns []string, opts ...Option) ([]*Result, error)
    ProcessCSV(inputPath, outputPath string) error
}

type Result struct {
    ISBN     string    `csv:"isbn"`
    Valid    bool      `csv:"valid"`  
    Format   string    `csv:"format"` // "ISBN-10" or "ISBN-13"
    Error    string    `csv:"error,omitempty"`
    Duration string    `csv:"duration_ms"`
}
```

---

## 🧪 Test Data

A sample CSV file (`test_isbns.csv`) is provided with test cases including:

**Valid Cases:**
- `0-306-40615-2` (ISBN-10 with hyphens)
- `9780306406157` (ISBN-13, same book)
- `978-0-321-14653-1` (ISBN-13 with hyphens)
- `ISBN-10: 123456789X` (with prefix text)
- `   978 0 306 40615 7   ` (extra whitespace)

**Invalid Cases:**
- `123456789Y` (invalid check digit)
- `978030640615` (wrong length)
- `abc-def-ghijk-l` (non-numeric)
- `""` (empty string)
- Various checksum mismatches
- Invalid 979-prefix ISBN-13 with wrong checksum

**CSV Format:**
```csv
isbn,expected_valid,format,description
0-306-40615-2,true,ISBN-10,Classic valid ISBN-10 with hyphens
1234567890,false,ISBN-10,Invalid checksum
```

Your CLI should process this CSV and output validation results in the same format with additional columns for actual results.

---

## 🚀 Deliverables

1. **Core Library** - Complete validation system
2. **CLI Tool** - Accept CSV file input, CSV output  
3. **Comprehensive Tests** - Unit, integration, benchmarks
4. **Documentation** - API docs and usage examples
5. **Performance Analysis** - Benchmark results and optimizations

---

## 🏆 Bonus Challenges *(+10 min each)*

1. **Format Conversion**: Convert between ISBN-10 ⟷ ISBN-13
2. **Publisher Validation**: Validate against real publisher ranges
3. **HTTP API**: REST endpoints with rate limiting
4. **Fuzzy Matching**: Handle OCR errors and suggest corrections

---

## 📊 Evaluation Criteria

- **Correctness** (40%): Handles all formats and edge cases
- **Performance** (25%): Meets throughput requirements  
- **Code Quality** (20%): Clean, testable, maintainable
- **Error Handling** (10%): Comprehensive error reporting
- **Documentation** (5%): Clear API and examples

---

## 🔧 Getting Started

1. **Analyze** the existing basic implementation
2. **Design** your architecture and interfaces
3. **Implement** core validation with proper algorithms
4. **Add** batch processing with concurrency
5. **Build** comprehensive test suite with benchmarks
6. **Optimize** for performance requirements
7. **Document** usage and API

The current basic implementation provides a foundation but requires significant enhancement to meet these advanced requirements.

---

👉 **This challenge evaluates advanced Go programming, concurrent processing, API design, and performance optimization skills.**  
