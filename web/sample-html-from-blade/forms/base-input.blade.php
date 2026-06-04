@extends('layouts.app')

@section('title', 'Base Input')

@push('styles')

@endpush

@section('content')
<div class="page-header">
    <div>
        <h1>Base Input</h1>
        <p>Semua tipe input HTML5 dengan berbagai variasi dan styling</p>
    </div>
    <div style="display: flex; gap: 10px;">
        <button class="btn btn-secondary">
            <i class="fa-solid fa-book"></i>
            <span>Documentation</span>
        </button>
        <button class="btn btn-primary">
            <i class="fa-solid fa-code"></i>
            <span>View Code</span>
        </button>
    </div>
</div>

<!-- Info Alert -->
<div class="content-card" style="margin-bottom: 24px; border-left: 4px solid var(--accent);">
    <div class="card-body" style="padding: 16px 20px;">
        <div style="display: flex; gap: 12px; align-items: start;">
            <i class="fa-solid fa-circle-info" style="color: var(--accent); font-size: 20px; margin-top: 2px;"></i>
            <div style="flex: 1;">
                <h4 style="margin-bottom: 4px; font-size: 14px;">HTML5 Input Types</h4>
                <p style="font-size: 13px; color: var(--text-secondary); margin: 0;">Halaman ini menampilkan semua tipe input HTML5 yang didukung. Setiap input sudah responsive dan mengikuti Fluent Design System.</p>
            </div>
        </div>
    </div>
</div>

<!-- Text Input Types -->
<div class="input-section-title">
    <i class="fa-solid fa-font"></i>
    Text Input Types <span class="badge badge-primary">Essential</span>
</div>

<div class="base-input-grid">
    <div class="content-card">
        <div class="card-header">
            <div class="card-header-left">
                <div class="card-icon bg-primary">
                    <i class="fa-solid fa-keyboard"></i>
                </div>
                <div>
                    <h3>Basic Text Inputs</h3>
                    <p class="card-subtitle">Text, email, password, search</p>
                </div>
            </div>
        </div>
        <div class="card-body">
            <div class="form-group">
                <label>Text Input</label>
                <input type="text" class="form-control" placeholder="Enter text">
                <div class="helper-text">
                    <i class="fa-solid fa-circle-info"></i>
                    Basic single-line text field
                </div>
            </div>

            <div class="form-group">
                <label>Email Input</label>
                <div class="input-with-icon">
                    <i class="fa-solid fa-envelope"></i>
                    <input type="email" class="form-control" placeholder="name@example.com">
                </div>
                <div class="helper-text">
                    <i class="fa-solid fa-shield-halved"></i>
                    Auto-validates email format
                </div>
            </div>

            <div class="form-group">
                <label>Password Input</label>
                <div class="input-with-icon-right">
                    <input type="password" class="form-control" placeholder="••••••••" value="MyPassword123">
                    <i class="fa-solid fa-eye" title="Toggle password visibility"></i>
                </div>
                <div class="helper-text">
                    <i class="fa-solid fa-lock"></i>
                    Masked input for security
                </div>
            </div>

            <div class="form-group">
                <label>Search Input</label>
                <div class="input-with-icon">
                    <i class="fa-solid fa-magnifying-glass"></i>
                    <input type="search" class="form-control" placeholder="Search...">
                </div>
                <div class="helper-text">
                    <i class="fa-solid fa-circle-info"></i>
                    With clear button (×)
                </div>
            </div>
        </div>
    </div>

    <div class="content-card">
        <div class="card-header">
            <div class="card-header-left">
                <div class="card-icon bg-success">
                    <i class="fa-solid fa-text-width"></i>
                </div>
                <div>
                    <h3>URL & Textarea</h3>
                    <p class="card-subtitle">Web addresses and multi-line</p>
                </div>
            </div>
        </div>
        <div class="card-body">
            <div class="form-group">
                <label>URL Input</label>
                <div class="input-with-icon">
                    <i class="fa-solid fa-link"></i>
                    <input type="url" class="form-control" placeholder="https://example.com">
                </div>
                <div class="helper-text">
                    <i class="fa-solid fa-circle-info"></i>
                    Validates URL format
                </div>
            </div>

            <div class="form-group">
                <label>Textarea</label>
                <textarea class="form-control" rows="4" placeholder="Enter multiple lines of text..."></textarea>
                <div class="helper-text">
                    <i class="fa-solid fa-align-left"></i>
                    Multi-line text input
                </div>
            </div>

            <div class="form-group">
                <label>Read-only Input</label>
                <input type="text" class="form-control" value="This cannot be edited" readonly style="background-color: var(--bg-secondary); cursor: not-allowed;">
                <div class="helper-text">
                    <i class="fa-solid fa-ban"></i>
                    Read-only field
                </div>
            </div>

            <div class="form-group">
                <label>Disabled Input</label>
                <input type="text" class="form-control" placeholder="This is disabled" disabled style="opacity: 0.5; cursor: not-allowed;">
                <div class="helper-text">
                    <i class="fa-solid fa-lock"></i>
                    Disabled field
                </div>
            </div>
        </div>
    </div>
