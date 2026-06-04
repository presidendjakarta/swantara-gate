@extends('layouts.app')

@section('title', 'Select2')

@push('styles')
<!-- Select2 CSS -->
<link href="https://cdn.jsdelivr.net/npm/select2@4.1.0-rc.0/dist/css/select2.min.css" rel="stylesheet" />
@endpush

@section('content')
<div class="page-header">
    <div>
        <h1>Select2</h1>
        <p>Enhanced select dropdown dengan search, multi-select, tagging, dan custom styling</p>
    </div>
    <div style="display: flex; gap: 10px;">
        <button class="btn btn-secondary">
            <i class="fa-solid fa-book"></i>
            <span>Documentation</span>
        </button>
        <button class="btn btn-primary">
            <i class="fa-solid fa-list"></i>
            <span>Try Select</span>
        </button>
    </div>
</div>

<!-- Info Alert -->
<div class="content-card" style="margin-bottom: 24px; border-left: 4px solid var(--accent);">
    <div class="card-body" style="padding: 16px 20px;">
        <div style="display: flex; gap: 12px; align-items: start;">
            <i class="fa-solid fa-circle-info" style="color: var(--accent); font-size: 20px; margin-top: 2px;"></i>
            <div style="flex: 1;">
                <h4 style="margin-bottom: 4px; font-size: 14px;">What is Select2?</h4>
                <p style="font-size: 13px; color: var(--text-secondary); margin: 0;">Select2 is a jQuery-based replacement for select boxes. It supports searching, remote data sets, and infinite scrolling of results. Examples below show various configurations and use cases.</p>
            </div>
        </div>
    </div>
</div>

<!-- Basic Single Select -->
<div class="s2-section-title">
    <i class="fa-solid fa-list-check"></i>
    Basic Single Select <span class="badge badge-primary">Essential</span>
</div>

