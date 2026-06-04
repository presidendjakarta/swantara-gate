@extends('layouts.app')

@section('title', 'Form Validation')

@push('styles')
@endpush

@push('scripts')
<script src="{{ asset('assets/js/form-validation.js') }}"></script>
@endpush

@section('content')
<div class="page-header">
    <div>
        <h1>Form Validation</h1>
        <p>Sistem validasi form lengkap dengan real-time feedback dan error handling</p>
    </div>
    <div style="display: flex; gap: 10px;">
        <button class="btn btn-secondary">
            <i class="fa-solid fa-book"></i>
            <span>Documentation</span>
        </button>
        <button class="btn btn-primary">
            <i class="fa-solid fa-shield-halved"></i>
            <span>Validation Rules</span>
        </button>
    </div>
</div>

<!-- Info Alert -->
<div class="content-card" style="margin-bottom: 24px; border-left: 4px solid var(--warning);">
    <div class="card-body" style="padding: 16px 20px;">
        <div style="display: flex; gap: 12px; align-items: start;">
            <i class="fa-solid fa-triangle-exclamation" style="color: var(--warning); font-size: 20px; margin-top: 2px;"></i>
            <div style="flex: 1;">
                <h4 style="margin-bottom: 4px; font-size: 14px;">Client-Side Validation</h4>
                <p style="font-size: 13px; color: var(--text-secondary); margin: 0;">Semua validasi di halaman ini menggunakan HTML5 validation dan JavaScript untuk real-time feedback. Server-side validation tetap diperlukan untuk keamanan.</p>
            </div>
        </div>
    </div>
</div>

<!-- Required Field Validation -->
<div class="validation-section-title">
    <i class="fa-solid fa-circle-exclamation"></i>
    Required Field Validation <span class="badge badge-danger">Required</span>
</div>

<div class="validation-demo-grid">
    <div class="content-card">
        <div class="card-header">
            <div class="card-header-left">
                <div class="card-icon bg-primary">
                    <i class="fa-solid fa-asterisk"></i>
                </div>
                <div>
                    <h3>Basic Required Fields</h3>
                    <p class="card-subtitle">Field wajib dengan asterisk (*)</p>
                </div>
            </div>
        </div>
        <div class="card-body">
            <form id="requiredForm">
                <div class="form-group">
                    <label>Full Name <span style="color: var(--danger);">*</span></label>
                    <input type="text" class="form-control is-invalid" value="" placeholder="Enter your full name" required>
                    <div class="validation-feedback invalid-feedback">
                        <i class="fa-solid fa-circle-xmark"></i>
                        <span>Nama lengkap wajib diisi</span>
                    </div>
                </div>

                <div class="form-group">
                    <label>Email Address <span style="color: var(--danger);">*</span></label>
                    <input type="email" class="form-control" placeholder="name@example.com" required>
                </div>

                <div class="form-group">
                    <label>Phone Number <span style="color: var(--danger);">*</span></label>
                    <input type="tel" class="form-control" placeholder="+62 812-3456-7890" required>
                </div>

                <div class="form-group">
                    <label>Subject <span style="color: var(--danger);">*</span></label>
                    <select class="form-control" required>
                        <option value="">-- Select Subject --</option>
                        <option value="general">General Inquiry</option>
                        <option value="support">Technical Support</option>
                        <option value="sales">Sales</option>
                    </select>
                </div>

                <button type="submit" class="btn btn-primary">
                    <i class="fa-solid fa-paper-plane"></i>
                    Submit Form
                </button>
            </form>
        </div>
    </div>

    <div class="content-card">
        <div class="card-header">
            <div class="card-header-left">
                <div class="card-icon bg-success">
                    <i class="fa-solid fa-circle-check"></i>
                </div>
                <div>
                    <h3>Validation States</h3>
                    <p class="card-subtitle">Error, success, and warning states</p>
                </div>
            </div>
        </div>
        <div class="card-body">
            <div class="form-group">
                <label style="color: var(--success);">Valid Input</label>
                <input type="text" class="form-control is-valid" value="John Doe">
                <div class="validation-feedback valid-feedback">
                    <i class="fa-solid fa-circle-check"></i>
                    <span>Looks good!</span>
                </div>
            </div>

            <div class="form-group">
                <label style="color: var(--danger);">Invalid Input</label>
                <input type="text" class="form-control is-invalid" value="abc">
                <div class="validation-feedback invalid-feedback">
                    <i class="fa-solid fa-circle-xmark"></i>
                    <span>Minimum 3 characters required</span>
                </div>
            </div>

            <div class="form-group">
                <label style="color: var(--warning);">Warning State</label>
                <input type="text" class="form-control" value="test@email" style="border-color: var(--warning); background-color: var(--warning-bg);">
                <div class="validation-feedback warning-feedback">
                    <i class="fa-solid fa-triangle-exclamation"></i>
                    <span>Email format mungkin tidak lengkap</span>
                </div>
            </div>

            <div class="divider"></div>

            <div class="form-group">
                <h4 style="font-size: 14px; margin-bottom: 12px;">Validation Features:</h4>
                <ul class="feature-list">
                    <li>
                        <i class="fa-solid fa-circle-check"></i>
                        <span>Real-time validation on input</span>
                    </li>
                    <li>
                        <i class="fa-solid fa-circle-check"></i>
                        <span>Color-coded feedback states</span>
                    </li>
                    <li>
                        <i class="fa-solid fa-circle-check"></i>
                        <span>Icon indicators (✓ / ✕)</span>
                    </li>
                    <li>
                        <i class="fa-solid fa-circle-check"></i>
                        <span>Descriptive error messages</span>
                    </li>
                </ul>
            </div>
        </div>
    </div>