</div>

<!-- Number & Range Inputs -->
<div class="input-section-title">
    <i class="fa-solid fa-hashtag"></i>
    Number & Range Inputs
</div>

<div class="base-input-grid three-cols">
    <div class="content-card">
        <div class="card-header">
            <div class="card-header-left">
                <div class="card-icon bg-info">
                    <i class="fa-solid fa-sort-numeric-up"></i>
                </div>
                <div>
                    <h3>Number Input</h3>
                    <p class="card-subtitle">Numeric values with controls</p>
                </div>
            </div>
        </div>
        <div class="card-body">
            <div class="form-group">
                <label>Basic Number</label>
                <input type="number" class="form-control" placeholder="0" min="0" max="100" step="1">
                <div class="helper-text">
                    <i class="fa-solid fa-circle-info"></i>
                    Spin buttons included
                </div>
            </div>

            <div class="form-group">
                <label>Decimal Number</label>
                <input type="number" class="form-control" placeholder="0.00" step="0.01" min="0" max="1000">
                <div class="helper-text">
                    <i class="fa-solid fa-coins"></i>
                    Step: 0.01 (2 decimals)
                </div>
            </div>

            <div class="form-group">
                <label>Number with Range</label>
                <input type="number" class="form-control" value="50" min="0" max="100" step="5">
                <div class="helper-text">
                    <i class="fa-solid fa-arrows-left-right"></i>
                    Min: 0, Max: 100, Step: 5
                </div>
            </div>
        </div>
    </div>

    <div class="content-card">
        <div class="card-header">
            <div class="card-header-left">
                <div class="card-icon bg-warning">
                    <i class="fa-solid fa-sliders"></i>
                </div>
                <div>
                    <h3>Range Slider</h3>
                    <p class="card-subtitle">Visual range selector</p>
                </div>
            </div>
        </div>
        <div class="card-body">
            <div class="form-group">
                <label>Volume (0-100)</label>
                <input type="range" class="form-control" min="0" max="100" value="75" style="padding: 0;">
                <div class="helper-text">
                    <i class="fa-solid fa-volume-high"></i>
                    Current: 75%
                </div>
            </div>

            <div class="form-group">
                <label>Price Range ($0-$1000)</label>
                <input type="range" class="form-control" min="0" max="1000" value="500" step="50" style="padding: 0;">
                <div class="helper-text">
                    <i class="fa-solid fa-dollar-sign"></i>
                    Current: $500
                </div>
            </div>

            <div class="form-group">
                <label>Brightness (0-255)</label>
                <input type="range" class="form-control" min="0" max="255" value="128" style="padding: 0;">
                <div class="helper-text">
                    <i class="fa-solid fa-sun"></i>
                    Current: 128
                </div>
            </div>
        </div>
    </div>

    <div class="content-card">
        <div class="card-header">
            <div class="card-header-left">
                <div class="card-icon bg-success">
                    <i class="fa-solid fa-phone"></i>
                </div>
                <div>
                    <h3>Telephone</h3>
                    <p class="card-subtitle">Phone number input</p>
                </div>
            </div>
        </div>
        <div class="card-body">
            <div class="form-group">
                <label>Phone Number</label>
                <div class="input-with-icon">
                    <i class="fa-solid fa-phone"></i>
                    <input type="tel" class="form-control" placeholder="+62 812-3456-7890">
                </div>
                <div class="helper-text">
                    <i class="fa-solid fa-circle-info"></i>
                    Opens numeric keyboard on mobile
                </div>
            </div>

            <div class="form-group">
                <label>Fax Number</label>
                <div class="input-with-icon">
                    <i class="fa-solid fa-fax"></i>
                    <input type="tel" class="form-control" placeholder="+62 21-123-4567">
                </div>
                <div class="helper-text">
                    <i class="fa-solid fa-circle-info"></i>
                    Same type as phone
                </div>
            </div>

            <div class="divider"></div>

            <div class="form-group">
                <h4 style="font-size: 14px; margin-bottom: 12px;">Features:</h4>
                <ul class="feature-list">
                    <li>
                        <i class="fa-solid fa-circle-check"></i>
                        <span>Numeric keyboard on mobile</span>
                    </li>
                    <li>
                        <i class="fa-solid fa-circle-check"></i>
                        <span>Pattern validation support</span>
                    </li>
                    <li>
                        <i class="fa-solid fa-circle-check"></i>
                        <span>Auto-formatting with JS</span>
                    </li>
                </ul>
            </div>
        </div>
    </div>