<div class="select2-grid">
    <div class="content-card">
        <div class="card-header">
            <div class="card-header-left">
                <div class="card-icon bg-primary">
                    <i class="fa-solid fa-globe"></i>
                </div>
                <div>
                    <h3>Country Selection</h3>
                    <p class="card-subtitle">Single country picker</p>
                </div>
            </div>
        </div>
        <div class="card-body">
            <div class="s2-example">
                <label class="s2-label">
                    Select Country
                    <span class="s2-hint">With flags</span>
                </label>
                <div class="select-wrapper">
                    <select>
                        <option value="">-- Select Country --</option>
                        <option value="id" selected>🇮🇩 Indonesia</option>
                        <option value="us">🇺🇸 United States</option>
                        <option value="uk">🇬🇧 United Kingdom</option>
                        <option value="jp">🇯🇵 Japan</option>
                        <option value="kr">🇰🇷 South Korea</option>
                        <option value="sg">🇸🇬 Singapore</option>
                        <option value="my">🇲🇾 Malaysia</option>
                        <option value="au">🇦🇺 Australia</option>
                    </select>
                </div>
                <div class="s2-helper">
                    <i class="fa-solid fa-circle-info"></i>
                    Flag emojis included
                </div>
            </div>

            <div class="s2-example">
                <label class="s2-label">
                    Select Language
                    <span class="s2-hint">Single choice</span>
                </label>
                <div class="select-wrapper">
                    <select>
                        <option value="">-- Select Language --</option>
                        <option value="en">🇬🇧 English</option>
                        <option value="id" selected>🇮🇩 Bahasa Indonesia</option>
                        <option value="jp">🇯🇵 日本語</option>
                        <option value="kr">🇰🇷 한국어</option>
                        <option value="cn">🇨🇳 中文</option>
                        <option value="es">🇪🇸 Español</option>
                        <option value="fr">🇫🇷 Français</option>
                    </select>
                </div>
                <div class="s2-helper">
                    <i class="fa-solid fa-circle-info"></i>
                    Pre-selected value
                </div>
            </div>

            <div class="s2-example">
                <label class="s2-label">
                    Select Timezone
                    <span class="s2-hint">UTC offset</span>
                </label>
                <div class="select-wrapper">
                    <select>
                        <option value="">-- Select Timezone --</option>
                        <option value="wib" selected>WIB (UTC+7)</option>
                        <option value="wita">WITA (UTC+8)</option>
                        <option value="wit">WIT (UTC+9)</option>
                        <option value="est">EST (UTC-5)</option>
                        <option value="pst">PST (UTC-8)</option>
                        <option value="gmt">GMT (UTC+0)</option>
                        <option value="jst">JST (UTC+9)</option>
                    </select>
                </div>
            </div>

            <div class="divider"></div>

            <div class="select-preview">
                <div class="select-preview-label">Selected Country:</div>
                <div class="select-preview-value">🇮🇩 Indonesia (ID)</div>
            </div>
        </div>
    </div>

    <div class="content-card">
        <div class="card-header">
            <div class="card-header-left">
                <div class="card-icon bg-success">
                    <i class="fa-solid fa-layer-group"></i>
                </div>
                <div>
                    <h3>Grouped Options</h3>
                    <p class="card-subtitle">Optgroup categories</p>
                </div>
            </div>
        </div>
        <div class="card-body">
            <div class="s2-example">
                <label class="s2-label">
                    Select Category
                    <span class="s2-hint">Grouped</span>
                </label>
                <div class="select-wrapper">
                    <select>
                        <option value="">-- Select Category --</option>
                        <optgroup label="Electronics">
                            <option value="phone">📱 Smartphone</option>
                            <option value="laptop">💻 Laptop</option>
                            <option value="tablet">📟 Tablet</option>
                        </optgroup>
                        <optgroup label="Clothing">
                            <option value="shirt">👕 Shirt</option>
                            <option value="pants">👖 Pants</option>
                            <option value="shoes">👟 Shoes</option>
                        </optgroup>
                        <optgroup label="Food">
                            <option value="fruit">🍎 Fruit</option>
                            <option value="vegetable">🥬 Vegetables</option>
                            <option value="meat">🥩 Meat</option>
                        </optgroup>
                    </select>
                </div>
                <div class="s2-helper">
                    <i class="fa-solid fa-circle-info"></i>
                    3 groups with icons
                </div>
            </div>

            <div class="s2-example">
                <label class="s2-label">
                    Select Department
                    <span class="s2-hint">Organizational</span>
                </label>
                <div class="select-wrapper">
                    <select>
                        <option value="">-- Select Department --</option>
                        <optgroup label="Engineering">
                            <option value="frontend">Frontend</option>
                            <option value="backend">Backend</option>
                            <option value="devops">DevOps</option>
                        </optgroup>
                        <optgroup label="Business">
                            <option value="sales">Sales</option>
                            <option value="marketing">Marketing</option>
                            <option value="finance">Finance</option>
                        </optgroup>
                        <optgroup label="Support">
                            <option value="hr">Human Resources</option>
                            <option value="it">IT Support</option>
                            <option value="customer">Customer Service</option>
                        </optgroup>
                    </select>
                </div>
            </div>

            <div class="s2-example">
                <label class="s2-label">
                    Select City
                    <span class="s2-hint">By region</span>
                </label>
                <div class="select-wrapper">
                    <select>
                        <option value="">-- Select City --</option>
                        <optgroup label="Java">
                            <option value="jkt">Jakarta</option>
                            <option value="sby">Surabaya</option>
                            <option value="bdg">Bandung</option>
                        </optgroup>
                        <optgroup label="Sumatra">
                            <option value="mdn">Medan</option>
                            <option value="pmb">Palembang</option>
                            <option value="pku">Pekanbaru</option>
                        </optgroup>
                        <optgroup label="Sulawesi">
                            <option value="mns">Manado</option>
                            <option value="mks">Makassar</option>
                        </optgroup>
                    </select>
                </div>
            </div>

            <div class="divider"></div>

            <div class="form-group">
                <h4 style="font-size: 14px; margin-bottom: 12px;">Features:</h4>
                <ul class="feature-list">
                    <li>
                        <i class="fa-solid fa-circle-check"></i>
                        <span>Optgroup support</span>
                    </li>
                    <li>
                        <i class="fa-solid fa-circle-check"></i>
                        <span>Category grouping</span>
                    </li>
                    <li>
                        <i class="fa-solid fa-circle-check"></i>
                        <span>Emoji icons</span>
                    </li>
                    <li>
                        <i class="fa-solid fa-circle-check"></i>
                        <span>Hierarchical options</span>
                    </li>
                </ul>
            </div>
        </div>
    </div>
</div>

<!-- Multi-Select -->
<div class="s2-section-title">
    <i class="fa-solid fa-check-double"></i>
    Multi-Select <span class="badge badge-success">Popular</span>
</div>