</div>

<!-- Pattern & Format Validation -->
<div class="validation-section-title">
    <i class="fa-solid fa-fingerprint"></i>
    Pattern & Format Validation
</div>

<div class="validation-demo-grid">
    <div class="content-card">
        <div class="card-header">
            <div class="card-header-left">
                <div class="card-icon bg-info">
                    <i class="fa-solid fa-envelope"></i>
                </div>
                <div>
                    <h3>Email Validation</h3>
                    <p class="card-subtitle">Format email dengan regex pattern</p>
                </div>
            </div>
        </div>
        <div class="card-body">
            <div class="form-group">
                <label>Email Address</label>
                <input type="email" class="form-control is-valid" value="user@example.com" pattern="[a-z0-9._%+-]+@[a-z0-9.-]+\.[a-z]{2,}$">
                <div class="validation-feedback valid-feedback">
                    <i class="fa-solid fa-circle-check"></i>
                    <span>Valid email format</span>
                </div>
            </div>

            <div class="form-group">
                <label>Email (Invalid)</label>
                <input type="email" class="form-control is-invalid" value="user@.com">
                <div class="validation-feedback invalid-feedback">
                    <i class="fa-solid fa-circle-xmark"></i>
                    <span>Please enter a valid email address</span>
                </div>
            </div>

            <div class="divider"></div>

            <div class="form-group">
                <label>Website URL</label>
                <input type="url" class="form-control is-valid" value="https://example.com" pattern="https://.*">
                <div class="validation-feedback valid-feedback">
                    <i class="fa-solid fa-circle-check"></i>
                    <span>Valid URL format</span>
                </div>
            </div>

            <div class="form-group">
                <label>Phone Number (Indonesian)</label>
                <input type="tel" class="form-control is-invalid" value="0812" pattern="^\\+62\\d{8,12}$">
                <div class="validation-feedback invalid-feedback">
                    <i class="fa-solid fa-circle-xmark"></i>
                    <span>Format: +62xxxxxxxxxx</span>
                </div>
            </div>
        </div>
    </div>

    <div class="content-card">
        <div class="card-header">
            <div class="card-header-left">
                <div class="card-icon bg-warning">
                    <i class="fa-solid fa-lock"></i>
                </div>
                <div>
                    <h3>Password Validation <span class="badge badge-warning">Important</span></h3>
                    <p class="card-subtitle">Strong password requirements</p>
                </div>
            </div>
        </div>
        <div class="card-body">
            <div class="form-group">
                <label>Password</label>
                <input type="password" id="passwordInput" class="form-control" placeholder="Enter password" value="MyP@ss123">
                
                <div class="password-strength">
                    <div class="strength-meter">
                        <div class="strength-meter-fill strength-strong"></div>
                    </div>
                    <div class="strength-text" style="color: var(--success);">Strong</div>
                </div>

                <ul class="requirement-list">
                    <li class="met">
                        <i class="fa-solid fa-circle-check"></i>
                        <span>Minimum 8 characters</span>
                    </li>
                    <li class="met">
                        <i class="fa-solid fa-circle-check"></i>
                        <span>At least one uppercase letter</span>
                    </li>
                    <li class="met">
                        <i class="fa-solid fa-circle-check"></i>
                        <span>At least one number</span>
                    </li>
                    <li class="met">
                        <i class="fa-solid fa-circle-check"></i>
                        <span>At least one special character</span>
                    </li>
                </ul>
            </div>

            <div class="divider"></div>

            <div class="form-group">
                <label>Confirm Password</label>
                <input type="password" class="form-control is-invalid" value="different123">
                <div class="validation-feedback invalid-feedback">
                    <i class="fa-solid fa-circle-xmark"></i>
                    <span>Passwords do not match</span>
                </div>
            </div>
        </div>
    </div>
