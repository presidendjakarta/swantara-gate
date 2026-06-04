@extends('layouts.app')

@section('title', 'Floating Labels')

@push('styles')
@endpush

@push('scripts')
<!-- IMask.js Library -->
<script src="https://unpkg.com/imask@7.1.3/dist/imask.min.js"></script>
<script>
    document.addEventListener('DOMContentLoaded', function() {
        // Initialize input masks
        
        // Phone number mask (Indonesian format)
        document.querySelectorAll('[data-mask="phone"]').forEach(input => {
            IMask(input, {
                mask: '+62 {000}-{0000}-{0000}'
            });
        });

        // Credit card mask
        document.querySelectorAll('[data-mask="credit-card"]').forEach(input => {
            IMask(input, {
                mask: '0000 0000 0000 0000'
            });
        });

        // Expiry date mask (MM/YY)
        document.querySelectorAll('[data-mask="expiry-date"]').forEach(input => {
            IMask(input, {
                mask: 'MM{/}YY',
                blocks: {
                    MM: {
                        mask: IMask.MaskedRange,
                        from: 1,
                        to: 12,
                        maxLength: 2
                    },
                    YY: {
                        mask: IMask.MaskedRange,
                        from: 0,
                        to: 99,
                        maxLength: 2
                    }
                }
            });
        });

        // CVV mask (3-4 digits)
        document.querySelectorAll('[data-mask="cvv"]').forEach(input => {
            IMask(input, {
                mask: '0009',
                lazy: false
            });
        });

        // Handle select elements - add 'has-value' class when changed
        document.querySelectorAll('.floating-label-group select').forEach(select => {
            // Add has-value class if option is selected
            if (select.value) {
                select.classList.add('has-value');
            }
            
            // Listen for changes
            select.addEventListener('change', function() {
                if (this.value) {
                    this.classList.add('has-value');
                } else {
                    this.classList.remove('has-value');
                }
            });
            
            // Trigger change event to ensure initial state
            select.dispatchEvent(new Event('change'));
        });

        // Handle form reset
        const forms = document.querySelectorAll('form');
        forms.forEach(form => {
            form.addEventListener('reset', function() {
                setTimeout(() => {
                    this.querySelectorAll('select').forEach(select => {
                        select.classList.remove('has-value');
                    });
                }, 0);
            });
        });
    });
</script>
@endpush

@section('content')
<div class="page-header">
    <div>
        <h1>Floating Labels</h1>
        <p>Modern label animation yang mengambang saat input terisi</p>
    </div>
    <div style="display: flex; gap: 10px;">
        <button class="btn btn-secondary">
            <i class="fa-solid fa-book"></i>
            <span>Documentation</span>
        </button>
        <button class="btn btn-primary">
            <i class="fa-solid fa-layer-group"></i>
            <span>View Examples</span>
        </button>
    </div>
</div>

<!-- Info Alert -->
<div class="content-card" style="margin-bottom: 24px; border-left: 4px solid var(--accent);">
    <div class="card-body" style="padding: 16px 20px;">
        <div style="display: flex; gap: 12px; align-items: start;">
            <i class="fa-solid fa-circle-info" style="color: var(--accent); font-size: 20px; margin-top: 2px;"></i>
            <div style="flex: 1;">
                <h4 style="margin-bottom: 4px; font-size: 14px;">How Floating Labels Work</h4>
                <p style="font-size: 13px; color: var(--text-secondary); margin: 0;">Labels start inside the input field and animate to the top border when the user focuses or enters text. This saves space while maintaining clarity. Uses CSS transitions for smooth animation.</p>
            </div>
        </div>
    </div>
</div>

<!-- Basic Floating Labels -->
<div class="fl-section-title">
    <i class="fa-solid fa-arrow-up"></i>
    Basic Floating Labels <span class="badge badge-primary">Essential</span>
</div>

