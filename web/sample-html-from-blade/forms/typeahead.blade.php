@extends('layouts.app')

@section('title', 'Typeahead')

@push('styles')
<style>
/* ============================================
   START: Typeahead Component Styles
   ============================================ */

/* Typeahead Grid */
.typeahead-grid {
    display: grid;
    grid-template-columns: repeat(2, 1fr);
    gap: 24px;
    margin-bottom: 24px;
}

.typeahead-grid.full-width {
    grid-template-columns: 1fr;
}

/* Section Title */
.ta-section-title {
    display: flex;
    align-items: center;
    gap: 10px;
    font-size: 18px;
    font-weight: 700;
    color: var(--text-primary);
    margin: 32px 0 16px 0;
    padding-bottom: 12px;
    border-bottom: 2px solid var(--border-color);
}

.ta-section-title i {
    color: var(--accent);
    font-size: 20px;
}

/* Typeahead Wrapper */
.typeahead-wrapper {
    position: relative;
    margin-bottom: 8px;
}

.typeahead-with-icon {
    position: relative;
}

.typeahead-with-icon > i {
    position: absolute;
    left: 14px;
    top: 50%;
    transform: translateY(-50%);
    color: var(--text-tertiary);
    font-size: 14px;
    pointer-events: none;
    z-index: 1;
    transition: color 0.2s;
}

.typeahead-with-icon .typeahead-input {
    padding-left: 40px !important;
}

.typeahead-with-icon:focus-within > i {
    color: var(--accent);
}

/* Typeahead Input */
.typeahead-input {
    width: 100%;
    padding: 12px 16px;
    border: 2px solid var(--border-color);
    border-radius: 8px;
    font-size: 14px;
    transition: all 0.2s;
    background: var(--bg-primary);
    color: var(--text-primary);
}

.typeahead-input:focus {
    outline: none;
    border-color: var(--accent);
    box-shadow: 0 0 0 3px rgba(0, 120, 212, 0.1);
}

.typeahead-input::placeholder {
    color: var(--text-tertiary);
}

/* Typeahead Dropdown */
.typeahead-dropdown {
    position: absolute;
    top: calc(100% + 4px);
    left: 0;
    right: 0;
    background: var(--bg-primary);
    border: 2px solid var(--border-color);
    border-radius: 8px;
    box-shadow: 0 4px 16px rgba(0, 0, 0, 0.1);
    z-index: 1000;
    display: none;
    max-height: 400px;
    overflow-y: auto;
    animation: dropdownSlide 0.2s ease;
}

@keyframes dropdownSlide {
    from {
        opacity: 0;
        transform: translateY(-8px);
    }
    to {
        opacity: 1;
        transform: translateY(0);
    }
}

.typeahead-dropdown.show {
    display: block;
}

/* Typeahead Header */
.typeahead-header {
    padding: 10px 14px;
    font-size: 12px;
    font-weight: 600;
    color: var(--text-tertiary);
    text-transform: uppercase;
    letter-spacing: 0.5px;
    border-bottom: 1px solid var(--border-color);
    background: var(--bg-secondary);
}

/* Typeahead Item */
.typeahead-item {
    display: flex;
    align-items: center;
    gap: 12px;
    padding: 12px 14px;
    cursor: pointer;
    transition: background 0.15s;
    border-bottom: 1px solid var(--border-light);
}

.typeahead-item:last-child {
    border-bottom: none;
}

.typeahead-item:hover,
.typeahead-item.active {
    background: var(--bg-secondary);
}

.typeahead-item-icon {
    font-size: 20px;
    flex-shrink: 0;
    width: 28px;
    height: 28px;
    display: flex;
    align-items: center;
    justify-content: center;
    background: var(--bg-secondary);
    border-radius: 6px;
}

.typeahead-item-content {
    flex: 1;
    min-width: 0;
}

.typeahead-item-title {
    font-size: 14px;
    font-weight: 600;
    color: var(--text-primary);
    margin-bottom: 2px;
}

.typeahead-item-subtitle {
    font-size: 12px;
    color: var(--text-tertiary);
}

.typeahead-item-badge {
    padding: 4px 8px;
    background: var(--accent-bg);
    color: var(--accent);
    font-size: 11px;
    font-weight: 600;
    border-radius: 4px;
    flex-shrink: 0;
}

/* Highlight */
.highlight {
    background: rgba(0, 120, 212, 0.15);
    color: var(--accent);
    padding: 0 2px;
    border-radius: 2px;
    font-weight: 700;
}

/* Typeahead Footer */
.typeahead-footer {
    padding: 10px 14px;
    font-size: 11px;
    color: var(--text-tertiary);
    text-align: center;
    border-top: 1px solid var(--border-color);
    background: var(--bg-secondary);
}

/* Empty State */
.typeahead-empty {
    display: flex;
    flex-direction: column;
    align-items: center;
    justify-content: center;
    padding: 40px 20px;
    color: var(--text-tertiary);
}

.typeahead-empty i {
    font-size: 48px;
    margin-bottom: 12px;
    opacity: 0.3;
}

.typeahead-empty > div:first-of-type {
    font-size: 14px;
    font-weight: 600;
    color: var(--text-secondary);
}

.typeahead-empty > div:last-of-type {
    font-size: 12px;
    margin-top: 4px;
}

/* Loading State */
.typeahead-loading .typeahead-input {
    background-image: url("data:image/svg+xml,%3Csvg xmlns='http://www.w3.org/2000/svg' viewBox='0 0 24 24' fill='none' stroke='%23666' stroke-width='2'%3E%3Ccircle cx='12' cy='12' r='10' stroke-opacity='0.25'/%3E%3Cpath d='M12 2a10 10 0 0 1 10 10' stroke-linecap='round'/%3E%3C/svg%3E");
    background-repeat: no-repeat;
    background-position: right 12px center;
    background-size: 16px;
    animation: pulse 1.5s ease-in-out infinite;
}