<div class="select2-grid">
    <div class="content-card">
        <div class="card-header">
            <div class="card-header-left">
                <div class="card-icon bg-info">
                    <i class="fa-solid fa-tags"></i>
                </div>
                <div>
                    <h3>Tags Selection</h3>
                    <p class="card-subtitle">Multiple tags with remove</p>
                </div>
            </div>
        </div>
        <div class="card-body">
            <div class="s2-example">
                <label class="s2-label">
                    Select Skills
                    <span class="s2-hint">Multi-select</span>
                </label>
                <div class="multi-select-tags">
                    <span class="multi-select-tag">
                        JavaScript
                        <span class="remove-tag">×</span>
                    </span>
                    <span class="multi-select-tag">
                        Laravel
                        <span class="remove-tag">×</span>
                    </span>
                    <span class="multi-select-tag">
                        React
                        <span class="remove-tag">×</span>
                    </span>
                    <input type="text" class="multi-select-input" placeholder="Add more...">
                </div>
                <div class="s2-helper">
                    <i class="fa-solid fa-circle-info"></i>
                    Click × to remove tag
                </div>
            </div>

            <div class="s2-example">
                <label class="s2-label">
                    Select Interests
                    <span class="s2-hint">Hobbies</span>
                </label>
                <div class="multi-select-tags">
                    <span class="multi-select-tag">
                        📚 Reading
                        <span class="remove-tag">×</span>
                    </span>
                    <span class="multi-select-tag">
                        🎵 Music
                        <span class="remove-tag">×</span>
                    </span>
                    <input type="text" class="multi-select-input" placeholder="Add more...">
                </div>
                <div class="s2-helper">
                    <i class="fa-solid fa-circle-info"></i>
                    Emoji tags supported
                </div>
            </div>

            <div class="s2-example">
                <label class="s2-label">
                    Select Technologies
                    <span class="s2-hint">Development stack</span>
                </label>
                <div class="multi-select-tags">
                    <span class="multi-select-tag">
                        HTML5
                        <span class="remove-tag">×</span>
                    </span>
                    <span class="multi-select-tag">
                        CSS3
                        <span class="remove-tag">×</span>
                    </span>
                    <span class="multi-select-tag">
                        JavaScript
                        <span class="remove-tag">×</span>
                    </span>
                    <span class="multi-select-tag">
                        PHP
                        <span class="remove-tag">×</span>
                    </span>
                    <span class="multi-select-tag">
                        MySQL
                        <span class="remove-tag">×</span>
                    </span>
                    <input type="text" class="multi-select-input" placeholder="Add more...">
                </div>
            </div>

            <div class="divider"></div>

            <div class="select-preview">
                <div class="select-preview-label">Selected Technologies:</div>
                <div class="select-preview-value">HTML5, CSS3, JavaScript, PHP, MySQL</div>
            </div>
        </div>
    </div>

    <div class="content-card">
        <div class="card-header">
            <div class="card-header-left">
                <div class="card-icon bg-warning">
                    <i class="fa-solid fa-list-check"></i>
                </div>
                <div>
                    <h3>Native Multi-Select</h3>
                    <p class="card-subtitle">Ctrl+Click selection</p>
                </div>
            </div>
        </div>
        <div class="card-body">
            <div class="s2-example">
                <label class="s2-label">
                    Select Permissions
                    <span class="s2-hint">Hold Ctrl</span>
                </label>
                <div class="select-wrapper">
                    <select multiple style="height: 120px;">
                        <option value="create" selected>✅ Create</option>
                        <option value="read" selected>👁️ Read</option>
                        <option value="update" selected>✏️ Update</option>
                        <option value="delete">🗑️ Delete</option>
                        <option value="export">📤 Export</option>
                        <option value="import">📥 Import</option>
                    </select>
                </div>
                <div class="s2-helper">
                    <i class="fa-solid fa-circle-info"></i>
                    Hold Ctrl to select multiple
                </div>
            </div>

            <div class="s2-example">
                <label class="s2-label">
                    Select Features
                    <span class="s2-hint">Multiple</span>
                </label>
                <div class="select-wrapper">
                    <select multiple style="height: 120px;">
                        <option value="dark">🌙 Dark Mode</option>
                        <option value="notifications" selected>🔔 Notifications</option>
                        <option value="analytics" selected>📊 Analytics</option>
                        <option value="api">🔌 API Access</option>
                        <option value="backup">💾 Backup</option>
                    </select>
                </div>
            </div>

            <div class="divider"></div>

            <div class="form-group">
                <h4 style="font-size: 14px; margin-bottom: 12px;">Multi-Select Features:</h4>
                <ul class="feature-list">
                    <li>
                        <i class="fa-solid fa-circle-check"></i>
                        <span>Tag-based selection</span>
                    </li>
                    <li>
                        <i class="fa-solid fa-circle-check"></i>
                        <span>Remove individual tags</span>
                    </li>
                    <li>
                        <i class="fa-solid fa-circle-check"></i>
                        <span>Native Ctrl+Click</span>
                    </li>
                    <li>
                        <i class="fa-solid fa-circle-check"></i>
                        <span>Search and filter</span>
                    </li>
                </ul>
            </div>
        </div>
    </div>