</div>

<!-- Length & Range Validation -->
<div class="validation-section-title">
    <i class="fa-solid fa-ruler"></i>
    Length & Range Validation
</div>

<div class="validation-demo-grid three-cols">
    <div class="content-card">
        <div class="card-header">
            <div class="card-header-left">
                <div class="card-icon bg-primary">
                    <i class="fa-solid fa-text-width"></i>
                </div>
                <div>
                    <h3>Min/Max Length</h3>
                    <p class="card-subtitle">Character length validation</p>
                </div>
            </div>
        </div>
        <div class="card-body">
            <div class="form-group">
                <label>Username (3-20 chars)</label>
                <input type="text" class="form-control is-valid" value="johndoe" minlength="3" maxlength="20">
                <div class="validation-feedback valid-feedback">
                    <i class="fa-solid fa-circle-check"></i>
                    <span>7/20 characters</span>
                </div>
            </div>

            <div class="form-group">
                <label>Username (Too Short)</label>
                <input type="text" class="form-control is-invalid" value="ab" minlength="3" maxlength="20">
                <div class="validation-feedback invalid-feedback">
                    <i class="fa-solid fa-circle-xmark"></i>
                    <span>Minimum 3 characters (2/20)</span>
                </div>
            </div>

            <div class="form-group">
                <label>Textarea (Max 200 chars)</label>
                <textarea class="form-control is-invalid" rows="3" maxlength="200">This is a very long text that exceeds the maximum allowed character limit for this field...</textarea>
                <div class="validation-feedback invalid-feedback">
                    <i class="fa-solid fa-circle-xmark"></i>
                    <span>Maximum 200 characters allowed</span>
                </div>
            </div>
        </div>
    </div>

    <div class="content-card">
        <div class="card-header">
            <div class="card-header-left">
                <div class="card-icon bg-success">
                    <i class="fa-solid fa-hashtag"></i>
                </div>
                <div>
                    <h3>Number Range</h3>
                    <p class="card-subtitle">Min/max value validation</p>
                </div>
            </div>
        </div>
        <div class="card-body">
            <div class="form-group">
                <label>Age (18-100)</label>
                <input type="number" class="form-control is-valid" value="25" min="18" max="100">
                <div class="validation-feedback valid-feedback">
                    <i class="fa-solid fa-circle-check"></i>
                    <span>Valid age</span>
                </div>
            </div>

            <div class="form-group">
                <label>Age (Under 18)</label>
                <input type="number" class="form-control is-invalid" value="15" min="18" max="100">
                <div class="validation-feedback invalid-feedback">
                    <i class="fa-solid fa-circle-xmark"></i>
                    <span>You must be at least 18 years old</span>
                </div>
            </div>

            <div class="form-group">
                <label>Quantity (1-100)</label>
                <input type="number" class="form-control" value="50" min="1" max="100" step="1">
                <div class="validation-feedback valid-feedback">
                    <i class="fa-solid fa-circle-check"></i>
                    <span>In stock</span>
                </div>
            </div>
        </div>
    </div>

    <div class="content-card">
        <div class="card-header">
            <div class="card-header-left">
                <div class="card-icon bg-info">
                    <i class="fa-solid fa-calendar"></i>
                </div>
                <div>
                    <h3>Date Validation</h3>
                    <p class="card-subtitle">Date range and format</p>
                </div>
            </div>
        </div>
        <div class="card-body">
            <div class="form-group">
                <label>Birth Date</label>
                <input type="date" class="form-control is-valid" value="1990-05-15" max="2006-01-01">
                <div class="validation-feedback valid-feedback">
                    <i class="fa-solid fa-circle-check"></i>
                    <span>Valid date (33 years old)</span>
                </div>
            </div>

            <div class="form-group">
                <label>Future Date</label>
                <input type="date" class="form-control is-invalid" value="2030-01-01" max="2026-12-31">
                <div class="validation-feedback invalid-feedback">
                    <i class="fa-solid fa-circle-xmark"></i>
                    <span>Date cannot be in the future</span>
                </div>
            </div>

            <div class="form-group">
                <label>Event Date</label>
                <input type="date" class="form-control" value="2026-06-15" min="2026-01-01">
                <div class="validation-feedback valid-feedback">
                    <i class="fa-solid fa-circle-check"></i>
                    <span>Upcoming event</span>
                </div>
            </div>
        </div>
    </div>