@keyframes pulse {
    0%, 100% { opacity: 1; }
    50% { opacity: 0.6; }
}

/* Labels & Hints */
.ta-label {
    display: block;
    font-size: 13px;
    font-weight: 600;
    color: var(--text-primary);
    margin-bottom: 8px;
}

.ta-hint {
    font-size: 11px;
    font-weight: 400;
    color: var(--text-tertiary);
    margin-left: 4px;
}

.ta-helper {
    display: flex;
    align-items: center;
    gap: 6px;
    font-size: 12px;
    color: var(--text-secondary);
    padding: 8px 0;
}

.ta-helper i {
    font-size: 14px;
    color: var(--accent);
}

/* Recent Search */
.recent-search {
    display: inline-flex;
    align-items: center;
    gap: 6px;
    padding: 6px 10px;
    background: var(--bg-secondary);
    border: 1px solid var(--border-color);
    border-radius: 6px;
    font-size: 13px;
    color: var(--text-primary);
    cursor: pointer;
    transition: all 0.2s;
}

.recent-search:hover {
    background: var(--bg-tertiary);
    border-color: var(--accent);
}

.recent-search i {
    font-size: 12px;
    color: var(--text-tertiary);
}

.recent-search-remove {
    font-size: 18px;
    line-height: 1;
    color: var(--text-tertiary);
    cursor: pointer;
    transition: color 0.2s;
    margin-left: 2px;
}

.recent-search-remove:hover {
    color: var(--danger);
}

/* Feature List */
.feature-list {
    list-style: none;
    padding: 0;
    margin: 0;
}

.feature-list li {
    display: flex;
    align-items: center;
    gap: 10px;
    padding: 8px 0;
    font-size: 13px;
    color: var(--text-secondary);
}

.feature-list li i {
    color: var(--success);
    font-size: 14px;
}

/* Code Block */
.code-block {
    background: var(--bg-secondary);
    border: 1px solid var(--border-color);
    border-radius: 8px;
    padding: 16px;
    font-family: 'Consolas', 'Monaco', monospace;
    font-size: 12px;
    overflow-x: auto;
}

.code-block code {
    display: block;
    line-height: 1.6;
}

/* Divider */
.divider {
    height: 1px;
    background: var(--border-color);
    margin: 20px 0;
}

/* Helper Text */
.helper-text {
    display: flex;
    align-items: center;
    gap: 8px;
    font-size: 12px;
    color: var(--text-secondary);
    padding: 12px;
    background: var(--bg-secondary);
    border-radius: 6px;
}

.helper-text i {
    color: var(--info);
}

/* Responsive */
@media (max-width: 768px) {
    .typeahead-grid {
        grid-template-columns: 1fr;
    }
}

/* ============================================
   END: Typeahead Component Styles
   ============================================ */
</style>
@endpush

@section('content')
<div class="page-header">
    <div>
        <h1>Typeahead</h1>
        <p>Autocomplete search suggestions dengan dropdown results dan keyboard navigation</p>
    </div>
    <div style="display: flex; gap: 10px;">
        <button class="btn btn-secondary">
            <i class="fa-solid fa-book"></i>
            <span>Documentation</span>
        </button>
        <button class="btn btn-primary">
            <i class="fa-solid fa-magnifying-glass"></i>
            <span>Try Typeahead</span>
        </button>
    </div>
</div>

<!-- Info Alert -->
<div class="content-card" style="margin-bottom: 24px; border-left: 4px solid var(--accent);">
    <div class="card-body" style="padding: 16px 20px;">
        <div style="display: flex; gap: 12px; align-items: start;">
            <i class="fa-solid fa-circle-info" style="color: var(--accent); font-size: 20px; margin-top: 2px;"></i>
            <div style="flex: 1;">
                <h4 style="margin-bottom: 4px; font-size: 14px;">What is Typeahead?</h4>
                <p style="font-size: 13px; color: var(--text-secondary); margin: 0;">Typeahead provides autocomplete suggestions as users type. It improves UX by showing relevant results, reducing typing effort, and guiding input. Popular implementations include <strong>Typeahead.js</strong>, <strong>Autocomplete.js</strong>, and custom solutions.</p>
            </div>
        </div>
    </div>
</div>

<!-- Basic Typeahead -->
<div class="ta-section-title">
    <i class="fa-solid fa-magnifying-glass"></i>
    Basic Typeahead <span class="badge badge-primary">Essential</span>
</div>