</div>

<!-- With Icons & Search -->
<div class="s2-section-title">
    <i class="fa-solid fa-magnifying-glass"></i>
    Select with Icons & Search
</div>

<div class="select2-grid">
    <div class="content-card">
        <div class="card-header">
            <div class="card-header-left">
                <div class="card-icon bg-primary">
                    <i class="fa-solid fa-icons"></i>
                </div>
                <div>
                    <h3>Select with Left Icon</h3>
                    <p class="card-subtitle">Icon inside select</p>
                </div>
            </div>
        </div>
        <div class="card-body">
            <div class="s2-example">
                <label class="s2-label">
                    User Role
                    <span class="s2-hint">👤 icon</span>
                </label>
                <div class="select-wrapper select-with-icon" data-icon="user-shield">
                    <i class="fa-solid fa-user-shield"></i>
                    <select>
                        <option value="">-- Select Role --</option>
                        <option value="admin" selected>Administrator</option>
                        <option value="editor">Editor</option>
                        <option value="author">Author</option>
                        <option value="user">User</option>
                        <option value="guest">Guest</option>
                    </select>
                </div>
            </div>

            <div class="s2-example">
                <label class="s2-label">
                    Payment Method
                    <span class="s2-hint">💳 icon</span>
                </label>
                <div class="select-wrapper select-with-icon" data-icon="credit-card">
                    <i class="fa-solid fa-credit-card"></i>
                    <select>
                        <option value="">-- Select Payment --</option>
                        <option value="visa" selected>💳 Visa</option>
                        <option value="mastercard">💳 Mastercard</option>
                        <option value="paypal">🅿️ PayPal</option>
                        <option value="bank">🏦 Bank Transfer</option>
                        <option value="crypto">₿ Crypto</option>
                    </select>
                </div>
            </div>

            <div class="s2-example">
                <label class="s2-label">
                    Priority Level
                    <span class="s2-hint">⚡ icon</span>
                </label>
                <div class="select-wrapper select-with-icon" data-icon="flag">
                    <i class="fa-solid fa-flag"></i>
                    <select>
                        <option value="">-- Select Priority --</option>
                        <option value="critical">🔴 Critical</option>
                        <option value="high">🟠 High</option>
                        <option value="medium" selected>🟡 Medium</option>
                        <option value="low">🟢 Low</option>
                    </select>
                </div>
            </div>

            <div class="divider"></div>

            <div class="helper-text">
                <i class="fa-solid fa-circle-info"></i>
                Icons improve visual identification and UX
            </div>
        </div>
    </div>

    <div class="content-card">
        <div class="card-header">
            <div class="card-header-left">
                <div class="card-icon bg-success">
                    <i class="fa-solid fa-magnifying-glass"></i>
                </div>
                <div>
                    <h3>Search Functionality</h3>
                    <p class="card-subtitle">Type to filter options</p>
                </div>
            </div>
        </div>
        <div class="card-body">
            <div class="s2-example">
                <label class="s2-label">
                    Search Country
                    <span class="s2-hint">Type to search</span>
                </label>
                <div class="select-wrapper">
                    <select>
                        <option value="">🔍 Search country...</option>
                        <option value="id">🇮🇩 Indonesia</option>
                        <option value="us">🇺🇸 United States</option>
                        <option value="uk">🇬🇧 United Kingdom</option>
                        <option value="jp">🇯🇵 Japan</option>
                        <option value="kr">🇰🇷 South Korea</option>
                        <option value="sg">🇸🇬 Singapore</option>
                        <option value="my">🇲🇾 Malaysia</option>
                        <option value="au">🇦🇺 Australia</option>
                        <option value="ca">🇨🇦 Canada</option>
                        <option value="de">🇩🇪 Germany</option>
                        <option value="fr">🇫🇷 France</option>
                        <option value="it">🇮🇹 Italy</option>
                    </select>
                </div>
                <div class="s2-helper">
                    <i class="fa-solid fa-circle-info"></i>
                    12+ countries, searchable
                </div>
            </div>

            <div class="s2-example">
                <label class="s2-label">
                    Search Product
                    <span class="s2-hint">Large dataset</span>
                </label>
                <div class="select-wrapper">
                    <select>
                        <option value="">🔍 Search product...</option>
                        <option value="laptop">💻 Laptop Pro 15</option>
                        <option value="phone">📱 Smartphone X</option>
                        <option value="tablet">📟 Tablet Air</option>
                        <option value="watch">⌚ Smart Watch</option>
                        <option value="headphone">🎧 Wireless Headphones</option>
                        <option value="camera">📷 Digital Camera</option>
                        <option value="speaker">🔊 Bluetooth Speaker</option>
                    </select>
                </div>
                <div class="s2-helper">
                    <i class="fa-solid fa-circle-info"></i>
                    Type "phone" to filter
                </div>
            </div>

            <div class="divider"></div>

            <div class="form-group">
                <h4 style="font-size: 14px; margin-bottom: 12px;">Search Features:</h4>
                <ul class="feature-list">
                    <li>
                        <i class="fa-solid fa-circle-check"></i>
                        <span>Real-time filtering</span>
                    </li>
                    <li>
                        <i class="fa-solid fa-circle-check"></i>
                        <span>Case-insensitive search</span>
                    </li>
                    <li>
                        <i class="fa-solid fa-circle-check"></i>
                        <span>Keyboard navigation</span>
                    </li>
                    <li>
                        <i class="fa-solid fa-circle-check"></i>
                        <span>Large dataset support</span>
                    </li>
                </ul>
            </div>
        </div>
    </div>