<div class="floating-labels-grid">
    <div class="content-card">
        <div class="card-header">
            <div class="card-header-left">
                <div class="card-icon bg-primary">
                    <i class="fa-solid fa-font"></i>
                </div>
                <div>
                    <h3>Text Inputs</h3>
                    <p class="card-subtitle">Basic floating label behavior</p>
                </div>
            </div>
        </div>
        <div class="card-body">
            <div class="floating-label-group">
                <input type="text" id="fname" placeholder=" " value="John">
                <label for="fname">First Name</label>
            </div>

            <div class="floating-label-group">
                <input type="text" id="lname" placeholder=" ">
                <label for="lname">Last Name</label>
            </div>

            <div class="floating-label-group">
                <input type="email" id="email" placeholder=" " value="john@example.com">
                <label for="email">Email Address</label>
            </div>

            <div class="floating-label-group">
                <input type="tel" id="phone" placeholder=" " data-mask="phone">
                <label for="phone">Phone Number</label>
            </div>
        </div>
    </div>

    <div class="content-card">
        <div class="card-header">
            <div class="card-header-left">
                <div class="card-icon bg-success">
                    <i class="fa-solid fa-lock"></i>
                </div>
                <div>
                    <h3>Password & Search</h3>
                    <p class="card-subtitle">Secure and search inputs</p>
                </div>
            </div>
        </div>
        <div class="card-body">
            <div class="floating-label-group">
                <input type="password" id="password" placeholder=" ">
                <label for="password">Password</label>
            </div>

            <div class="floating-label-group">
                <input type="password" id="confirm" placeholder=" " value="••••••••">
                <label for="confirm">Confirm Password</label>
            </div>

            <div class="floating-label-group">
                <input type="search" id="search" placeholder=" ">
                <label for="search">Search</label>
            </div>

            <div class="floating-label-group">
                <input type="url" id="website" placeholder=" ">
                <label for="website">Website URL</label>
            </div>
        </div>
    </div>
</div>

<!-- With Icons -->
<div class="fl-section-title">
    <i class="fa-solid fa-icons"></i>
    Floating Labels with Icons <span class="badge badge-success">Popular</span>
</div>

<div class="floating-labels-grid">
    <div class="content-card">
        <div class="card-header">
            <div class="card-header-left">
                <div class="card-icon bg-info">
                    <i class="fa-solid fa-user"></i>
                </div>
                <div>
                    <h3>User Information</h3>
                    <p class="card-subtitle">Icons for better UX</p>
                </div>
            </div>
        </div>
        <div class="card-body">
            <div class="floating-label-group floating-label-with-icon">
                <i class="fa-solid fa-user"></i>
                <input type="text" id="username" placeholder=" " value="johndoe">
                <label for="username">Username</label>
            </div>

            <div class="floating-label-group floating-label-with-icon">
                <i class="fa-solid fa-envelope"></i>
                <input type="email" id="uemail" placeholder=" ">
                <label for="uemail">Email Address</label>
            </div>

            <div class="floating-label-group floating-label-with-icon">
                <i class="fa-solid fa-phone"></i>
                <input type="tel" id="uphone" placeholder=" " data-mask="phone">
                <label for="uphone">Phone Number</label>
            </div>

            <div class="floating-label-group floating-label-with-icon">
                <i class="fa-solid fa-globe"></i>
                <input type="url" id="uwebsite" placeholder=" ">
                <label for="uwebsite">Website</label>
            </div>
        </div>
    </div>

    <div class="content-card">
        <div class="card-header">
            <div class="card-header-left">
                <div class="card-icon bg-warning">
                    <i class="fa-solid fa-credit-card"></i>
                </div>
                <div>
                    <h3>Payment Information</h3>
                    <p class="card-subtitle">Financial form icons</p>
                </div>
            </div>
        </div>
        <div class="card-body">
            <div class="floating-label-group floating-label-with-icon">
                <i class="fa-solid fa-credit-card"></i>
                <input type="text" id="card" placeholder=" " value="4111 1111 1111 1111" data-mask="credit-card">
                <label for="card">Card Number</label>
            </div>

            <div class="floating-label-group floating-label-with-icon">
                <i class="fa-solid fa-user-pen"></i>
                <input type="text" id="cname" placeholder=" ">
                <label for="cname">Cardholder Name</label>
            </div>

            <div class="floating-label-group floating-label-with-icon">
                <i class="fa-solid fa-calendar"></i>
                <input type="text" id="expiry" placeholder=" " value="12/25" data-mask="expiry-date">
                <label for="expiry">Expiry Date</label>
            </div>

            <div class="floating-label-group floating-label-with-icon">
                <i class="fa-solid fa-lock"></i>
                <input type="text" id="cvv" placeholder=" " data-mask="cvv" maxlength="4">
                <label for="cvv">CVV</label>
            </div>
        </div>
    </div>
</div>

<!-- Textarea & Select -->
<div class="fl-section-title">
    <i class="fa-solid fa-align-left"></i>
    Textarea & Select
</div>