</div>

<!-- Date & Time Inputs -->
<div class="input-section-title">
    <i class="fa-solid fa-calendar-days"></i>
    Date & Time Inputs <span class="badge badge-success">Popular</span>
</div>

<div class="base-input-grid">
    <div class="content-card">
        <div class="card-header">
            <div class="card-header-left">
                <div class="card-icon bg-primary">
                    <i class="fa-solid fa-calendar"></i>
                </div>
                <div>
                    <h3>Date Inputs</h3>
                    <p class="card-subtitle">Date, month, week, time</p>
                </div>
            </div>
        </div>
        <div class="card-body">
            <div class="form-group">
                <label>Date Picker</label>
                <input type="date" class="form-control" value="2026-05-31">
                <div class="helper-text">
                    <i class="fa-solid fa-calendar-days"></i>
                    Format: YYYY-MM-DD
                </div>
            </div>

            <div class="form-group">
                <label>Month Picker</label>
                <input type="month" class="form-control" value="2026-05">
                <div class="helper-text">
                    <i class="fa-solid fa-calendar-week"></i>
                    Format: YYYY-MM
                </div>
            </div>

            <div class="form-group">
                <label>Week Picker</label>
                <input type="week" class="form-control" value="2026-W22">
                <div class="helper-text">
                    <i class="fa-solid fa-calendar"></i>
                    Format: YYYY-W##
                </div>
            </div>

            <div class="form-group">
                <label>Time Picker</label>
                <input type="time" class="form-control" value="14:30">
                <div class="helper-text">
                    <i class="fa-solid fa-clock"></i>
                    Format: HH:MM (24h)
                </div>
            </div>
        </div>
    </div>

    <div class="content-card">
        <div class="card-header">
            <div class="card-header-left">
                <div class="card-icon bg-info">
                    <i class="fa-solid fa-calendar-check"></i>
                </div>
                <div>
                    <h3>DateTime Local</h3>
                    <p class="card-subtitle">Combined date and time</p>
                </div>
            </div>
        </div>
        <div class="card-body">
            <div class="form-group">
                <label>Event Date & Time</label>
                <input type="datetime-local" class="form-control" value="2026-06-15T14:30">
                <div class="helper-text">
                    <i class="fa-solid fa-calendar-check"></i>
                    Date + Time combined
                </div>
            </div>

            <div class="divider"></div>

            <div class="form-group">
                <label>Date with Min/Max</label>
                <input type="date" class="form-control" min="2026-01-01" max="2026-12-31">
                <div class="helper-text">
                    <i class="fa-solid fa-arrows-left-right"></i>
                    Range: Jan 1 - Dec 31, 2026
                </div>
            </div>

            <div class="form-group">
                <label>Disabled Weekends</label>
                <input type="date" class="form-control" value="2026-06-01">
                <div class="helper-text">
                    <i class="fa-solid fa-ban"></i>
                    JavaScript validation required
                </div>
            </div>

            <div class="divider"></div>

            <div class="form-group">
                <h4 style="font-size: 14px; margin-bottom: 12px;">Supported Pickers:</h4>
                <ul class="feature-list">
                    <li>
                        <i class="fa-solid fa-circle-check"></i>
                        <span><strong>date:</strong> Calendar date picker</span>
                    </li>
                    <li>
                        <i class="fa-solid fa-circle-check"></i>
                        <span><strong>month:</strong> Month selector</span>
                    </li>
                    <li>
                        <i class="fa-solid fa-circle-check"></i>
                        <span><strong>week:</strong> Week selector</span>
                    </li>
                    <li>
                        <i class="fa-solid fa-circle-check"></i>
                        <span><strong>time:</strong> Time picker</span>
                    </li>
                    <li>
                        <i class="fa-solid fa-circle-check"></i>
                        <span><strong>datetime-local:</strong> Date + Time</span>
                    </li>
                </ul>
            </div>
        </div>
    </div>
