@extends('layouts.app')

@section('title', 'Input Masks')

@push('styles')
@endpush

@push('scripts')
<!-- IMask.js Library -->
<script src="https://unpkg.com/imask@7.1.3/dist/imask.min.js"></script>
<script src="{{ asset('assets/js/input-mask.js') }}"></script>
@endpush

@section('content')
<div class="page-header">
    <div>
        <h1>Input Masks</h1>
        <p>Auto-formatting input fields dengan pattern masking untuk data validation</p>
    </div>
    <div style="display: flex; gap: 10px;">
        <button class="btn btn-secondary">
            <i class="fa-solid fa-book"></i>
            <span>Documentation</span>
        </button>
        <button class="btn btn-primary">
            <i class="fa-solid fa-wand-magic-sparkles"></i>
            <span>Try Examples</span>
        </button>
    </div>
</div>

<!-- Info Alert -->
<div class="content-card" style="margin-bottom: 24px; border-left: 4px solid var(--warning);">
    <div class="card-body" style="padding: 16px 20px;">
        <div style="display: flex; gap: 12px; align-items: start;">
            <i class="fa-solid fa-triangle-exclamation" style="color: var(--warning); font-size: 20px; margin-top: 2px;"></i>
            <div style="flex: 1;">
                <h4 style="margin-bottom: 4px; font-size: 14px;">Input Mask Library Required</h4>
                <p style="font-size: 13px; color: var(--text-secondary); margin: 0;">Input masks require JavaScript library like <strong>IMask.js</strong> or <strong>jQuery InputMask</strong>. Examples below show formatted outputs. Include the library in your project for live masking.</p>
            </div>
        </div>
    </div>
</div>

<!-- Phone Number Masks -->
<div class="im-section-title">
    <i class="fa-solid fa-phone"></i>
    Phone Number Masks <span class="badge badge-primary">Essential</span>
</div>