<div class="floating-labels-grid">
    <div class="content-card">
        <div class="card-header">
            <div class="card-header-left">
                <div class="card-icon bg-primary">
                    <i class="fa-solid fa-text-width"></i>
                </div>
                <div>
                    <h3>Textarea</h3>
                    <p class="card-subtitle">Multi-line floating labels</p>
                </div>
            </div>
        </div>
        <div class="card-body">
            <div class="floating-label-group">
                <textarea id="bio" rows="3" placeholder=" "></textarea>
                <label for="bio">Biography</label>
            </div>

            <div class="floating-label-group">
                <textarea id="message" rows="4" placeholder=" ">Hello, I would like to...</textarea>
                <label for="message">Message</label>
            </div>

            <div class="floating-label-group">
                <textarea id="address" rows="3" placeholder=" "></textarea>
                <label for="address">Shipping Address</label>
            </div>

            <div class="divider"></div>

            <div class="form-group">
                <h4 style="font-size: 14px; margin-bottom: 12px;">Features:</h4>
                <ul class="feature-list">
                    <li>
                        <i class="fa-solid fa-circle-check"></i>
                        <span>Label stays at top when filled</span>
                    </li>
                    <li>
                        <i class="fa-solid fa-circle-check"></i>
                        <span>Smooth animation transition</span>
                    </li>
                    <li>
                        <i class="fa-solid fa-circle-check"></i>
                        <span>Works with any rows attribute</span>
                    </li>
                    <li>
                        <i class="fa-solid fa-circle-check"></i>
                        <span>Resize handle compatible</span>
                    </li>
                </ul>
            </div>
        </div>
    </div>

    <div class="content-card">
        <div class="card-header">
            <div class="card-header-left">
                <div class="card-icon bg-success">
                    <i class="fa-solid fa-list"></i>
                </div>
                <div>
                    <h3>Select Dropdown</h3>
                    <p class="card-subtitle">Floating labels for selects</p>
                </div>
            </div>
        </div>
        <div class="card-body">
            <div class="floating-label-group">
                <select id="country" class="has-value">
                    <option value="">Select Country</option>
                    <option value="id" selected>🇮🇩 Indonesia</option>
                    <option value="us">🇺🇸 United States</option>
                    <option value="uk">🇬🇧 United Kingdom</option>
                    <option value="jp">🇯🇵 Japan</option>
                </select>
                <label for="country">Country</label>
            </div>

            <div class="floating-label-group">
                <select id="language">
                    <option value="">Select Language</option>
                    <option value="en">English</option>
                    <option value="id">Bahasa Indonesia</option>
                    <option value="jp">日本語</option>
                </select>
                <label for="language">Language</label>
            </div>

            <div class="floating-label-group">
                <select id="timezone" class="has-value">
                    <option value="">Select Timezone</option>
                    <option value="wib" selected>WIB (UTC+7)</option>
                    <option value="wita">WITA (UTC+8)</option>
                    <option value="wit">WIT (UTC+9)</option>
                </select>
                <label for="timezone">Timezone</label>
            </div>

            <div class="divider"></div>

            <div class="helper-text">
                <i class="fa-solid fa-circle-info"></i>
                Add class "has-value" to pre-selected options
            </div>
        </div>
    </div>
</div>

<!-- Validation States -->
<div class="fl-section-title">
    <i class="fa-solid fa-circle-check"></i>
    Validation States
</div>