</div>

<!-- File Upload Validation -->
<div class="validation-section-title">
    <i class="fa-solid fa-file-arrow-up"></i>
    File Upload Validation
</div>

<div class="validation-demo-grid">
    <div class="content-card">
        <div class="card-header">
            <div class="card-header-left">
                <div class="card-icon bg-primary">
                    <i class="fa-solid fa-image"></i>
                </div>
                <div>
                    <h3>Image Upload</h3>
                    <p class="card-subtitle">File type and size validation</p>
                </div>
            </div>
        </div>
        <div class="card-body">
            <div class="form-group">
                <label>Profile Photo <span style="color: var(--danger);">*</span></label>
                <input type="file" class="form-control" accept="image/jpeg,image/png,image/gif" max="2097152">
                <div class="validation-feedback valid-feedback">
                    <i class="fa-solid fa-circle-info"></i>
                    <span>Accepted: JPG, PNG, GIF (Max 2MB)</span>
                </div>
            </div>

            <div class="form-group">
                <label>Invalid File Type</label>
                <input type="file" class="form-control is-invalid" accept="image/*">
                <div class="validation-feedback invalid-feedback">
                    <i class="fa-solid fa-circle-xmark"></i>
                    <span>Only image files are allowed</span>
                </div>
            </div>

            <div class="divider"></div>

            <div class="form-group">
                <h4 style="font-size: 14px; margin-bottom: 12px;">File Validation Rules:</h4>
                <ul class="feature-list">
                    <li>
                        <i class="fa-solid fa-circle-check"></i>
                        <span>File type restriction (MIME type)</span>
                    </li>
                    <li>
                        <i class="fa-solid fa-circle-check"></i>
                        <span>Maximum file size limit</span>
                    </li>
                    <li>
                        <i class="fa-solid fa-circle-check"></i>
                        <span>Image dimension validation</span>
                    </li>
                    <li>
                        <i class="fa-solid fa-circle-check"></i>
                        <span>Client-side preview before upload</span>
                    </li>
                </ul>
            </div>
        </div>
    </div>

    <div class="content-card">
        <div class="card-header">
            <div class="card-header-left">
                <div class="card-icon bg-success">
                    <i class="fa-solid fa-file-pdf"></i>
                </div>
                <div>
                    <h3>Document Upload</h3>
                    <p class="card-subtitle">PDF and document validation</p>
                </div>
            </div>
        </div>
        <div class="card-body">
            <div class="form-group">
                <label>Resume (PDF)</label>
                <input type="file" class="form-control is-valid" accept=".pdf,application/pdf">
                <div class="validation-feedback valid-feedback">
                    <i class="fa-solid fa-circle-check"></i>
                    <span>Valid PDF document</span>
                </div>
            </div>

            <div class="form-group">
                <label>Invalid File</label>
                <input type="file" class="form-control is-invalid" accept=".pdf">
                <div class="validation-feedback invalid-feedback">
                    <i class="fa-solid fa-circle-xmark"></i>
                    <span>File must be in PDF format</span>
                </div>
            </div>

            <div class="divider"></div>

            <div class="code-block">
                <div style="color: var(--text-tertiary); margin-bottom: 8px;">HTML Validation:</div>
                <code style="color: var(--accent);">
                    &lt;input type="file" <br>
                    &nbsp;&nbsp;accept=".pdf,application/pdf"<br>
                    &nbsp;&nbsp;max-size="5242880"&gt;<br>
                    <br>
                    <span style="color: var(--text-tertiary);">// 5MB = 5242880 bytes</span>
                </code>
            </div>
        </div>
    </div>