</div>

<!-- States & Disabled -->
<div class="s2-section-title">
    <i class="fa-solid fa-circle-check"></i>
    States & Disabled Options
</div>

<div class="select2-grid">
    <div class="content-card">
        <div class="card-header">
            <div class="card-header-left">
                <div class="card-icon bg-success">
                    <i class="fa-solid fa-check"></i>
                </div>
                <div>
                    <h3>Valid State</h3>
                    <p class="card-subtitle">Green border, success</p>
                </div>
            </div>
        </div>
        <div class="card-body">
            <div class="s2-example">
                <label class="s2-label">Country (Valid)</label>
                <div class="select-wrapper">
                    <select class="is-valid">
                        <option value="">-- Select Country --</option>
                        <option value="id" selected>🇮🇩 Indonesia</option>
                        <option value="us">🇺🇸 United States</option>
                        <option value="uk">🇬🇧 United Kingdom</option>
                    </select>
                </div>
                <div class="s2-helper" style="color: var(--success);">
                    <i class="fa-solid fa-circle-check"></i>
                    <span>Valid selection</span>
                </div>
            </div>

            <div class="s2-example">
                <label class="s2-label">Language (Valid)</label>
                <div class="select-wrapper">
                    <select class="is-valid">
                        <option value="">-- Select Language --</option>
                        <option value="en">🇬🇧 English</option>
                        <option value="id" selected>🇮🇩 Bahasa Indonesia</option>
                        <option value="jp">🇯🇵 日本語</option>
                    </select>
                </div>
                <div class="s2-helper" style="color: var(--success);">
                    <i class="fa-solid fa-circle-check"></i>
                    <span>Language selected</span>
                </div>
            </div>

            <div class="divider"></div>

            <div class="helper-text">
                <i class="fa-solid fa-circle-info"></i>
                Green border indicates valid selection
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
                    <p class="card-subtitle">Red border, error message</p>
                </div>
            </div>
        </div>
        <div class="card-body">
            <div class="s2-example">
                <label class="s2-label">Country (Required)</label>
                <div class="select-wrapper">
                    <select class="is-invalid">
                        <option value="">-- Select Country --</option>
                        <option value="id">🇮🇩 Indonesia</option>
                        <option value="us">🇺🇸 United States</option>
                        <option value="uk">🇬🇧 United Kingdom</option>
                    </select>
                </div>
                <div class="s2-helper" style="color: var(--danger);">
                    <i class="fa-solid fa-circle-exclamation"></i>
                    <span>Country is required</span>
                </div>
            </div>

            <div class="s2-example">
                <label class="s2-label">Role (Required)</label>
                <div class="select-wrapper">
                    <select class="is-invalid">
                        <option value="">-- Select Role --</option>
                        <option value="admin">Administrator</option>
                        <option value="editor">Editor</option>
                        <option value="user">User</option>
                    </select>
                </div>
                <div class="s2-helper" style="color: var(--danger);">
                    <i class="fa-solid fa-circle-exclamation"></i>
                    <span>Please select a role</span>
                </div>
            </div>

            <div class="divider"></div>

            <div class="helper-text">
                <i class="fa-solid fa-circle-info"></i>
                Red border indicates missing required field
            </div>
        </div>
    </div>