<div class="floating-labels-grid">
    <div class="content-card">
        <div class="card-header">
            <div class="card-header-left">
                <div class="card-icon bg-success">
                    <i class="fa-solid fa-check"></i>
                </div>
                <div>
                    <h3>Valid State</h3>
                    <p class="card-subtitle">Green border and label</p>
                </div>
            </div>
        </div>
        <div class="card-body">
            <div class="floating-label-group">
                <input type="text" id="valid-name" class="is-valid" placeholder=" " value="John Doe">
                <label for="valid-name">Full Name</label>
            </div>
            <div class="helper-text success-text">
                <i class="fa-solid fa-circle-check"></i>
                <span>Looks good!</span>
            </div>

            <div class="floating-label-group" style="margin-top: 20px;">
                <input type="email" id="valid-email" class="is-valid" placeholder=" " value="john@example.com">
                <label for="valid-email">Email Address</label>
            </div>
            <div class="helper-text success-text">
                <i class="fa-solid fa-circle-check"></i>
                <span>Valid email format</span>
            </div>

            <div class="floating-label-group" style="margin-top: 20px;">
                <input type="tel" id="valid-phone" class="is-valid" placeholder=" " value="+62 812-3456-7890" data-mask="phone">
                <label for="valid-phone">Phone Number</label>
            </div>
            <div class="helper-text success-text">
                <i class="fa-solid fa-circle-check"></i>
                <span>Valid phone number</span>
            </div>
        </div>
    </div>

    <div class="content-card">
        <div class="card-header">
            <div class="card-header-left">
                <div class="card-icon bg-danger">
                    <i class="fa-solid fa-circle-xmark"></i>
                </div>
                <div>
                    <h3>Invalid State</h3>
                    <p class="card-subtitle">Red border and error message</p>
                </div>
            </div>
        </div>
        <div class="card-body">
            <div class="floating-label-group">
                <input type="text" id="invalid-name" class="is-invalid" placeholder=" " value="">
                <label for="invalid-name">Full Name</label>
            </div>
            <div class="helper-text error-text">
                <i class="fa-solid fa-circle-exclamation"></i>
                <span>Name is required</span>
            </div>

            <div class="floating-label-group" style="margin-top: 20px;">
                <input type="email" id="invalid-email" class="is-invalid" placeholder=" " value="invalid-email">
                <label for="invalid-email">Email Address</label>
            </div>
            <div class="helper-text error-text">
                <i class="fa-solid fa-circle-exclamation"></i>
                <span>Please enter a valid email</span>
            </div>

            <div class="floating-label-group" style="margin-top: 20px;">
                <input type="password" id="invalid-pass" class="is-invalid" placeholder=" " value="123">
                <label for="invalid-pass">Password</label>
            </div>
            <div class="helper-text error-text">
                <i class="fa-solid fa-circle-exclamation"></i>
                <span>Minimum 8 characters required</span>
            </div>
        </div>
    </div>
</div>

<!-- Disabled & Readonly -->
<div class="fl-section-title">
    <i class="fa-solid fa-ban"></i>
    Disabled & Read-only States
</div>

<div class="floating-labels-grid">
    <div class="content-card">
        <div class="card-header">
            <div class="card-header-left">
                <div class="card-icon bg-warning">
                    <i class="fa-solid fa-lock"></i>
                </div>
                <div>
                    <h3>Disabled Inputs</h3>
                    <p class="card-subtitle">Non-interactive fields</p>
                </div>
            </div>
        </div>
        <div class="card-body">
            <div class="floating-label-group">
                <input type="text" id="disabled1" placeholder=" " value="Cannot edit this" disabled>
                <label for="disabled1">Disabled Field</label>
            </div>

            <div class="floating-label-group">
                <input type="email" id="disabled2" placeholder=" " value="user@example.com" disabled>
                <label for="disabled2">Disabled Email</label>
            </div>

            <div class="floating-label-group">
                <input type="text" id="disabled3" placeholder=" " disabled>
                <label for="disabled3">Empty Disabled</label>
            </div>

            <div class="divider"></div>

            <div class="helper-text">
                <i class="fa-solid fa-circle-info"></i>
                Disabled inputs have 50% opacity and cannot be focused
            </div>
        </div>
    </div>

    <div class="content-card">
        <div class="card-header">
            <div class="card-header-left">
                <div class="card-icon bg-info">
                    <i class="fa-solid fa-eye"></i>
                </div>
                <div>
                    <h3>Read-only Inputs</h3>
                    <p class="card-subtitle">View-only fields</p>
                </div>
            </div>
        </div>
        <div class="card-body">
            <div class="floating-label-group">
                <input type="text" id="readonly1" placeholder=" " value="User ID: 12345" readonly>
                <label for="readonly1">User ID</label>
            </div>

            <div class="floating-label-group">
                <input type="email" id="readonly2" placeholder=" " value="john@example.com" readonly>
                <label for="readonly2">Email (Verified)</label>
            </div>

            <div class="floating-label-group">
                <input type="text" id="readonly3" placeholder=" " value="Created: 2024-01-15" readonly>
                <label for="readonly3">Created Date</label>
            </div>

            <div class="divider"></div>

            <div class="helper-text">
                <i class="fa-solid fa-circle-info"></i>
                Read-only inputs can be focused but not edited
            </div>
        </div>
    </div>
</div>

<!-- Complete Registration Form -->
<div class="fl-section-title">
    <i class="fa-solid fa-clipboard-list"></i>
    Complete Registration Form
</div>