<div class="typeahead-grid">
    <div class="content-card">
        <div class="card-header">
            <div class="card-header-left">
                <div class="card-icon bg-primary">
                    <i class="fa-solid fa-globe"></i>
                </div>
                <div>
                    <h3>Country Search</h3>
                    <p class="card-subtitle">Type to search countries</p>
                </div>
            </div>
        </div>
        <div class="card-body">
            <div class="ta-example">
                <label class="ta-label">
                    Search Country
                    <span class="ta-hint">Type "Ind" or "Uni"</span>
                </label>
                <div class="typeahead-wrapper typeahead-with-icon">
                    <i class="fa-solid fa-magnifying-glass"></i>
                    <input type="text" class="typeahead-input" placeholder="Start typing country name..." value="Ind">
                    <div class="typeahead-dropdown">
                        <div class="typeahead-header">Countries (4 results)</div>
                        <div class="typeahead-item active">
                            <div class="typeahead-item-icon">🇮🇩</div>
                            <div class="typeahead-item-content">
                                <div class="typeahead-item-title"><span class="highlight">Ind</span>onesia</div>
                                <div class="typeahead-item-subtitle">Jakarta • Southeast Asia</div>
                            </div>
                            <div class="typeahead-item-badge">Popular</div>
                        </div>
                        <div class="typeahead-item">
                            <div class="typeahead-item-icon">🇮🇳</div>
                            <div class="typeahead-item-content">
                                <div class="typeahead-item-title"><span class="highlight">Ind</span>ia</div>
                                <div class="typeahead-item-subtitle">New Delhi • South Asia</div>
                            </div>
                        </div>
                        <div class="typeahead-item">
                            <div class="typeahead-item-icon">🇮🇸</div>
                            <div class="typeahead-item-content">
                                <div class="typeahead-item-title"><span class="highlight">Ind</span>iana (USA)</div>
                                <div class="typeahead-item-subtitle">Indianapolis • North America</div>
                            </div>
                        </div>
                        <div class="typeahead-item">
                            <div class="typeahead-item-icon">🇬🇧</div>
                            <div class="typeahead-item-content">
                                <div class="typeahead-item-title"><span class="highlight">Uni</span>ted Kingdom</div>
                                <div class="typeahead-item-subtitle">London • Europe</div>
                            </div>
                        </div>
                        <div class="typeahead-footer">↑↓ Navigate • Enter Select • Esc Close</div>
                    </div>
                </div>
                <div class="ta-helper">
                    <i class="fa-solid fa-circle-info"></i>
                    4 results found for "Ind"
                </div>
            </div>

            <div class="ta-example">
                <label class="ta-label">
                    Search Language
                    <span class="ta-hint">No results</span>
                </label>
                <div class="typeahead-wrapper typeahead-with-icon">
                    <i class="fa-solid fa-magnifying-glass"></i>
                    <input type="text" class="typeahead-input" placeholder="Start typing language..." value="xyz">
                    <div class="typeahead-dropdown">
                        <div class="typeahead-empty">
                            <i class="fa-solid fa-circle-xmark"></i>
                            <div>No languages found for "xyz"</div>
                            <div style="font-size: 11px; margin-top: 4px;">Try different keywords</div>
                        </div>
                    </div>
                </div>
            </div>

            <div class="divider"></div>

            <div class="form-group">
                <h4 style="font-size: 14px; margin-bottom: 12px;">Features:</h4>
                <ul class="feature-list">
                    <li>
                        <i class="fa-solid fa-circle-check"></i>
                        <span>Real-time suggestions</span>
                    </li>
                    <li>
                        <i class="fa-solid fa-circle-check"></i>
                        <span>Highlight matching text</span>
                    </li>
                    <li>
                        <i class="fa-solid fa-circle-check"></i>
                        <span>Keyboard navigation</span>
                    </li>
                    <li>
                        <i class="fa-solid fa-circle-check"></i>
                        <span>Empty state handling</span>
                    </li>
                </ul>
            </div>
        </div>
    </div>

    <div class="content-card">
        <div class="card-header">
            <div class="card-header-left">
                <div class="card-icon bg-success">
                    <i class="fa-solid fa-user"></i>
                </div>
                <div>
                    <h3>User Search</h3>
                    <p class="card-subtitle">Find users by name/email</p>
                </div>
            </div>
        </div>
        <div class="card-body">
            <div class="ta-example">
                <label class="ta-label">
                    Search Users
                    <span class="ta-hint">Type "John"</span>
                </label>
                <div class="typeahead-wrapper typeahead-with-icon">
                    <i class="fa-solid fa-user"></i>
                    <input type="text" class="typeahead-input" placeholder="Search users..." value="John">
                    <div class="typeahead-dropdown">
                        <div class="typeahead-header">Users (3 results)</div>
                        <div class="typeahead-item active">
                            <div class="typeahead-item-icon">👤</div>
                            <div class="typeahead-item-content">
                                <div class="typeahead-item-title"><span class="highlight">John</span> Doe</div>
                                <div class="typeahead-item-subtitle">john@example.com • Admin</div>
                            </div>
                            <div class="typeahead-item-badge">Online</div>
                        </div>
                        <div class="typeahead-item">
                            <div class="typeahead-item-icon">👤</div>
                            <div class="typeahead-item-content">
                                <div class="typeahead-item-title"><span class="highlight">John</span> Smith</div>
                                <div class="typeahead-item-subtitle">john.smith@example.com • User</div>
                            </div>
                        </div>
                        <div class="typeahead-item">
                            <div class="typeahead-item-icon">👤</div>
                            <div class="typeahead-item-content">
                                <div class="typeahead-item-title"><span class="highlight">John</span>ny Cash</div>
                                <div class="typeahead-item-subtitle">jcash@example.com • Editor</div>
                            </div>
                        </div>
                        <div class="typeahead-footer">Showing 3 of 1,234 users</div>
                    </div>
                </div>
            </div>

            <div class="ta-example">
                <label class="ta-label">
                    Search Teams
                    <span class="ta-hint">Type "Dev"</span>
                </label>
                <div class="typeahead-wrapper typeahead-with-icon">
                    <i class="fa-solid fa-users"></i>
                    <input type="text" class="typeahead-input" placeholder="Search teams..." value="Dev">
                    <div class="typeahead-dropdown">
                        <div class="typeahead-header">Teams (2 results)</div>
                        <div class="typeahead-item">
                            <div class="typeahead-item-icon">👥</div>
                            <div class="typeahead-item-content">
                                <div class="typeahead-item-title"><span class="highlight">Dev</span>elopment Team</div>
                                <div class="typeahead-item-subtitle">12 members • Engineering</div>
                            </div>
                        </div>
                        <div class="typeahead-item">
                            <div class="typeahead-item-icon">👥</div>
                            <div class="typeahead-item-content">
                                <div class="typeahead-item-title"><span class="highlight">Dev</span>Ops Team</div>
                                <div class="typeahead-item-subtitle">8 members • Infrastructure</div>
                            </div>
                        </div>
                    </div>
                </div>
            </div>

            <div class="divider"></div>

            <div class="helper-text">
                <i class="fa-solid fa-circle-info"></i>
                Click on a result to select the user/team
            </div>
        </div>
    </div>