</div>

<!-- Special Input Types -->
<div class="input-section-title">
    <i class="fa-solid fa-wand-magic-sparkles"></i>
    Special Input Types
</div>

<div class="base-input-grid three-cols">
    <div class="content-card">
        <div class="card-header">
            <div class="card-header-left">
                <div class="card-icon bg-warning">
                    <i class="fa-solid fa-palette"></i>
                </div>
                <div>
                    <h3>Color Picker</h3>
                    <p class="card-subtitle">Visual color selection</p>
                </div>
            </div>
        </div>
        <div class="card-body">
            <div class="form-group">
                <label>Primary Color</label>
                <div style="display: flex; gap: 12px; align-items: center;">
                    <input type="color" class="form-control" value="#0078D4" style="width: 60px; height: 40px; padding: 2px; cursor: pointer;">
                    <input type="text" class="form-control" value="#0078D4" style="flex: 1;" readonly>
                </div>
                <div class="helper-text">
                    <i class="fa-solid fa-palette"></i>
                    Click to open color picker
                </div>
            </div>

            <div class="form-group">
                <label>Background Color</label>
                <div style="display: flex; gap: 12px; align-items: center;">
                    <input type="color" class="form-control" value="#16C60C" style="width: 60px; height: 40px; padding: 2px; cursor: pointer;">
                    <input type="text" class="form-control" value="#16C60C" style="flex: 1;" readonly>
                </div>
                <div class="helper-text">
                    <i class="fa-solid fa-fill-drip"></i>
                    HEX color code
                </div>
            </div>

            <div class="divider"></div>

            <div class="code-block">
                <div style="color: var(--text-tertiary); margin-bottom: 8px;">HTML:</div>
                <code style="color: var(--accent);">
                    &lt;input type="color" <br>
                    &nbsp;&nbsp;value="#0078D4"&gt;
                </code>
            </div>
        </div>
    </div>

    <div class="content-card">
        <div class="card-header">
            <div class="card-header-left">
                <div class="card-icon bg-success">
                    <i class="fa-solid fa-file-arrow-up"></i>
                </div>
                <div>
                    <h3>File Upload</h3>
                    <p class="card-subtitle">File selection input</p>
                </div>
            </div>
        </div>
        <div class="card-body">
            <div class="form-group">
                <label>Single File</label>
                <input type="file" class="form-control">
                <div class="helper-text">
                    <i class="fa-solid fa-file"></i>
                    Select one file
                </div>
            </div>

            <div class="form-group">
                <label>Multiple Files</label>
                <input type="file" class="form-control" multiple>
                <div class="helper-text">
                    <i class="fa-solid fa-file-zipper"></i>
                    Select multiple files
                </div>
            </div>

            <div class="form-group">
                <label>Image Only</label>
                <input type="file" class="form-control" accept="image/*">
                <div class="helper-text">
                    <i class="fa-solid fa-image"></i>
                    accept="image/*"
                </div>
            </div>

            <div class="form-group">
                <label>PDF Documents</label>
                <input type="file" class="form-control" accept=".pdf,application/pdf">
                <div class="helper-text">
                    <i class="fa-solid fa-file-pdf"></i>
                    accept=".pdf"
                </div>
            </div>
        </div>
    </div>

    <div class="content-card">
        <div class="card-header">
            <div class="card-header-left">
                <div class="card-icon bg-info">
                    <i class="fa-solid fa-eye-dropper"></i>
                </div>
                <div>
                    <h3>Hidden Input</h3>
                    <p class="card-subtitle">Invisible form data</p>
                </div>
            </div>
        </div>
        <div class="card-body">
            <div class="form-group">
                <label>Hidden Input Example</label>
                <input type="hidden" name="csrf_token" value="abc123xyz">
                <div class="code-block" style="margin-top: 0;">
                    <code style="color: var(--accent);">
                        &lt;input type="hidden" <br>
                        &nbsp;&nbsp;name="csrf_token"<br>
                        &nbsp;&nbsp;value="abc123xyz"&gt;
                    </code>
                </div>
                <div class="helper-text" style="margin-top: 12px;">
                    <i class="fa-solid fa-shield-halved"></i>
                    Not visible to users
                </div>
            </div>

            <div class="divider"></div>

            <div class="form-group">
                <h4 style="font-size: 14px; margin-bottom: 12px;">Use Cases:</h4>
                <ul class="feature-list">
                    <li>
                        <i class="fa-solid fa-circle-check"></i>
                        <span>CSRF tokens</span>
                    </li>
                    <li>
                        <i class="fa-solid fa-circle-check"></i>
                        <span>Form submission IDs</span>
                    </li>
                    <li>
                        <i class="fa-solid fa-circle-check"></i>
                        <span>User session data</span>
                    </li>
                    <li>
                        <i class="fa-solid fa-circle-check"></i>
                        <span>Tracking parameters</span>
                    </li>
                </ul>
            </div>
        </div>
    </div>