</div>

<!-- Custom Validation -->
<div class="validation-section-title">
    <i class="fa-solid fa-wand-magic-sparkles"></i>
    Custom Validation Rules
</div>

<div class="validation-demo-grid full-width">
    <div class="content-card">
        <div class="card-header">
            <div class="card-header-left">
                <div class="card-icon bg-info">
                    <i class="fa-solid fa-code"></i>
                </div>
                <div>
                    <h3>Custom JavaScript Validation</h3>
                    <p class="card-subtitle">Complex validation with custom logic</p>
                </div>
            </div>
        </div>
        <div class="card-body">
            <div class="validation-demo-grid three-cols">
                <div>
                    <div class="form-group">
                        <label>Credit Card Number</label>
                        <input type="text" class="form-control is-valid" value="4111 1111 1111 1111" maxlength="19" pattern="\\d{4} \\d{4} \\d{4} \\d{4}">
                        <div class="validation-feedback valid-feedback">
                            <i class="fa-solid fa-circle-check"></i>
                            <span>Valid card number</span>
                        </div>
                    </div>
                </div>

                <div>
                    <div class="form-group">
                        <label>ZIP/Postal Code</label>
                        <input type="text" class="form-control is-valid" value="12345" maxlength="5" pattern="\\d{5}">
                        <div class="validation-feedback valid-feedback">
                            <i class="fa-solid fa-circle-check"></i>
                            <span>Valid ZIP code</span>
                        </div>
                    </div>
                </div>

                <div>
                    <div class="form-group">
                        <label>SSN/Tax ID</label>
                        <input type="text" class="form-control is-invalid" value="123-45" maxlength="11" pattern="\\d{3}-\\d{2}-\\d{4}">
                        <div class="validation-feedback invalid-feedback">
                            <i class="fa-solid fa-circle-xmark"></i>
                            <span>Format: XXX-XX-XXXX</span>
                        </div>
                    </div>
                </div>
            </div>

            <div class="divider"></div>

            <div class="form-group">
                <h4 style="font-size: 14px; margin-bottom: 12px;">Validation Capabilities:</h4>
                <ul class="feature-list">
                    <li>
                        <i class="fa-solid fa-circle-check"></i>
                        <span><strong>HTML5 Validation:</strong> Required, pattern, min/max, type</span>
                    </li>
                    <li>
                        <i class="fa-solid fa-circle-check"></i>
                        <span><strong>JavaScript Validation:</strong> Custom logic, async validation</span>
                    </li>
                    <li>
                        <i class="fa-solid fa-circle-check"></i>
                        <span><strong>Server Validation:</strong> Laravel validation rules</span>
                    </li>
                    <li>
                        <i class="fa-solid fa-circle-check"></i>
                        <span><strong>Real-time Feedback:</strong> Instant validation on input</span>
                    </li>
                    <li>
                        <i class="fa-solid fa-circle-check"></i>
                        <span><strong>Password Strength:</strong> Entropy-based strength meter</span>
                    </li>
                </ul>
            </div>
        </div>
    </div>
</div>

<!-- Complete Validation Form -->
<div class="validation-section-title">
    <i class="fa-solid fa-file-shield"></i>
    Complete Validation Form
</div>