<div class="input-masks-grid">
    <div class="content-card">
        <div class="card-header">
            <div class="card-header-left">
                <div class="card-icon bg-primary">
                    <i class="fa-solid fa-mobile-screen"></i>
                </div>
                <div>
                    <h3>Indonesian Phone</h3>
                    <p class="card-subtitle">Local phone format masking</p>
                </div>
            </div>
        </div>
        <div class="card-body">
            <div class="mask-example">
                <label class="mask-label">
                    Mobile Phone
                    <span class="mask-pattern">+62 ###-####-####</span>
                </label>
                <div class="input-group">
                    <span class="input-group-text">
                        <i class="fa-solid fa-mobile-screen"></i>
                    </span>
                    <input type="text" class="form-control" placeholder="+62 812-3456-7890" data-mask="phone">
                </div>
                <div class="mask-placeholder">
                    <i class="fa-solid fa-circle-info"></i>
                    Format: <code>+62 812-3456-7890</code>
                </div>
            </div>

            <div class="mask-example">
                <label class="mask-label">
                    Landline
                    <span class="mask-pattern">(0##) ###-####</span>
                </label>
                <div class="input-group">
                    <span class="input-group-text">
                        <i class="fa-solid fa-phone"></i>
                    </span>
                    <input type="text" class="form-control" placeholder="(021) 123-4567" data-mask="landline">
                </div>
                <div class="mask-placeholder">
                    <i class="fa-solid fa-circle-info"></i>
                    Format: <code>(021) 123-4567</code>
                </div>
            </div>

            <div class="mask-example">
                <label class="mask-label">
                    Toll Free
                    <span class="mask-pattern">###-###-####</span>
                </label>
                <div class="input-group">
                    <span class="input-group-text">
                        <i class="fa-solid fa-headset"></i>
                    </span>
                    <input type="text" class="form-control" placeholder="800-123-4567" data-mask="toll-free">
                </div>
                <div class="mask-placeholder">
                    <i class="fa-solid fa-circle-info"></i>
                    Format: <code>800-123-4567</code>
                </div>
            </div>

            <div class="divider"></div>

            <div class="live-preview">
                <div class="live-preview-label">Live Preview:</div>
                <div class="live-preview-value">+62 812-3456-7890</div>
            </div>
        </div>
    </div>

    <div class="content-card">
        <div class="card-header">
            <div class="card-header-left">
                <div class="card-icon bg-success">
                    <i class="fa-solid fa-globe"></i>
                </div>
                <div>
                    <h3>International Phone</h3>
                    <p class="card-subtitle">Global phone formats</p>
                </div>
            </div>
        </div>
        <div class="card-body">
            <div class="mask-example">
                <label class="mask-label">
                    US Phone
                    <span class="mask-pattern">(###) ###-####</span>
                </label>
                <div class="input-group">
                    <span class="input-group-text">🇺🇸</span>
                    <input type="text" class="form-control" placeholder="(555) 123-4567" data-mask="phone-us">
                </div>
                <div class="mask-placeholder">
                    <i class="fa-solid fa-circle-info"></i>
                    Format: <code>(555) 123-4567</code>
                </div>
            </div>

            <div class="mask-example">
                <label class="mask-label">
                    UK Phone
                    <span class="mask-pattern">+## #### ######</span>
                </label>
                <div class="input-group">
                    <span class="input-group-text">🇬🇧</span>
                    <input type="text" class="form-control" placeholder="+44 20 7946 0958" data-mask="phone-uk">
                </div>
                <div class="mask-placeholder">
                    <i class="fa-solid fa-circle-info"></i>
                    Format: <code>+44 20 7946 0958</code>
                </div>
            </div>

            <div class="mask-example">
                <label class="mask-label">
                    Japan Phone
                    <span class="mask-pattern">+##-##-####-####</span>
                </label>
                <div class="input-group">
                    <span class="input-group-text">🇯🇵</span>
                    <input type="text" class="form-control" placeholder="+81-90-1234-5678" data-mask="phone-jp">
                </div>
                <div class="mask-placeholder">
                    <i class="fa-solid fa-circle-info"></i>
                    Format: <code>+81-90-1234-5678</code>
                </div>
            </div>

            <div class="divider"></div>

            <div class="form-group">
                <h4 style="font-size: 14px; margin-bottom: 12px;">Mask Features:</h4>
                <ul class="feature-list">
                    <li>
                        <i class="fa-solid fa-circle-check"></i>
                        <span>Auto-formatting on input</span>
                    </li>
                    <li>
                        <i class="fa-solid fa-circle-check"></i>
                        <span>Country code support</span>
                    </li>
                    <li>
                        <i class="fa-solid fa-circle-check"></i>
                        <span>Digit validation</span>
                    </li>
                    <li>
                        <i class="fa-solid fa-circle-check"></i>
                        <span>Visual separators</span>
                    </li>
                </ul>
            </div>
        </div>
    </div>
</div>

<!-- Date & Time Masks -->
<div class="im-section-title">
    <i class="fa-solid fa-calendar-days"></i>
    Date & Time Masks <span class="badge badge-success">Popular</span>
</div>

<div class="input-masks-grid">
    <div class="content-card">
        <div class="card-header">
            <div class="card-header-left">
                <div class="card-icon bg-info">
                    <i class="fa-solid fa-calendar"></i>
                </div>
                <div>
                    <h3>Date Formats</h3>
                    <p class="card-subtitle">Various date mask patterns</p>
                </div>
            </div>
        </div>
        <div class="card-body">
            <div class="mask-example">
                <label class="mask-label">
                    Date (DD/MM/YYYY)
                    <span class="mask-pattern">##/##/####</span>
                </label>
                <div class="input-group">
                    <span class="input-group-text">
                        <i class="fa-solid fa-calendar-days"></i>
                    </span>
                    <input type="text" class="form-control" placeholder="DD/MM/YYYY" data-mask="date-dmy">
                </div>
                <div class="mask-placeholder">
                    <i class="fa-solid fa-circle-info"></i>
                    Format: <code>DD/MM/YYYY</code>
                </div>
            </div>

            <div class="mask-example">
                <label class="mask-label">
                    Date (MM-DD-YYYY)
                    <span class="mask-pattern">##-##-####</span>
                </label>
                <div class="input-group">
                    <span class="input-group-text">
                        <i class="fa-solid fa-calendar-check"></i>
                    </span>
                    <input type="text" class="form-control" placeholder="MM-DD-YYYY" data-mask="date-mdy">
                </div>
                <div class="mask-placeholder">
                    <i class="fa-solid fa-circle-info"></i>
                    Format: <code>MM-DD-YYYY</code>
                </div>
            </div>

            <div class="mask-example">
                <label class="mask-label">
                    Date (YYYY.MM.DD)
                    <span class="mask-pattern">####.##.##</span>
                </label>
                <div class="input-group">
                    <span class="input-group-text">
                        <i class="fa-solid fa-calendar-week"></i>
                    </span>
                    <input type="text" class="form-control" placeholder="YYYY.MM.DD" data-mask="date-ymd">
                </div>
                <div class="mask-placeholder">
                    <i class="fa-solid fa-circle-info"></i>
                    Format: <code>YYYY.MM.DD</code>
                </div>
            </div>

            <div class="divider"></div>

            <div class="live-preview">
                <div class="live-preview-label">Current Date:</div>
                <div class="live-preview-value">31/05/2026</div>
            </div>
        </div>
    </div>

    <div class="content-card">
        <div class="card-header">
            <div class="card-header-left">
                <div class="card-icon bg-warning">
                    <i class="fa-solid fa-clock"></i>
                </div>
                <div>
                    <h3>Time Formats</h3>
                    <p class="card-subtitle">Time mask patterns</p>
                </div>
            </div>
        </div>
        <div class="card-body">
            <div class="mask-example">
                <label class="mask-label">
                    Time 24h (HH:MM)
                    <span class="mask-pattern">##:##</span>
                </label>
                <div class="input-group">
                    <span class="input-group-text">
                        <i class="fa-solid fa-clock"></i>
                    </span>
                    <input type="text" class="form-control" placeholder="HH:MM" data-mask="time-24h">
                </div>
                <div class="mask-placeholder">
                    <i class="fa-solid fa-circle-info"></i>
                    Format: <code>HH:MM</code> (24-hour)
                </div>
            </div>

            <div class="mask-example">
                <label class="mask-label">
                    Time 12h (HH:MM AM/PM)
                    <span class="mask-pattern">##:## ##</span>
                </label>
                <div class="input-group">
                    <span class="input-group-text">
                        <i class="fa-solid fa-sun"></i>
                    </span>
                    <input type="text" class="form-control" placeholder="HH:MM AM/PM" data-mask="time-12h">
                </div>
                <div class="mask-placeholder">
                    <i class="fa-solid fa-circle-info"></i>
                    Format: <code>HH:MM AM/PM</code>
                </div>
            </div>

            <div class="mask-example">
                <label class="mask-label">
                    DateTime (DD/MM/YYYY HH:MM)
                    <span class="mask-pattern">##/##/#### ##:##</span>
                </label>
                <div class="input-group">
                    <span class="input-group-text">
                        <i class="fa-solid fa-calendar-clock"></i>
                    </span>
                    <input type="text" class="form-control" placeholder="DD/MM/YYYY HH:MM" data-mask="datetime">
                </div>
                <div class="mask-placeholder">
                    <i class="fa-solid fa-circle-info"></i>
                    Format: <code>DD/MM/YYYY HH:MM</code>
                </div>
            </div>

            <div class="divider"></div>

            <div class="form-group">
                <h4 style="font-size: 14px; margin-bottom: 12px;">Validation Rules:</h4>
                <ul class="feature-list">
                    <li>
                        <i class="fa-solid fa-circle-check"></i>
                        <span>Day: 01-31</span>
                    </li>
                    <li>
                        <i class="fa-solid fa-circle-check"></i>
                        <span>Month: 01-12</span>
                    </li>
                    <li>
                        <i class="fa-solid fa-circle-check"></i>
                        <span>Hour: 00-23 (24h)</span>
                    </li>
                    <li>
                        <i class="fa-solid fa-circle-check"></i>
                        <span>Minute: 00-59</span>
                    </li>
                </ul>
            </div>
        </div>
    </div>
</div>

<!-- Number & Currency Masks -->
<div class="im-section-title">
    <i class="fa-solid fa-hashtag"></i>
    Number & Currency Masks
</div>

<div class="input-masks-grid three-cols">
    <div class="content-card">
        <div class="card-header">
            <div class="card-header-left">
                <div class="card-icon bg-primary">
                    <i class="fa-solid fa-sort-numeric-up"></i>
                </div>
                <div>
                    <h3>Number Format</h3>
                    <p class="card-subtitle">Numeric masking</p>
                </div>
            </div>
        </div>
        <div class="card-body">
            <div class="mask-example">
                <label class="mask-label">
                    Integer
                    <span class="mask-pattern">#########</span>
                </label>
                <div class="input-group">
                    <span class="input-group-text">#</span>
                    <input type="text" class="form-control" placeholder="123456789" data-mask="integer">
                </div>
            </div>

            <div class="mask-example">
                <label class="mask-label">
                    Decimal (2 places)
                    <span class="mask-pattern">###.##</span>
                </label>
                <div class="input-group">
                    <span class="input-group-text">#</span>
                    <input type="text" class="form-control" placeholder="123.45" data-mask="decimal">
                </div>
            </div>

            <div class="mask-example">
                <label class="mask-label">
                    Thousands Separator
                    <span class="mask-pattern">###,###,###</span>
                </label>
                <div class="input-group">
                    <span class="input-group-text">#</span>
                    <input type="text" class="form-control" placeholder="1,234,567" data-mask="number-thousands">
                </div>
            </div>
        </div>
    </div>

    <div class="content-card">
        <div class="card-header">
            <div class="card-header-left">
                <div class="card-icon bg-success">
                    <i class="fa-solid fa-dollar-sign"></i>
                </div>
                <div>
                    <h3>Currency</h3>
                    <p class="card-subtitle">Money format masking</p>
                </div>
            </div>
        </div>
        <div class="card-body">
            <div class="mask-example">
                <label class="mask-label">
                    USD
                    <span class="mask-pattern">$#,###.##</span>
                </label>
                <div class="input-group">
                    <span class="input-group-text">$
                    </span>
                    <input type="text" class="form-control" placeholder="$1,234.56" data-mask="currency-usd">
                </div>
            </div>

            <div class="mask-example">
                <label class="mask-label">
                    IDR (Rupiah)
                    <span class="mask-pattern">Rp ###.###.###</span>
                </label>
                <div class="input-group">
                    <span class="input-group-text">Rp</span>
                    <input type="text" class="form-control" placeholder="Rp 1.234.567" data-mask="currency-idr">
                </div>
            </div>

            <div class="mask-example">
                <label class="mask-label">
                    EUR (Euro)
                    <span class="mask-pattern">€#,###.##</span>
                </label>
                <div class="input-group">
                    <span class="input-group-text">€</span>
                    <input type="text" class="form-control" placeholder="€1.234,56" data-mask="currency-eur">
                </div>
            </div>
        </div>
    </div>

    <div class="content-card">
        <div class="card-header">
            <div class="card-header-left">
                <div class="card-icon bg-info">
                    <i class="fa-solid fa-percent"></i>
                </div>
                <div>
                    <h3>Percentage</h3>
                    <p class="card-subtitle">Percent format masking</p>
                </div>
            </div>
        </div>
        <div class="card-body">
            <div class="mask-example">
                <label class="mask-label">
                    Percentage
                    <span class="mask-pattern">###%</span>
                </label>
                <div class="input-group">
                    <input type="text" class="form-control" placeholder="85" data-mask="percentage">
                    <span class="input-group-text">%</span>
                </div>
            </div>

            <div class="mask-example">
                <label class="mask-label">
                    Decimal Percent
                    <span class="mask-pattern">##.##%</span>
                </label>
                <div class="input-group">
                    <input type="text" class="form-control" placeholder="12.50" data-mask="percentage-decimal">
                    <span class="input-group-text">%</span>
                </div>
            </div>

            <div class="mask-example">
                <label class="mask-label">
                    Progress (0-100)
                    <span class="mask-pattern">###</span>
                </label>
                <div class="input-group">
                    <input type="text" class="form-control" placeholder="75" data-mask="percentage">
                    <span class="input-group-text">/100</span>
                </div>
            </div>
        </div>
    </div>
</div>

<!-- ID & Document Masks -->
<div class="im-section-title">
    <i class="fa-solid fa-id-card"></i>
    ID & Document Masks
</div>

<div class="input-masks-grid">
    <div class="content-card">
        <div class="card-header">
            <div class="card-header-left">
                <div class="card-icon bg-warning">
                    <i class="fa-solid fa-credit-card"></i>
                </div>
                <div>
                    <h3>Financial Documents</h3>
                    <p class="card-subtitle">Card and account numbers</p>
                </div>
            </div>
        </div>
        <div class="card-body">
            <div class="mask-example">
                <label class="mask-label">
                    Credit Card
                    <span class="mask-pattern">#### #### #### ####</span>
                </label>
                <div class="input-group">
                    <span class="input-group-text">
                        <i class="fa-solid fa-credit-card"></i>
                    </span>
                    <input type="text" class="form-control" placeholder="1234 5678 9012 3456" maxlength="19" data-mask="credit-card">
                </div>
                <div class="mask-placeholder">
                    <i class="fa-solid fa-circle-info"></i>
                    Format: <code>4111 1111 1111 1111</code>
                </div>
            </div>

            <div class="mask-example">
                <label class="mask-label">
                    CVV
                    <span class="mask-pattern">###</span>
                </label>
                <div class="input-group">
                    <span class="input-group-text">
                        <i class="fa-solid fa-lock"></i>
                    </span>
                    <input type="text" class="form-control" placeholder="123" maxlength="4" data-mask="cvv">
                </div>
                <div class="mask-placeholder">
                    <i class="fa-solid fa-circle-info"></i>
                    Format: <code>123</code> or <code>1234</code>
                </div>
            </div>

            <div class="mask-example">
                <label class="mask-label">
                    Bank Account
                    <span class="mask-pattern">############</span>
                </label>
                <div class="input-group">
                    <span class="input-group-text">
                        <i class="fa-solid fa-building-columns"></i>
                    </span>
                    <input type="text" class="form-control" placeholder="1234567890" maxlength="12" data-mask="bank-account">
                </div>
                <div class="mask-placeholder">
                    <i class="fa-solid fa-circle-info"></i>
                    Format: <code>1234567890</code>
                </div>
            </div>

            <div class="divider"></div>

            <div class="live-preview">
                <div class="live-preview-label">Card Preview:</div>
                <div class="live-preview-value">4111 **** **** 1111</div>
            </div>
        </div>
    </div>

    <div class="content-card">
        <div class="card-header">
            <div class="card-header-left">
                <div class="card-icon bg-danger">
                    <i class="fa-solid fa-passport"></i>
                </div>
                <div>
                    <h3>Government IDs</h3>
                    <p class="card-subtitle">National ID and passport</p>
                </div>
            </div>
        </div>
        <div class="card-body">
            <div class="mask-example">
                <label class="mask-label">
                    Indonesian NIK (KTP)
                    <span class="mask-pattern">################</span>
                </label>
                <div class="input-group">
                    <span class="input-group-text">
                        <i class="fa-solid fa-id-card"></i>
                    </span>
                    <input type="text" class="form-control" placeholder="3201234567890001" maxlength="16" data-mask="nik">
                </div>
                <div class="mask-placeholder">
                    <i class="fa-solid fa-circle-info"></i>
                    Format: <code>16 digits</code>
                </div>
            </div>

            <div class="mask-example">
                <label class="mask-label">
                    Passport Number
                    <span class="mask-pattern">AA########</span>
                </label>
                <div class="input-group">
                    <span class="input-group-text">
                        <i class="fa-solid fa-passport"></i>
                    </span>
                    <input type="text" class="form-control" placeholder="A1234567" maxlength="9" data-mask="passport">
                </div>
                <div class="mask-placeholder">
                    <i class="fa-solid fa-circle-info"></i>
                    Format: <code>A1234567</code>
                </div>
            </div>

            <div class="mask-example">
                <label class="mask-label">
                    Tax ID (NPWP)
                    <span class="mask-pattern">##.###.###.#.###.###</span>
                </label>
                <div class="input-group">
                    <span class="input-group-text">
                        <i class="fa-solid fa-file-invoice-dollar"></i>
                    </span>
                    <input type="text" class="form-control" placeholder="12.345.678.9-012.345" data-mask="npwp">
                </div>
                <div class="mask-placeholder">
                    <i class="fa-solid fa-circle-info"></i>
                    Format: <code>12.345.678.9-012.345</code>
                </div>
            </div>

            <div class="divider"></div>

            <div class="form-group">
                <h4 style="font-size: 14px; margin-bottom: 12px;">Security Features:</h4>
                <ul class="feature-list">
                    <li>
                        <i class="fa-solid fa-circle-check"></i>
                        <span>Digit count validation</span>
                    </li>
                    <li>
                        <i class="fa-solid fa-circle-check"></i>
                        <span>Luhn algorithm (credit cards)</span>
                    </li>
                    <li>
                        <i class="fa-solid fa-circle-check"></i>
                        <span>Auto-spacing for readability</span>
                    </li>
                    <li>
                        <i class="fa-solid fa-circle-check"></i>
                        <span>Masked display for sensitive data</span>
                    </li>
                </ul>
            </div>
        </div>
    </div>
</div>

<!-- Custom Pattern Masks -->
<div class="im-section-title">
    <i class="fa-solid fa-wand-magic-sparkles"></i>
    Custom Pattern Masks
</div>

<div class="input-masks-grid">
    <div class="content-card">
        <div class="card-header">
            <div class="card-header-left">
                <div class="card-icon bg-primary">
                    <i class="fa-solid fa-barcode"></i>
                </div>
                <div>
                    <h3>Product Codes</h3>
                    <p class="card-subtitle">SKU and barcode masking</p>
                </div>
            </div>
        </div>
        <div class="card-body">
            <div class="mask-example">
                <label class="mask-label">
                    ISBN (Book)
                    <span class="mask-pattern">###-#-#####-###-#</span>
                </label>
                <div class="input-group">
                    <span class="input-group-text">
                        <i class="fa-solid fa-book"></i>
                    </span>
                    <input type="text" class="form-control" placeholder="978-3-16148-410-0" data-mask="isbn">
                </div>
            </div>

            <div class="mask-example">
                <label class="mask-label">
                    Serial Number
                    <span class="mask-pattern">AAA-###-#####</span>
                </label>
                <div class="input-group">
                    <span class="input-group-text">
                        <i class="fa-solid fa-fingerprint"></i>
                    </span>
                    <input type="text" class="form-control" placeholder="ABC-123-45678" data-mask="serial">
                </div>
            </div>

            <div class="mask-example">
                <label class="mask-label">
                    License Plate
                    <span class="mask-pattern">AA ### AAA</span>
                </label>
                <div class="input-group">
                    <span class="input-group-text">
                        <i class="fa-solid fa-car"></i>
                    </span>
                    <input type="text" class="form-control" placeholder="B 1234 XYZ" data-mask="license-plate">
                </div>
            </div>
        </div>
    </div>

    <div class="content-card">
        <div class="card-header">
            <div class="card-header-left">
                <div class="card-icon bg-success">
                    <i class="fa-solid fa-network-wired"></i>
                </div>
                <div>
                    <h3>Network & Tech</h3>
                    <p class="card-subtitle">IP and MAC address masking</p>
                </div>
            </div>
        </div>
        <div class="card-body">
            <div class="mask-example">
                <label class="mask-label">
                    IPv4 Address
                    <span class="mask-pattern">###.###.###.###</span>
                </label>
                <div class="input-group">
                    <span class="input-group-text">
                        <i class="fa-solid fa-wifi"></i>
                    </span>
                    <input type="text" class="form-control" placeholder="192.168.1.100" data-mask="ip-address">
                </div>
                <div class="mask-placeholder">
                    <i class="fa-solid fa-circle-info"></i>
                    Format: <code>192.168.1.100</code>
                </div>
            </div>

            <div class="mask-example">
                <label class="mask-label">
                    MAC Address
                    <span class="mask-pattern">AA:AA:AA:AA:AA:AA</span>
                </label>
                <div class="input-group">
                    <span class="input-group-text">
                        <i class="fa-solid fa-ethernet"></i>
                    </span>
                    <input type="text" class="form-control" placeholder="00:1B:44:11:3A:B7" data-mask="mac-address">
                </div>
                <div class="mask-placeholder">
                    <i class="fa-solid fa-circle-info"></i>
                    Format: <code>00:1B:44:11:3A:B7</code>
                </div>
            </div>

            <div class="mask-example">
                <label class="mask-label">
                    Color Code (HEX)
                    <span class="mask-pattern">#AAAAAA</span>
                </label>
                <div class="input-group">
                    <span class="input-group-text">
                        <i class="fa-solid fa-palette"></i>
                    </span>
                    <input type="text" class="form-control" placeholder="#0078D4" maxlength="7" data-mask="hex-color">
                </div>
                <div class="mask-placeholder">
                    <i class="fa-solid fa-circle-info"></i>
                    Format: <code>#0078D4</code>
                </div>
            </div>
        </div>
    </div>
</div>

<!-- Complete Form Example -->
<div class="im-section-title">
    <i class="fa-solid fa-clipboard-list"></i>
    Complete Payment Form
</div>

<div class="input-masks-grid full-width">
    <div class="content-card">
        <div class="card-header">
            <div class="card-header-left">
                <div class="card-icon bg-primary">
                    <i class="fa-solid fa-credit-card"></i>
                </div>
                <div>
                    <h3>Payment Form with Input Masks</h3>
                    <p class="card-subtitle">Real-world example with various masks</p>
                </div>
            </div>
        </div>
        <div class="card-body">
            <form style="max-width: 800px;">
                <div class="form-group">
                    <label style="font-weight: 600; margin-bottom: 8px; display: block;">Cardholder Name</label>
                    <input type="text" class="form-control" placeholder="JOHN DOE" style="text-transform: uppercase;">
                </div>

                <div class="form-group">
                    <label style="font-weight: 600; margin-bottom: 8px; display: block;">Card Number <span style="color: var(--danger);">*</span></label>
                    <div class="input-group">
                        <span class="input-group-text">
                            <i class="fa-solid fa-credit-card"></i>
                        </span>
                        <input type="text" class="form-control" placeholder="1234 5678 9012 3456" maxlength="19" data-mask="credit-card">
                    </div>
                    <div class="mask-placeholder">
                        <i class="fa-solid fa-circle-info"></i>
                        Auto-format: <code>#### #### #### ####</code>
                    </div>
                </div>

                <div style="display: grid; grid-template-columns: 1fr 1fr; gap: 16px;">
                    <div class="form-group">
                        <label style="font-weight: 600; margin-bottom: 8px; display: block;">Expiry Date <span style="color: var(--danger);">*</span></label>
                        <div class="input-group">
                            <input type="text" class="form-control" placeholder="MM/YY" maxlength="5" data-mask="expiry-date">
                            <span class="input-group-text">
                                <i class="fa-solid fa-calendar"></i>
                            </span>
                        </div>
                        <div class="mask-placeholder">
                            <i class="fa-solid fa-circle-info"></i>
                            Format: <code>MM/YY</code>
                        </div>
                    </div>

                    <div class="form-group">
                        <label style="font-weight: 600; margin-bottom: 8px; display: block;">CVV <span style="color: var(--danger);">*</span></label>
                        <div class="input-group">
                            <input type="text" class="form-control" placeholder="123" maxlength="4" data-mask="cvv">
                            <span class="input-group-text">
                                <i class="fa-solid fa-lock"></i>
                            </span>
                        </div>
                        <div class="mask-placeholder">
                            <i class="fa-solid fa-circle-info"></i>
                            Format: <code>###</code> or <code>####</code>
                        </div>
                    </div>
                </div>

                <div class="divider"></div>

                <div class="form-group">
                    <label style="font-weight: 600; margin-bottom: 8px; display: block;">Billing Phone</label>
                    <div class="input-group">
                        <select class="form-control" style="max-width: 120px;">
                            <option>🇮🇩 +62</option>
                            <option>🇺🇸 +1</option>
                        </select>
                        <input type="tel" class="form-control" placeholder="812-3456-7890" data-mask="phone">
                    </div>
                </div>

                <div class="form-group">
                    <label style="font-weight: 600; margin-bottom: 8px; display: block;">Payment Amount</label>
                    <div class="input-group">
                        <span class="input-group-text">Rp</span>
                        <input type="text" class="form-control" placeholder="1.234.567" data-mask="currency-idr">
                    </div>
                </div>

                <div style="display: flex; gap: 12px; margin-top: 24px;">
                    <button type="submit" class="btn btn-primary" style="flex: 1;">
                        <i class="fa-solid fa-lock"></i>
                        Process Payment
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

<!-- Implementation Guide -->
<div class="input-masks-grid full-width">
    <div class="content-card">
        <div class="card-header">
            <div class="card-header-left">
                <div class="card-icon bg-info">
                    <i class="fa-solid fa-code"></i>
                </div>
                <div>
                    <h3>Implementation Guide</h3>
                    <p class="card-subtitle">How to implement input masks with IMask.js</p>
                </div>
            </div>
        </div>
        <div class="card-body">
            <div class="code-block">
                <div style="color: var(--text-tertiary); margin-bottom: 8px;">1. Include IMask.js:</div>
                <code style="color: var(--accent);">
                    &lt;script src="https://unpkg.com/imask"&gt;&lt;/script&gt;
                </code>
            </div>

            <div class="code-block">
                <div style="color: var(--text-tertiary); margin-bottom: 8px; margin-top: 16px;">2. Initialize Phone Mask:</div>
                <code style="color: var(--success);">
                    const phoneMask = IMask(<br>
                    &nbsp;&nbsp;document.getElementById('phone'), {<br>
                    &nbsp;&nbsp;&nbsp;&nbsp;mask: '+62 {000}-{0000}-{0000}'<br>
                    &nbsp;&nbsp;}<br>
                    );
                </code>
            </div>

            <div class="code-block">
                <div style="color: var(--text-tertiary); margin-bottom: 8px; margin-top: 16px;">3. Initialize Date Mask:</div>
                <code style="color: var(--warning);">
                    const dateMask = IMask(<br>
                    &nbsp;&nbsp;document.getElementById('date'), {<br>
                    &nbsp;&nbsp;&nbsp;&nbsp;mask: 'DD/MM/YYYY',<br>
                    &nbsp;&nbsp;&nbsp;&nbsp;blocks: {<br>
                    &nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;DD: { mask: IMask.MaskedRange, from: 1, to: 31 },<br>
                    &nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;MM: { mask: IMask.MaskedRange, from: 1, to: 12 },<br>
                    &nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;YYYY: { mask: IMask.MaskedRange, from: 1900, to: 2099 }<br>
                    &nbsp;&nbsp;&nbsp;&nbsp;}<br>
                    &nbsp;&nbsp;}<br>
                    );
                </code>
            </div>

            <div class="divider"></div>

            <div class="form-group">
                <h4 style="font-size: 14px; margin-bottom: 12px;">Popular Libraries:</h4>
                <ul class="feature-list">
                    <li>
                        <i class="fa-solid fa-circle-check"></i>
                        <span><strong>IMask.js</strong> - Lightweight, vanilla JS</span>
                    </li>
                    <li>
                        <i class="fa-solid fa-circle-check"></i>
                        <span><strong>jQuery InputMask</strong> - jQuery plugin</span>
                    </li>
                    <li>
                        <i class="fa-solid fa-circle-check"></i>
                        <span><strong>Cleave.js</strong> - Auto-formatting</span>
                    </li>
                    <li>
                        <i class="fa-solid fa-circle-check"></i>
                        <span><strong>React Input Mask</strong> - React component</span>
                    </li>
                </ul>
            </div>
        </div>
    </div>
</div>
@endsection