<div class="floating-labels-grid full-width">
    <div class="content-card">
        <div class="card-header">
            <div class="card-header-left">
                <div class="card-icon bg-primary">
                    <i class="fa-solid fa-user-plus"></i>
                </div>
                <div>
                    <h3>User Registration</h3>
                    <p class="card-subtitle">Complete form with floating labels</p>
                </div>
            </div>
        </div>
        <div class="card-body">
            <form style="max-width: 800px;">
                <div style="display: grid; grid-template-columns: 1fr 1fr; gap: 20px;">
                    <div class="floating-label-group">
                        <input type="text" id="reg-fname" placeholder=" " required>
                        <label for="reg-fname">First Name <span style="color: var(--danger);">*</span></label>
                    </div>

                    <div class="floating-label-group">
                        <input type="text" id="reg-lname" placeholder=" " required>
                        <label for="reg-lname">Last Name <span style="color: var(--danger);">*</span></label>
                    </div>
                </div>

                <div class="floating-label-group">
                    <input type="email" id="reg-email" placeholder=" " required>
                    <label for="reg-email">Email Address <span style="color: var(--danger);">*</span></label>
                </div>

                <div class="floating-label-group">
                    <input type="tel" id="reg-phone" placeholder=" " data-mask="phone">
                    <label for="reg-phone">Phone Number</label>
                </div>

                <div class="floating-label-group">
                    <input type="password" id="reg-password" placeholder=" " required>
                    <label for="reg-password">Password <span style="color: var(--danger);">*</span></label>
                </div>

                <div class="floating-label-group">
                    <input type="password" id="reg-confirm" placeholder=" " required>
                    <label for="reg-confirm">Confirm Password <span style="color: var(--danger);">*</span></label>
                </div>

                <div class="floating-label-group">
                    <textarea id="reg-bio" rows="3" placeholder=" "></textarea>
                    <label for="reg-bio">Biography (Optional)</label>
                </div>

                <div class="floating-label-group">
                    <select id="reg-country">
                        <option value="">Select Country</option>
                        <option value="id">🇮🇩 Indonesia</option>
                        <option value="us">🇺🇸 United States</option>
                        <option value="uk">🇬🇧 United Kingdom</option>
                    </select>
                    <label for="reg-country">Country</label>
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
                </div>
            </form>
        </div>
    </div>
</div>

<!-- Comparison & Best Practices -->
<div class="floating-labels-grid full-width">
    <div class="content-card">
        <div class="card-header">
            <div class="card-header-left">
                <div class="card-icon bg-info">
                    <i class="fa-solid fa-lightbulb"></i>
                </div>
                <div>
                    <h3>Best Practices</h3>
                    <p class="card-subtitle">When and how to use floating labels</p>
                </div>
            </div>
        </div>
        <div class="card-body">
            <div class="floating-labels-grid three-cols">
                <div>
                    <h4 style="font-size: 14px; margin-bottom: 12px; display: flex; align-items: center; gap: 8px;">
                        <span style="font-size: 24px;">✅</span>
                        When to Use
                    </h4>
                    <ul class="feature-list">
                        <li><i class="fa-solid fa-check"></i> Space-constrained forms</li>
                        <li><i class="fa-solid fa-check"></i> Modern UI design</li>
                        <li><i class="fa-solid fa-check"></i> Fewer form fields</li>
                        <li><i class="fa-solid fa-check"></i> Simple data entry</li>
                        <li><i class="fa-solid fa-check"></i> Mobile-first design</li>
                    </ul>
                </div>
                <div>
                    <h4 style="font-size: 14px; margin-bottom: 12px; display: flex; align-items: center; gap: 8px;">
                        <span style="font-size: 24px;">❌</span>
                        When to Avoid
                    </h4>
                    <ul class="feature-list">
                        <li><i class="fa-solid fa-xmark"></i> Complex multi-field forms</li>
                        <li><i class="fa-solid fa-xmark"></i> Users need to scan quickly</li>
                        <li><i class="fa-solid fa-xmark"></i> Accessibility priority</li>
                        <li><i class="fa-solid fa-xmark"></i> Long label text</li>
                        <li><i class="fa-solid fa-xmark"></i> Elderly users</li>
                    </ul>
                </div>
                <div>
                    <h4 style="font-size: 14px; margin-bottom: 12px; display: flex; align-items: center; gap: 8px;">
                        <span style="font-size: 24px;">💡</span>
                        Tips
                    </h4>
                    <ul class="feature-list">
                        <li><i class="fa-solid fa-check"></i> Use placeholder=" " trick</li>
                        <li><i class="fa-solid fa-check"></i> Keep labels short</li>
                        <li><i class="fa-solid fa-check"></i> Ensure contrast</li>
                        <li><i class="fa-solid fa-check"></i> Test on mobile</li>
                        <li><i class="fa-solid fa-check"></i> Add validation states</li>
                    </ul>
                </div>
            </div>
        </div>
    </div>
</div>
@endsection