<div class="content-card">
    <div class="card-header">
        <div class="card-header-left">
            <div class="card-icon bg-primary">
                <i class="fa-solid fa-user-plus"></i>
            </div>
            <div>
                <h3>Registration Form with Validation</h3>
                <p class="card-subtitle">Full form example with comprehensive validation</p>
            </div>
        </div>
    </div>
    <div class="card-body">
        <form id="registrationForm" style="max-width: 800px;">
            <div style="display: grid; grid-template-columns: 1fr 1fr; gap: 20px;">
                <div class="form-group">
                    <label>First Name <span style="color: var(--danger);">*</span></label>
                    <input type="text" class="form-control is-invalid" placeholder="John" required minlength="2">
                    <div class="validation-feedback invalid-feedback">
                        <i class="fa-solid fa-circle-xmark"></i>
                        <span>First name is required (min 2 characters)</span>
                    </div>
                </div>
                <div class="form-group">
                    <label>Last Name <span style="color: var(--danger);">*</span></label>
                    <input type="text" class="form-control is-valid" value="Doe" required minlength="2">
                    <div class="validation-feedback valid-feedback">
                        <i class="fa-solid fa-circle-check"></i>
                        <span>Looks good!</span>
                    </div>
                </div>
            </div>

            <div class="form-group">
                <label>Email Address <span style="color: var(--danger);">*</span></label>
                <input type="email" class="form-control is-invalid" value="invalid-email" required pattern="[a-z0-9._%+-]+@[a-z0-9.-]+\.[a-z]{2,}$">
                <div class="validation-feedback invalid-feedback">
                    <i class="fa-solid fa-circle-xmark"></i>
                    <span>Please enter a valid email address</span>
                </div>
            </div>

            <div class="form-group">
                <label>Password <span style="color: var(--danger);">*</span></label>
                <input type="password" class="form-control is-invalid" value="12345" required minlength="8">
                
                <div class="password-strength">
                    <div class="strength-meter">
                        <div class="strength-meter-fill strength-weak"></div>
                    </div>
                    <div class="strength-text" style="color: var(--danger);">Weak - Add uppercase, numbers, and special characters</div>
                </div>

                <ul class="requirement-list">
                    <li class="unmet">
                        <i class="fa-solid fa-circle-xmark"></i>
                        <span>Minimum 8 characters (5/8)</span>
                    </li>
                    <li class="unmet">
                        <i class="fa-solid fa-circle-xmark"></i>
                        <span>At least one uppercase letter</span>
                    </li>
                    <li class="met">
                        <i class="fa-solid fa-circle-check"></i>
                        <span>At least one number</span>
                    </li>
                    <li class="unmet">
                        <i class="fa-solid fa-circle-xmark"></i>
                        <span>At least one special character</span>
                    </li>
                </ul>
            </div>

            <div class="form-group">
                <label>Confirm Password <span style="color: var(--danger);">*</span></label>
                <input type="password" class="form-control is-invalid" value="different" required>
                <div class="validation-feedback invalid-feedback">
                    <i class="fa-solid fa-circle-xmark"></i>
                    <span>Passwords do not match</span>
                </div>
            </div>

            <div class="form-group">
                <label>Phone Number</label>
                <input type="tel" class="form-control is-invalid" value="0812" pattern="^\\+62\\d{8,12}$">
                <div class="validation-feedback invalid-feedback">
                    <i class="fa-solid fa-circle-xmark"></i>
                    <span>Format: +62xxxxxxxxxx (10-13 digits)</span>
                </div>
            </div>

            <div class="form-group">
                <label>Agreement <span style="color: var(--danger);">*</span></label>
                <label style="display: flex; align-items: center; gap: 8px; cursor: pointer; padding: 8px 0;">
                    <input type="checkbox" required>
                    <span>I agree to the <a href="#" style="color: var(--accent);">Terms of Service</a> and <a href="#" style="color: var(--accent);">Privacy Policy</a></span>
                </label>
            </div>

            <div style="display: flex; gap: 12px; margin-top: 24px;">
                <button type="submit" class="btn btn-primary">
                    <i class="fa-solid fa-user-plus"></i>
                    Create Account
                </button>
                <button type="reset" class="btn btn-secondary">
                    <i class="fa-solid fa-rotate-left"></i>
                    Reset Form
                </button>
                <button type="button" class="btn btn-outline-info">
                    <i class="fa-solid fa-circle-info"></i>
                    View Rules
                </button>
            </div>
        </form>
    </div>
</div>
@endsection