</div>

<!-- Input Sizes & States -->
<div class="input-section-title">
    <i class="fa-solid fa-up-right-and-down-left-from-center"></i>
    Input Sizes & States
</div>

<div class="base-input-grid">
    <div class="content-card">
        <div class="card-header">
            <div class="card-header-left">
                <div class="card-icon bg-primary">
                    <i class="fa-solid fa-expand"></i>
                </div>
                <div>
                    <h3>Input Sizes</h3>
                    <p class="card-subtitle">Small, default, and large</p>
                </div>
            </div>
        </div>
        <div class="card-body">
            <div class="input-size-demo">
                <div class="size-label">Small</div>
                <input type="text" class="form-control form-control-sm" placeholder="Small input" style="padding: 6px 12px; font-size: 12px;">
            </div>

            <div class="input-size-demo">
                <div class="size-label">Default</div>
                <input type="text" class="form-control" placeholder="Default input">
            </div>

            <div class="input-size-demo">
                <div class="size-label">Large</div>
                <input type="text" class="form-control form-control-lg" placeholder="Large input" style="padding: 14px 16px; font-size: 16px;">
            </div>

            <div class="divider"></div>

            <div class="form-group">
                <label>Full Width Input</label>
                <input type="text" class="form-control" placeholder="This takes full width" style="width: 100%;">
            </div>
        </div>
    </div>

    <div class="content-card">
        <div class="card-header">
            <div class="card-header-left">
                <div class="card-icon bg-success">
                    <i class="fa-solid fa-circle-check"></i>
                </div>
                <div>
                    <h3>Input States</h3>
                    <p class="card-subtitle">Valid, invalid, and disabled</p>
                </div>
            </div>
        </div>
        <div class="card-body">
            <div class="input-state-demo">
                <label>Valid State</label>
                <input type="text" class="form-control is-valid" value="Valid input" style="border-color: var(--success);">
                <div class="helper-text" style="color: var(--success);">
                    <i class="fa-solid fa-circle-check"></i>
                    Looks good!
                </div>
            </div>

            <div class="input-state-demo">
                <label>Invalid State</label>
                <input type="text" class="form-control is-invalid" value="Invalid" style="border-color: var(--danger);">
                <div class="helper-text" style="color: var(--danger);">
                    <i class="fa-solid fa-circle-xmark"></i>
                    Please fix this field
                </div>
            </div>

            <div class="input-state-demo">
                <label>Disabled State</label>
                <input type="text" class="form-control" value="Disabled" disabled style="opacity: 0.5; cursor: not-allowed; background-color: var(--bg-secondary);">
            </div>

            <div class="input-state-demo">
                <label>Read-only State</label>
                <input type="text" class="form-control" value="Read-only" readonly style="background-color: var(--bg-secondary); cursor: not-allowed;">
            </div>
        </div>
    </div>
</div>

<!-- Complete Input Examples -->
<div class="input-section-title">
    <i class="fa-solid fa-clipboard-list"></i>
    Complete Input Examples
</div>