</div>

<!-- Disabled Selects -->
<div class="s2-section-title">
    <i class="fa-solid fa-ban"></i>
    Disabled Selects & Options
</div>

<div class="select2-grid">
    <div class="content-card">
        <div class="card-header">
            <div class="card-header-left">
                <div class="card-icon bg-warning">
                    <i class="fa-solid fa-lock"></i>
                </div>
                <div>
                    <h3>Disabled Select</h3>
                    <p class="card-subtitle">Non-interactive dropdown</p>
                </div>
            </div>
        </div>
        <div class="card-body">
            <div class="s2-example">
                <label class="s2-label">Disabled Country</label>
                <div class="select-wrapper">
                    <select disabled>
                        <option value="id" selected>🇮🇩 Indonesia</option>
                        <option value="us">🇺🇸 United States</option>
                        <option value="uk">🇬🇧 United Kingdom</option>
                    </select>
                </div>
                <div class="s2-helper">
                    <i class="fa-solid fa-circle-info"></i>
                    Cannot be changed
                </div>
            </div>

            <div class="s2-example">
                <label class="s2-label">Disabled Role</label>
                <div class="select-wrapper">
                    <select disabled>
                        <option value="admin" selected>Administrator</option>
                        <option value="editor">Editor</option>
                        <option value="user">User</option>
                    </select>
                </div>
                <div class="s2-helper">
                    <i class="fa-solid fa-circle-info"></i>
                    Read-only value
                </div>
            </div>

            <div class="divider"></div>

            <div class="helper-text">
                <i class="fa-solid fa-circle-info"></i>
                Disabled selects have 50% opacity and no pointer events
            </div>
        </div>
    </div>

    <div class="content-card">
        <div class="card-header">
            <div class="card-header-left">
                <div class="card-icon bg-info">
                    <i class="fa-solid fa-ban"></i>
                </div>
                <div>
                    <h3>Disabled Options</h3>
                    <p class="card-subtitle">Individual disabled items</p>
                </div>
            </div>
        </div>
        <div class="card-body">
            <div class="s2-example">
                <label class="s2-label">Plan Selection</label>
                <div class="select-wrapper">
                    <select>
                        <option value="">-- Select Plan --</option>
                        <option value="free">Free Plan</option>
                        <option value="basic" selected>Basic Plan ($9/mo)</option>
                        <option value="pro">Pro Plan ($29/mo)</option>
                        <option value="enterprise" disabled>Enterprise (Coming Soon)</option>
                    </select>
                </div>
                <div class="s2-helper">
                    <i class="fa-solid fa-circle-info"></i>
                    Enterprise option is disabled
                </div>
            </div>

            <div class="s2-example">
                <label class="s2-label">Payment Method</label>
                <div class="select-wrapper">
                    <select>
                        <option value="">-- Select Payment --</option>
                        <option value="visa" selected>💳 Visa</option>
                        <option value="mastercard">💳 Mastercard</option>
                        <option value="paypal" disabled>🅿️ PayPal (Unavailable)</option>
                        <option value="bank">🏦 Bank Transfer</option>
                    </select>
                </div>
                <div class="s2-helper">
                    <i class="fa-solid fa-circle-info"></i>
                    PayPal temporarily unavailable
                </div>
            </div>

            <div class="divider"></div>

            <div class="helper-text">
                <i class="fa-solid fa-circle-info"></i>
                Disabled options appear grayed out in dropdown
            </div>
        </div>
    </div>
</div>

<!-- Registration Form Example -->
<div class="s2-section-title">
    <i class="fa-solid fa-clipboard-list"></i>
    Registration Form Example
</div>