</div>

<!-- Advanced Typeahead -->
<div class="ta-section-title">
    <i class="fa-solid fa-wand-magic-sparkles"></i>
    Advanced Typeahead <span class="badge badge-success">Popular</span>
</div>

<div class="typeahead-grid">
    <div class="content-card">
        <div class="card-header">
            <div class="card-header-left">
                <div class="card-icon bg-info">
                    <i class="fa-solid fa-box"></i>
                </div>
                <div>
                    <h3>Product Search</h3>
                    <p class="card-subtitle">E-commerce product finder</p>
                </div>
            </div>
        </div>
        <div class="card-body">
            <div class="ta-example">
                <label class="ta-label">
                    Search Products
                    <span class="ta-hint">Type "Phone"</span>
                </label>
                <div class="typeahead-wrapper typeahead-with-icon">
                    <i class="fa-solid fa-box"></i>
                    <input type="text" class="typeahead-input" placeholder="Search products..." value="Phone">
                    <div class="typeahead-dropdown">
                        <div class="typeahead-header">Products (5 results)</div>
                        <div class="typeahead-item active">
                            <div class="typeahead-item-icon">📱</div>
                            <div class="typeahead-item-content">
                                <div class="typeahead-item-title">iPhone 15 Pro</div>
                                <div class="typeahead-item-subtitle">Apple • $999 • ⭐ 4.8</div>
                            </div>
                            <div class="typeahead-item-badge">New</div>
                        </div>
                        <div class="typeahead-item">
                            <div class="typeahead-item-icon">📱</div>
                            <div class="typeahead-item-content">
                                <div class="typeahead-item-title">Samsung Galaxy S24</div>
                                <div class="typeahead-item-subtitle">Samsung • $899 • ⭐ 4.7</div>
                            </div>
                        </div>
                        <div class="typeahead-item">
                            <div class="typeahead-item-icon">📱</div>
                            <div class="typeahead-item-content">
                                <div class="typeahead-item-title">Google Pixel 8</div>
                                <div class="typeahead-item-subtitle">Google • $699 • ⭐ 4.6</div>
                            </div>
                            <div class="typeahead-item-badge" style="background: var(--success-bg); color: var(--success);">Sale</div>
                        </div>
                        <div class="typeahead-item">
                            <div class="typeahead-item-icon">📱</div>
                            <div class="typeahead-item-content">
                                <div class="typeahead-item-title">OnePlus 12</div>
                                <div class="typeahead-item-subtitle">OnePlus • $799 • ⭐ 4.5</div>
                            </div>
                        </div>
                        <div class="typeahead-item">
                            <div class="typeahead-item-icon">📱</div>
                            <div class="typeahead-item-content">
                                <div class="typeahead-item-title">Xiaomi 14 Pro</div>
                                <div class="typeahead-item-subtitle">Xiaomi • $649 • ⭐ 4.4</div>
                            </div>
                        </div>
                        <div class="typeahead-footer">Showing 5 of 2,456 products</div>
                    </div>
                </div>
            </div>

            <div class="divider"></div>

            <div class="form-group">
                <h4 style="font-size: 14px; margin-bottom: 12px;">Product Features:</h4>
                <ul class="feature-list">
                    <li>
                        <i class="fa-solid fa-circle-check"></i>
                        <span>Product images/icons</span>
                    </li>
                    <li>
                        <i class="fa-solid fa-circle-check"></i>
                        <span>Price and rating display</span>
                    </li>
                    <li>
                        <i class="fa-solid fa-circle-check"></i>
                        <span>Category badges (New/Sale)</span>
                    </li>
                    <li>
                        <i class="fa-solid fa-circle-check"></i>
                        <span>Total results count</span>
                    </li>
                </ul>
            </div>
        </div>
    </div>

    <div class="content-card">
        <div class="card-header">
            <div class="card-header-left">
                <div class="card-icon bg-warning">
                    <i class="fa-solid fa-location-dot"></i>
                </div>
                <div>
                    <h3>Location Search</h3>
                    <p class="card-subtitle">Address autocomplete</p>
                </div>
            </div>
        </div>
        <div class="card-body">
            <div class="ta-example">
                <label class="ta-label">
                    Search Address
                    <span class="ta-hint">Type "Jakarta"</span>
                </label>
                <div class="typeahead-wrapper typeahead-with-icon">
                    <i class="fa-solid fa-location-dot"></i>
                    <input type="text" class="typeahead-input" placeholder="Enter address..." value="Jakarta">
                    <div class="typeahead-dropdown">
                        <div class="typeahead-header">Locations (4 results)</div>
                        <div class="typeahead-item active">
                            <div class="typeahead-item-icon">🏙️</div>
                            <div class="typeahead-item-content">
                                <div class="typeahead-item-title"><span class="highlight">Jakarta</span> Pusat</div>
                                <div class="typeahead-item-subtitle">DKI Jakarta • Indonesia</div>
                            </div>
                        </div>
                        <div class="typeahead-item">
                            <div class="typeahead-item-icon">🏙️</div>
                            <div class="typeahead-item-content">
                                <div class="typeahead-item-title"><span class="highlight">Jakarta</span> Selatan</div>
                                <div class="typeahead-item-subtitle">DKI Jakarta • Indonesia</div>
                            </div>
                        </div>
                        <div class="typeahead-item">
                            <div class="typeahead-item-icon">🏙️</div>
                            <div class="typeahead-item-content">
                                <div class="typeahead-item-title"><span class="highlight">Jakarta</span> Barat</div>
                                <div class="typeahead-item-subtitle">DKI Jakarta • Indonesia</div>
                            </div>
                        </div>
                        <div class="typeahead-item">
                            <div class="typeahead-item-icon">🏙️</div>
                            <div class="typeahead-item-content">
                                <div class="typeahead-item-title"><span class="highlight">Jakarta</span> Utara</div>
                                <div class="typeahead-item-subtitle">DKI Jakarta • Indonesia</div>
                            </div>
                        </div>
                        <div class="typeahead-footer">Powered by Google Places API</div>
                    </div>
                </div>
            </div>

            <div class="divider"></div>

            <div class="helper-text">
                <i class="fa-solid fa-circle-info"></i>
                Integrates with Google Places API for accurate addresses
            </div>
        </div>
    </div>