<div class="base-input-grid full-width">
    <div class="content-card">
        <div class="card-header">
            <div class="card-header-left">
                <div class="card-icon bg-primary">
                    <i class="fa-solid fa-user-pen"></i>
                </div>
                <div>
                    <h3>Contact Information Form</h3>
                    <p class="card-subtitle">Real-world example with various input types</p>
                </div>
            </div>
        </div>
        <div class="card-body">
            <form style="max-width: 800px;">
                <div style="display: grid; grid-template-columns: 1fr 1fr; gap: 20px;">
                    <div class="form-group">
                        <label>First Name <span style="color: var(--danger);">*</span></label>
                        <input type="text" class="form-control" placeholder="John" required>
                    </div>
                    <div class="form-group">
                        <label>Last Name <span style="color: var(--danger);">*</span></label>
                        <input type="text" class="form-control" placeholder="Doe" required>
                    </div>
                </div>

                <div class="form-group">
                    <label>Email Address <span style="color: var(--danger);">*</span></label>
                    <div class="input-with-icon">
                        <i class="fa-solid fa-envelope"></i>
                        <input type="email" class="form-control" placeholder="john@example.com" required>
                    </div>
                </div>

                <div class="form-group">
                    <label>Phone Number</label>
                    <div class="input-with-icon">
                        <i class="fa-solid fa-phone"></i>
                        <input type="tel" class="form-control" placeholder="+62 812-3456-7890">
                    </div>
                </div>

                <div class="form-group">
                    <label>Website</label>
                    <div class="input-with-icon">
                        <i class="fa-solid fa-globe"></i>
                        <input type="url" class="form-control" placeholder="https://johndoe.com">
                    </div>
                </div>

                <div class="form-group">
                    <label>Birth Date</label>
                    <input type="date" class="form-control">
                </div>

                <div class="form-group">
                    <label>Profile Photo</label>
                    <input type="file" class="form-control" accept="image/*">
                </div>

                <div class="form-group">
                    <label>Bio</label>
                    <textarea class="form-control" rows="4" placeholder="Tell us about yourself..."></textarea>
                </div>

                <div style="display: flex; gap: 12px; margin-top: 24px;">
                    <button type="submit" class="btn btn-primary">
                        <i class="fa-solid fa-paper-plane"></i>
                        Save Changes
                    </button>
                    <button type="reset" class="btn btn-secondary">
                        <i class="fa-solid fa-rotate-left"></i>
                        Reset
                    </button>
                </div>
            </form>
        </div>
    </div>
</div>

<!-- Input Features Summary -->
<div class="base-input-grid full-width">
    <div class="content-card">
        <div class="card-header">
            <div class="card-header-left">
                <div class="card-icon bg-info">
                    <i class="fa-solid fa-list-check"></i>
                </div>
                <div>
                    <h3>All Supported Input Types</h3>
                    <p class="card-subtitle">Complete list of HTML5 input types</p>
                </div>
            </div>
        </div>
        <div class="card-body">
            <div class="base-input-grid three-cols">
                <div>
                    <h4 style="font-size: 14px; margin-bottom: 12px;">Text Inputs:</h4>
                    <ul class="feature-list">
                        <li><i class="fa-solid fa-check"></i> text</li>
                        <li><i class="fa-solid fa-check"></i> email</li>
                        <li><i class="fa-solid fa-check"></i> password</li>
                        <li><i class="fa-solid fa-check"></i> search</li>
                        <li><i class="fa-solid fa-check"></i> tel</li>
                        <li><i class="fa-solid fa-check"></i> url</li>
                    </ul>
                </div>
                <div>
                    <h4 style="font-size: 14px; margin-bottom: 12px;">Number & Date:</h4>
                    <ul class="feature-list">
                        <li><i class="fa-solid fa-check"></i> number</li>
                        <li><i class="fa-solid fa-check"></i> range</li>
                        <li><i class="fa-solid fa-check"></i> date</li>
                        <li><i class="fa-solid fa-check"></i> month</li>
                        <li><i class="fa-solid fa-check"></i> week</li>
                        <li><i class="fa-solid fa-check"></i> time</li>
                        <li><i class="fa-solid fa-check"></i> datetime-local</li>
                    </ul>
                </div>
                <div>
                    <h4 style="font-size: 14px; margin-bottom: 12px;">Special:</h4>
                    <ul class="feature-list">
                        <li><i class="fa-solid fa-check"></i> color</li>
                        <li><i class="fa-solid fa-check"></i> file</li>
                        <li><i class="fa-solid fa-check"></i> hidden</li>
                        <li><i class="fa-solid fa-check"></i> checkbox</li>
                        <li><i class="fa-solid fa-check"></i> radio</li>
                        <li><i class="fa-solid fa-check"></i> submit</li>
                    </ul>
                </div>
            </div>
        </div>
    </div>
</div>
@endsection