<div class="select2-grid full-width">
    <div class="content-card">
        <div class="card-header">
            <div class="card-header-left">
                <div class="card-icon bg-primary">
                    <i class="fa-solid fa-user-plus"></i>
                </div>
                <div>
                    <h3>User Registration</h3>
                    <p class="card-subtitle">Complete form with various select types</p>
                </div>
            </div>
        </div>
        <div class="card-body">
            <form style="max-width: 800px;">
                <div style="display: grid; grid-template-columns: 1fr 1fr; gap: 20px;">
                    <div class="s2-example">
                        <label class="s2-label">Country <span style="color: var(--danger);">*</span></label>
                        <div class="select-wrapper">
                            <select required>
                                <option value="">-- Select Country --</option>
                                <option value="id">🇮🇩 Indonesia</option>
                                <option value="us">🇺🇸 United States</option>
                                <option value="uk">🇬🇧 United Kingdom</option>
                                <option value="jp">🇯🇵 Japan</option>
                                <option value="sg">🇸🇬 Singapore</option>
                            </select>
                        </div>
                    </div>

                    <div class="s2-example">
                        <label class="s2-label">Language <span style="color: var(--danger);">*</span></label>
                        <div class="select-wrapper">
                            <select required>
                                <option value="">-- Select Language --</option>
                                <option value="en">🇬🇧 English</option>
                                <option value="id">🇮🇩 Bahasa Indonesia</option>
                                <option value="jp">🇯🇵 日本語</option>
                                <option value="kr">🇰🇷 한국어</option>
                            </select>
                        </div>
                    </div>
                </div>

                <div class="s2-example">
                    <label class="s2-label">Role <span style="color: var(--danger);">*</span></label>
                    <div class="select-wrapper">
                        <select required>
                            <option value="">-- Select Role --</option>
                            <option value="user">👤 User</option>
                            <option value="creator">✍️ Content Creator</option>
                            <option value="admin">👨‍💼 Administrator</option>
                        </select>
                    </div>
                </div>

                <div class="s2-example">
                    <label class="s2-label">Interests (Multi-select)</label>
                    <div class="multi-select-tags">
                        <span class="multi-select-tag">
                            💻 Technology
                            <span class="remove-tag">×</span>
                        </span>
                        <span class="multi-select-tag">
                            📱 Mobile
                            <span class="remove-tag">×</span>
                        </span>
                        <input type="text" class="multi-select-input" placeholder="Add more...">
                    </div>
                </div>

                <div class="s2-example">
                    <label class="s2-label">Timezone</label>
                    <div class="select-wrapper">
                        <select>
                            <option value="">-- Select Timezone --</option>
                            <option value="wib" selected>WIB (UTC+7)</option>
                            <option value="wita">WITA (UTC+8)</option>
                            <option value="wit">WIT (UTC+9)</option>
                            <option value="est">EST (UTC-5)</option>
                            <option value="pst">PST (UTC-8)</option>
                        </select>
                    </div>
                </div>

                <div style="display: flex; gap: 12px; margin-top: 24px;">
                    <button type="submit" class="btn btn-primary" style="flex: 1;">
                        <i class="fa-solid fa-user-plus"></i>
                        Register
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
<div class="select2-grid full-width">
    <div class="content-card">
        <div class="card-header">
            <div class="card-header-left">
                <div class="card-icon bg-success">
                    <i class="fa-solid fa-code"></i>
                </div>
                <div>
                    <h3>Implementation Guide</h3>
                    <p class="card-subtitle">How to implement Select2 with jQuery</p>
                </div>
            </div>
        </div>
        <div class="card-body">
            <div class="code-block">
                <div style="color: var(--text-tertiary); margin-bottom: 8px;">1. Include Select2:</div>
                <code style="color: var(--accent);">
                    &lt;link href="https://cdn.jsdelivr.net/npm/select2@4.1.0-rc.0/dist/css/select2.min.css" rel="stylesheet" /&gt;<br>
                    &lt;script src="https://cdn.jsdelivr.net/npm/select2@4.1.0-rc.0/dist/js/select2.min.js"&gt;&lt;/script&gt;
                </code>
            </div>

            <div class="code-block">
                <div style="color: var(--text-tertiary); margin-bottom: 8px; margin-top: 16px;">2. Initialize Basic Select2:</div>
                <code style="color: var(--success);">
                    $(document).ready(function() {<br>
                    &nbsp;&nbsp;$('#countrySelect').select2({<br>
                    &nbsp;&nbsp;&nbsp;&nbsp;placeholder: 'Select a country',<br>
                    &nbsp;&nbsp;&nbsp;&nbsp;allowClear: true,<br>
                    &nbsp;&nbsp;&nbsp;&nbsp;width: '100%'<br>
                    &nbsp;&nbsp;});<br>
                    });
                </code>
            </div>

            <div class="code-block">
                <div style="color: var(--text-tertiary); margin-bottom: 8px; margin-top: 16px;">3. Initialize Multi-Select:</div>
                <code style="color: var(--warning);">
                    $('#skillsSelect').select2({<br>
                    &nbsp;&nbsp;placeholder: 'Select skills',<br>
                    &nbsp;&nbsp;closeOnSelect: false,<br>
                    &nbsp;&nbsp;tags: true,<br>
                    &nbsp;&nbsp;width: '100%'<br>
                    });
                </code>
            </div>

            <div class="code-block">
                <div style="color: var(--text-tertiary); margin-bottom: 8px; margin-top: 16px;">4. With AJAX Search:</div>
                <code style="color: var(--info);">
                    $('#searchSelect').select2({<br>
                    &nbsp;&nbsp;ajax: {<br>
                    &nbsp;&nbsp;&nbsp;&nbsp;url: '/api/search',<br>
                    &nbsp;&nbsp;&nbsp;&nbsp;delay: 250,<br>
                    &nbsp;&nbsp;&nbsp;&nbsp;data: function (params) {<br>
                    &nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;return { q: params.term };<br>
                    &nbsp;&nbsp;&nbsp;&nbsp;},<br>
                    &nbsp;&nbsp;&nbsp;&nbsp;minimumInputLength: 2<br>
                    &nbsp;&nbsp;}<br>
                    });
                </code>
            </div>

            <div class="divider"></div>

            <div class="form-group">
                <h4 style="font-size: 14px; margin-bottom: 12px;">Popular Libraries:</h4>
                <ul class="feature-list">
                    <li>
                        <i class="fa-solid fa-circle-check"></i>
                        <span><strong>Select2</strong> - jQuery-based, feature-rich</span>
                    </li>
                    <li>
                        <i class="fa-solid fa-circle-check"></i>
                        <span><strong>Choices.js</strong> - Vanilla JS, modern</span>
                    </li>
                    <li>
                        <i class="fa-solid fa-circle-check"></i>
                        <span><strong>Tom Select</strong> - Lightweight, taggable</span>
                    </li>
                    <li>
                        <i class="fa-solid fa-circle-check"></i>
                        <span><strong>React Select</strong> - React component</span>
                    </li>
                </ul>
            </div>
        </div>
    </div>