</div>

<!-- Loading & States -->
<div class="ta-section-title">
    <i class="fa-solid fa-spinner"></i>
    Loading & States
</div>

<div class="typeahead-grid">
    <div class="content-card">
        <div class="card-header">
            <div class="card-header-left">
                <div class="card-icon bg-info">
                    <i class="fa-solid fa-circle-notch"></i>
                </div>
                <div>
                    <h3>Loading State</h3>
                    <p class="card-subtitle">Fetching results from API</p>
                </div>
            </div>
        </div>
        <div class="card-body">
            <div class="ta-example">
                <label class="ta-label">Search with Loading</label>
                <div class="typeahead-wrapper typeahead-with-icon typeahead-loading">
                    <i class="fa-solid fa-magnifying-glass"></i>
                    <input type="text" class="typeahead-input" placeholder="Loading..." value="Search">
                </div>
                <div class="ta-helper">
                    <i class="fa-solid fa-circle-notch fa-spin"></i>
                    Fetching results from server...
                </div>
            </div>

            <div class="divider"></div>

            <div class="helper-text">
                <i class="fa-solid fa-circle-info"></i>
                Loading spinner appears while fetching data
            </div>
        </div>
    </div>

    <div class="content-card">
        <div class="card-header">
            <div class="card-header-left">
                <div class="card-icon bg-success">
                    <i class="fa-solid fa-clock-rotate-left"></i>
                </div>
                <div>
                    <h3>Recent Searches</h3>
                    <p class="card-subtitle">Quick access to history</p>
                </div>
            </div>
        </div>
        <div class="card-body">
            <div class="ta-example">
                <label class="ta-label">Recent Searches</label>
                <div style="display: flex; flex-wrap: wrap; gap: 8px; margin-bottom: 12px;">
                    <div class="recent-search">
                        <i class="fa-solid fa-clock"></i>
                        <span>Laravel</span>
                        <span class="recent-search-remove">×</span>
                    </div>
                    <div class="recent-search">
                        <i class="fa-solid fa-clock"></i>
                        <span>React</span>
                        <span class="recent-search-remove">×</span>
                    </div>
                    <div class="recent-search">
                        <i class="fa-solid fa-clock"></i>
                        <span>Vue.js</span>
                        <span class="recent-search-remove">×</span>
                    </div>
                    <div class="recent-search">
                        <i class="fa-solid fa-clock"></i>
                        <span>JavaScript</span>
                        <span class="recent-search-remove">×</span>
                    </div>
                </div>
                <div class="ta-helper">
                    <i class="fa-solid fa-circle-info"></i>
                    Click to search again, × to remove
                </div>
            </div>

            <div class="divider"></div>

            <div class="helper-text">
                <i class="fa-solid fa-circle-info"></i>
                Recent searches improve user experience
            </div>
        </div>
    </div>
</div>

<!-- Search Form Example -->
<div class="ta-section-title">
    <i class="fa-solid fa-clipboard-list"></i>
    Search Form Example
</div>