</div>
@endsection

@push('scripts')
<!-- Select2 JS -->
<script src="https://cdn.jsdelivr.net/npm/select2@4.1.0-rc.0/dist/js/select2.min.js"></script>
<script>
$(document).ready(function() {
    // Initialize all basic select elements with Select2
    $('.select-wrapper select').not('[multiple]').each(function() {
        const $select = $(this);
        const $wrapper = $select.closest('.select-wrapper');
        
        $select.select2({
            placeholder: function() {
                return $(this).data('placeholder') || 'Select an option';
            },
            allowClear: true,
            width: '100%',
            language: {
                noResults: function() {
                    return 'No results found';
                },
                searching: function() {
                    return 'Searching...';
                }
            }
        });
        
        // Mark wrapper as initialized to hide original icon
        $wrapper.addClass('select2-initialized');
    });
    
    // Initialize multi-select with Select2
    $('select[multiple]').select2({
        placeholder: 'Select options',
        closeOnSelect: false,
        width: '100%',
        allowClear: true
    });
    
    // Multi-select tag removal
    $(document).on('click', '.remove-tag', function() {
        $(this).parent('.multi-select-tag').fadeOut(200, function() {
            $(this).remove();
            updatePreview();
        });
    });
    
    // Add new tag on enter
    $('.multi-select-input').on('keypress', function(e) {
        if (e.which === 13 && $(this).val().trim() !== '') {
            const tag = $('<span class="multi-select-tag">' + 
                $(this).val().trim() + 
                ' <span class="remove-tag">×</span></span>');
            $(this).before(tag);
            $(this).val('');
            updatePreview();
        }
    });
    
    // Update preview box
    function updatePreview() {
        const previewValue = $('.select-preview-value');
        if (previewValue.length) {
            const tags = [];
            $('.multi-select-tag').each(function() {
                tags.push($(this).text().replace('×', '').trim());
            });
            if (tags.length > 0) {
                previewValue.text(tags.join(', '));
            }
        }
    }
    
    // Live update for single select preview
    $('.select-wrapper select').on('change', function() {
        const previewValue = $('.select-preview-value');
        if (previewValue.length && !$(this).attr('multiple')) {
            const selectedText = $(this).find('option:selected').text();
            const selectedValue = $(this).val();
            if (selectedValue) {
                previewValue.text(selectedText);
            }
        }
    });
});
</script>
@endpush