<div class="typeahead-grid full-width">
    <div class="content-card">
        <div class="card-header">
            <div class="card-header-left">
                <div class="card-icon bg-primary">
                    <i class="fa-solid fa-magnifying-glass"></i>
                </div>
                <div>
                    <h3>Advanced Search Form</h3>
                    <p class="card-subtitle">Multi-field search with typeahead</p>
                </div>
            </div>
        </div>
        <div class="card-body">
            <form style="max-width: 800px;">
                <div class="ta-example">
                    <label class="ta-label">Search Products <span style="color: var(--danger);">*</span></label>
                    <div class="typeahead-wrapper typeahead-with-icon">
                        <i class="fa-solid fa-magnifying-glass"></i>
                        <input type="text" class="typeahead-input" placeholder="What are you looking for?" required>
                    </div>
                </div>

                <div style="display: grid; grid-template-columns: 1fr 1fr; gap: 20px;">
                    <div class="ta-example">
                        <label class="ta-label">Category</label>
                        <div class="typeahead-wrapper">
                            <input type="text" class="typeahead-input" placeholder="Select category...">
                        </div>
                    </div>

                    <div class="ta-example">
                        <label class="ta-label">Brand</label>
                        <div class="typeahead-wrapper">
                            <input type="text" class="typeahead-input" placeholder="Select brand...">
                        </div>
                    </div>
                </div>

                <div style="display: grid; grid-template-columns: 1fr 1fr; gap: 20px;">
                    <div class="ta-example">
                        <label class="ta-label">Price Range</label>
                        <div style="display: grid; grid-template-columns: 1fr 1fr; gap: 12px;">
                            <input type="number" class="typeahead-input" placeholder="Min" min="0">
                            <input type="number" class="typeahead-input" placeholder="Max" min="0">
                        </div>
                    </div>

                    <div class="ta-example">
                        <label class="ta-label">Rating</label>
                        <div class="typeahead-wrapper">
                            <input type="text" class="typeahead-input" placeholder="Min rating...">
                        </div>
                    </div>
                </div>

                <div class="divider"></div>

                <div style="display: flex; gap: 12px;">
                    <button type="submit" class="btn btn-primary" style="flex: 1;">
                        <i class="fa-solid fa-magnifying-glass"></i>
                        Search
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
<div class="typeahead-grid full-width">
    <div class="content-card">
        <div class="card-header">
            <div class="card-header-left">
                <div class="card-icon bg-success">
                    <i class="fa-solid fa-code"></i>
                </div>
                <div>
                    <h3>Implementation Guide</h3>
                    <p class="card-subtitle">How to implement typeahead with JavaScript</p>
                </div>
            </div>
        </div>
        <div class="card-body">
            <div class="code-block">
                <div style="color: var(--text-tertiary); margin-bottom: 8px;">1. Include Typeahead.js:</div>
                <code style="color: var(--accent);">
                    &lt;script src="https://cdnjs.cloudflare.com/ajax/libs/typeahead.js/0.11.1/typeahead.bundle.min.js"&gt;&lt;/script&gt;
                </code>
            </div>

            <div class="code-block">
                <div style="color: var(--text-tertiary); margin-bottom: 8px; margin-top: 16px;">2. Initialize Typeahead:</div>
                <code style="color: var(--success);">
                    $('#searchInput').typeahead({<br>
                    &nbsp;&nbsp;hint: true,<br>
                    &nbsp;&nbsp;highlight: true,<br>
                    &nbsp;&nbsp;minLength: 1<br>
                    }, {<br>
                    &nbsp;&nbsp;name: 'countries',<br>
                    &nbsp;&nbsp;source: substringMatcher(countries)<br>
                    });
                </code>
            </div>

            <div class="code-block">
                <div style="color: var(--text-tertiary); margin-bottom: 8px; margin-top: 16px;">3. AJAX Data Source:</div>
                <code style="color: var(--warning);">
                    $('#searchInput').typeahead(null, {<br>
                    &nbsp;&nbsp;source: function(query, syncResults, asyncResults) {<br>
                    &nbsp;&nbsp;&nbsp;&nbsp;$.ajax({<br>
                    &nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;url: '/api/search',<br>
                    &nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;data: { q: query },<br>
                    &nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;success: function(data) {<br>
                    &nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;asyncResults(data);<br>
                    &nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;}<br>
                    &nbsp;&nbsp;&nbsp;&nbsp;});<br>
                    &nbsp;&nbsp;}<br>
                    });
                </code>
            </div>

            <div class="code-block">
                <div style="color: var(--text-tertiary); margin-bottom: 8px; margin-top: 16px;">4. Custom Template:</div>
                <code style="color: var(--info);">
                    $('#searchInput').typeahead(null, {<br>
                    &nbsp;&nbsp;display: 'name',<br>
                    &nbsp;&nbsp;templates: {<br>
                    &nbsp;&nbsp;&nbsp;&nbsp;empty: '&lt;div class="empty"&gt;No results&lt;/div&gt;',<br>
                    &nbsp;&nbsp;&nbsp;&nbsp;suggestion: function(data) {<br>
                    &nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;return '&lt;div&gt;' + data.name + '&lt;/div&gt;';<br>
                    &nbsp;&nbsp;&nbsp;&nbsp;}<br>
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
                        <span><strong>Typeahead.js</strong> - Twitter's implementation</span>
                    </li>
                    <li>
                        <i class="fa-solid fa-circle-check"></i>
                        <span><strong>Autocomplete.js</strong> - Lightweight, modern</span>
                    </li>
                    <li>
                        <i class="fa-solid fa-circle-check"></i>
                        <span><strong>Choices.js</strong> - Select + search</span>
                    </li>
                    <li>
                        <i class="fa-solid fa-circle-check"></i>
                        <span><strong>Custom Solution</strong> - Full control</span>
                    </li>
                </ul>
            </div>
        </div>
    </div>
</div>
@endsection

@push('scripts')
<script>
$(document).ready(function() {
    // Sample data for autocomplete
    const countries = [
        { name: 'Indonesia', code: '🇮🇩', capital: 'Jakarta', region: 'Southeast Asia', popular: true },
        { name: 'India', code: '🇮🇳', capital: 'New Delhi', region: 'South Asia' },
        { name: 'United States', code: '🇺🇸', capital: 'Washington D.C.', region: 'North America', popular: true },
        { name: 'United Kingdom', code: '🇬🇧', capital: 'London', region: 'Europe' },
        { name: 'France', code: '🇫🇷', capital: 'Paris', region: 'Europe' },
        { name: 'Germany', code: '🇩🇪', capital: 'Berlin', region: 'Europe' },
        { name: 'Japan', code: '🇯🇵', capital: 'Tokyo', region: 'East Asia' },
        { name: 'China', code: '🇨🇳', capital: 'Beijing', region: 'East Asia' },
        { name: 'Australia', code: '🇦🇺', capital: 'Canberra', region: 'Oceania' },
        { name: 'Brazil', code: '🇧🇷', capital: 'Brasília', region: 'South America' }
    ];

    const users = [
        { name: 'John Doe', email: 'john@example.com', role: 'Admin', online: true },
        { name: 'John Smith', email: 'john.smith@example.com', role: 'User' },
        { name: 'Johnny Cash', email: 'jcash@example.com', role: 'Editor' },
        { name: 'Jane Doe', email: 'jane@example.com', role: 'User', online: true },
        { name: 'Alice Johnson', email: 'alice@example.com', role: 'Admin' }
    ];

    const products = [
        { name: 'iPhone 15 Pro', brand: 'Apple', price: '$999', rating: '4.8', badge: 'New', icon: '📱' },
        { name: 'Samsung Galaxy S24', brand: 'Samsung', price: '$899', rating: '4.7', icon: '📱' },
        { name: 'Google Pixel 8', brand: 'Google', price: '$699', rating: '4.6', badge: 'Sale', icon: '📱' },
        { name: 'OnePlus 12', brand: 'OnePlus', price: '$799', rating: '4.5', icon: '📱' },
        { name: 'Xiaomi 14 Pro', brand: 'Xiaomi', price: '$649', rating: '4.4', icon: '📱' }
    ];

    const locations = [
        { name: 'Jakarta Pusat', city: 'DKI Jakarta', country: 'Indonesia', icon: '🏙️' },
        { name: 'Jakarta Selatan', city: 'DKI Jakarta', country: 'Indonesia', icon: '🏙️' },
        { name: 'Jakarta Barat', city: 'DKI Jakarta', country: 'Indonesia', icon: '🏙️' },
        { name: 'Jakarta Utara', city: 'DKI Jakarta', country: 'Indonesia', icon: '🏙️' },
        { name: 'Jakarta Timur', city: 'DKI Jakarta', country: 'Indonesia', icon: '🏙️' },
        { name: 'Surabaya', city: 'Jawa Timur', country: 'Indonesia', icon: '🏙️' },
        { name: 'Bandung', city: 'Jawa Barat', country: 'Indonesia', icon: '🏙️' }
    ];

    // Initialize all typeahead inputs
    $('.typeahead-input').each(function() {
        const $input = $(this);
        const $wrapper = $input.closest('.typeahead-wrapper');
        let $dropdown = $wrapper.find('.typeahead-dropdown');
        
        // Create dropdown if not exists
        if ($dropdown.length === 0) {
            $dropdown = $('<div class="typeahead-dropdown"></div>').appendTo($wrapper);
        }
        
        let debounceTimer;
        let activeIndex = -1;
        let currentItems = [];
        
        // Determine data source based on placeholder or label
        let dataSource = countries;
        const placeholder = $input.attr('placeholder').toLowerCase();
        const label = $wrapper.siblings('.ta-label').text().toLowerCase();
        
        if (placeholder.includes('user') || label.includes('user')) {
            dataSource = users;
        } else if (placeholder.includes('product') || label.includes('product')) {
            dataSource = products;
        } else if (placeholder.includes('address') || placeholder.includes('location') || label.includes('address') || label.includes('location')) {
            dataSource = locations;
        }
        
        // Input event handler
        $input.on('input', function() {
            const query = $(this).val().trim();
            
            // Clear previous timer
            clearTimeout(debounceTimer);
            
            if (query.length === 0) {
                hideDropdown();
                return;
            }
            
            // Debounce input
            debounceTimer = setTimeout(() => {
                filterAndShowDropdown(query, dataSource);
            }, 200);
        });
        
        // Focus event
        $input.on('focus', function() {
            const query = $(this).val().trim();
            if (query.length > 0) {
                filterAndShowDropdown(query, dataSource);
            }
        });
        
        // Keyboard navigation
        $input.on('keydown', function(e) {
            const $items = $dropdown.find('.typeahead-item');
            
            if (!$dropdown.hasClass('show') || $items.length === 0) {
                if (e.key === 'ArrowDown' || e.key === 'ArrowUp') {
                    e.preventDefault();
                    const query = $(this).val().trim();
                    if (query.length > 0) {
                        filterAndShowDropdown(query, dataSource);
                    }
                }
                return;
            }
            
            if (e.key === 'ArrowDown') {
                e.preventDefault();
                activeIndex = Math.min(activeIndex + 1, $items.length - 1);
                updateActiveItem($items);
            } else if (e.key === 'ArrowUp') {
                e.preventDefault();
                activeIndex = Math.max(activeIndex - 1, 0);
                updateActiveItem($items);
            } else if (e.key === 'Enter' && activeIndex >= 0) {
                e.preventDefault();
                const $activeItem = $items.eq(activeIndex);
                selectItem($activeItem, $input);
            } else if (e.key === 'Escape') {
                hideDropdown();
            }
        });
        
        // Click outside to close
        $(document).on('click', function(e) {
            if (!$wrapper.is(e.target) && $wrapper.has(e.target).length === 0) {
                hideDropdown();
            }
        });
        
        // Filter and show dropdown
        function filterAndShowDropdown(query, data) {
            const filtered = data.filter(item => 
                item.name.toLowerCase().includes(query.toLowerCase())
            );
            
            currentItems = filtered;
            activeIndex = -1;
            
            if (filtered.length === 0) {
                showEmptyState(query);
                return;
            }
            
            renderDropdown(filtered, query, dataSource);
            showDropdown();
        }
        
        // Render dropdown items
        function renderDropdown(items, query, source) {
            let html = `<div class="typeahead-header">${source === countries ? 'Countries' : source === users ? 'Users' : source === products ? 'Products' : 'Locations'} (${items.length} results)</div>`;
            
            items.forEach((item, index) => {
                const highlightedName = highlightMatch(item.name, query);
                
                html += `<div class="typeahead-item" data-index="${index}">`;
                html += `<div class="typeahead-item-icon">${item.code || item.icon || '👤'}</div>`;
                html += `<div class="typeahead-item-content">`;
                html += `<div class="typeahead-item-title">${highlightedName}</div>`;
                
                // Subtitle based on data source
                if (source === countries) {
                    html += `<div class="typeahead-item-subtitle">${item.capital} • ${item.region}</div>`;
                } else if (source === users) {
                    html += `<div class="typeahead-item-subtitle">${item.email} • ${item.role}</div>`;
                } else if (source === products) {
                    html += `<div class="typeahead-item-subtitle">${item.brand} • ${item.price} • ⭐ ${item.rating}</div>`;
                } else if (source === locations) {
                    html += `<div class="typeahead-item-subtitle">${item.city} • ${item.country}</div>`;
                }
                
                html += `</div>`;
                
                // Badge
                if (item.popular) {
                    html += `<div class="typeahead-item-badge">Popular</div>`;
                } else if (item.badge) {
                    const badgeClass = item.badge === 'Sale' ? 'style="background: var(--success-bg); color: var(--success);"' : '';
                    html += `<div class="typeahead-item-badge" ${badgeClass}>${item.badge}</div>`;
                } else if (item.online) {
                    html += `<div class="typeahead-item-badge">Online</div>`;
                }
                
                html += `</div>`;
            });
            
            html += `<div class="typeahead-footer">↑↓ Navigate • Enter Select • Esc Close</div>`;
            
            $dropdown.html(html);
            
            // Add click handlers
            $dropdown.find('.typeahead-item').on('click', function() {
                selectItem($(this), $input);
            });
        }
        
        // Show empty state
        function showEmptyState(query) {
            const sourceName = dataSource === countries ? 'countries' : dataSource === users ? 'users' : dataSource === products ? 'products' : 'locations';
            $dropdown.html(`
                <div class="typeahead-empty">
                    <i class="fa-solid fa-circle-xmark"></i>
                    <div>No ${sourceName} found for "${escapeHtml(query)}"</div>
                    <div style="font-size: 11px; margin-top: 4px;">Try different keywords</div>
                </div>
            `);
            showDropdown();
        }
        
        // Highlight matching text
        function highlightMatch(text, query) {
            const regex = new RegExp(`(${escapeRegex(query)})`, 'gi');
            return text.replace(regex, '<span class="highlight">$1</span>');
        }
        
        // Escape regex special characters
        function escapeRegex(string) {
            return string.replace(/[.*+?^${}()|[\]\\]/g, '\\$&');
        }
        
        // Escape HTML
        function escapeHtml(string) {
            const div = document.createElement('div');
            div.textContent = string;
            return div.innerHTML;
        }
        
        // Update active item
        function updateActiveItem($items) {
            $items.removeClass('active');
            if (activeIndex >= 0) {
                $items.eq(activeIndex).addClass('active');
                // Scroll into view if needed
                const itemEl = $items[activeIndex];
                if (itemEl) {
                    itemEl.scrollIntoView({ block: 'nearest' });
                }
            }
        }
        
        // Select item
        function selectItem($item, $input) {
            const title = $item.find('.typeahead-item-title').text();
            $input.val(title);
            hideDropdown();
            
            // Show success feedback
            Swal.fire({
                toast: true,
                position: 'top-end',
                icon: 'success',
                title: 'Selected: ' + title,
                showConfirmButton: false,
                timer: 2000
            });
        }
        
        // Show dropdown
        function showDropdown() {
            $dropdown.addClass('show');
        }
        
        // Hide dropdown
        function hideDropdown() {
            $dropdown.removeClass('show');
            activeIndex = -1;
        }
    });
    
    // Recent search functionality
    $('.recent-search').on('click', function(e) {
        if ($(e.target).hasClass('recent-search-remove')) {
            // Remove search
            $(this).fadeOut(200, function() {
                $(this).remove();
            });
        } else {
            // Search again
            const searchText = $(this).find('span').first().text();
            Swal.fire({
                toast: true,
                position: 'top-end',
                icon: 'info',
                title: 'Searching: ' + searchText,
                showConfirmButton: false,
                timer: 2000
            });
        }
    });
    
    // Advanced Search Form
    $('form').on('submit', function(e) {
        e.preventDefault();
        
        const $form = $(this);
        const formData = {};
        
        // Collect all form data
        $form.find('input[type="text"], input[type="number"]').each(function() {
            const $input = $(this);
            const placeholder = $input.attr('placeholder').toLowerCase();
            const value = $input.val().trim();
            
            if (value) {
                if (placeholder.includes('min')) {
                    formData.minPrice = value;
                } else if (placeholder.includes('max')) {
                    formData.maxPrice = value;
                } else if (placeholder.includes('rating')) {
                    formData.minRating = value;
                } else if (placeholder.includes('category')) {
                    formData.category = value;
                } else if (placeholder.includes('brand')) {
                    formData.brand = value;
                } else if (placeholder.includes('looking')) {
                    formData.searchQuery = value;
                }
            }
        });
        
        // Validate required field
        if (!formData.searchQuery) {
            Swal.fire({
                toast: true,
                position: 'top-end',
                icon: 'warning',
                title: 'Please enter a search query',
                showConfirmButton: false,
                timer: 2000
            });
            $form.find('input[placeholder="What are you looking for?"]').focus();
            return;
        }
        
        // Show search results
        let message = `Searching: "${formData.searchQuery}"`;
        if (formData.category) message += `\nCategory: ${formData.category}`;
        if (formData.brand) message += `\nBrand: ${formData.brand}`;
        if (formData.minPrice || formData.maxPrice) message += `\nPrice: $${formData.minPrice || '0'} - $${formData.maxPrice || '∞'}`;
        if (formData.minRating) message += `\nMin Rating: ${formData.minRating}⭐`;
        
        Swal.fire({
            icon: 'success',
            title: 'Search Submitted',
            html: `<pre style="text-align: left; font-size: 13px;">${message}</pre>`,
            confirmButtonText: 'OK',
            confirmButtonColor: 'var(--accent)'
        });
        
        console.log('Search Form Data:', formData);
    });
    
    // Form reset
    $('form').on('reset', function(e) {
        e.preventDefault();
        
        Swal.fire({
            toast: true,
            position: 'top-end',
            icon: 'info',
            title: 'Form reset',
            showConfirmButton: false,
            timer: 1500
        });
        
        // Clear all inputs
        $(this).find('input').val('');
    });
});
</script>
@endpush
