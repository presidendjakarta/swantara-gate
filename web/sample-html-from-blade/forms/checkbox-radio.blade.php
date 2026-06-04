@extends('layouts.app')

@section('title', 'Checkbox & Radio')

@push('styles')

@endpush

@section('content')
<div class="page-header">
    <div>
        <h1>Checkbox & Radio</h1>
        <p>Custom styled checkboxes, radio buttons, and toggle switches dengan Fluent Design</p>
    </div>
    <div style="display: flex; gap: 10px;">
        <button class="btn btn-secondary">
            <i class="fa-solid fa-book"></i>
            <span>Documentation</span>
        </button>
        <button class="btn btn-primary">
            <i class="fa-solid fa-palette"></i>
            <span>Customize</span>
        </button>
    </div>
</div>

<!-- Info Alert -->
<div class="content-card" style="margin-bottom: 24px; border-left: 4px solid var(--accent);">
    <div class="card-body" style="padding: 16px 20px;">
        <div style="display: flex; gap: 12px; align-items: start;">
            <i class="fa-solid fa-circle-info" style="color: var(--accent); font-size: 20px; margin-top: 2px;"></i>
            <div style="flex: 1;">
                <h4 style="margin-bottom: 4px; font-size: 14px;">Custom Styled Controls</h4>
                <p style="font-size: 13px; color: var(--text-secondary); margin: 0;">Semua checkbox dan radio button menggunakan custom CSS styling untuk tampilan yang konsisten dan modern. Mendukung keyboard navigation dan accessibility.</p>
            </div>
        </div>
    </div>
</div>

<!-- Basic Checkboxes -->
<div class="cr-section-title">
    <i class="fa-solid fa-square-check"></i>
    Basic Checkboxes <span class="badge badge-primary">Essential</span>
</div>

<div class="checkbox-radio-grid">
    <div class="content-card">
        <div class="card-header">
            <div class="card-header-left">
                <div class="card-icon bg-primary">
                    <i class="fa-solid fa-check-square"></i>
                </div>
                <div>
                    <h3>Default Checkboxes</h3>
                    <p class="card-subtitle">Basic checkbox states</p>
                </div>
            </div>
        </div>
        <div class="card-body">
            <div class="cr-group">
                <label class="custom-checkbox">
                    <input type="checkbox">
                    <span>Unchecked</span>
                </label>

                <label class="custom-checkbox">
                    <input type="checkbox" checked>
                    <span>Checked</span>
                </label>

                <label class="custom-checkbox">
                    <input type="checkbox" disabled>
                    <span>Disabled</span>
                </label>

                <label class="custom-checkbox">
                    <input type="checkbox" checked disabled>
                    <span>Checked & Disabled</span>
                </label>
            </div>

            <div class="divider"></div>

            <div class="form-group">
                <h4 style="font-size: 14px; margin-bottom: 12px;">Features:</h4>
                <ul class="feature-list">
                    <li>
                        <i class="fa-solid fa-circle-check"></i>
                        <span>Custom styled with CSS</span>
                    </li>
                    <li>
                        <i class="fa-solid fa-circle-check"></i>
                        <span>Hover effect</span>
                    </li>
                    <li>
                        <i class="fa-solid fa-circle-check"></i>
                        <span>Focus outline for accessibility</span>
                    </li>
                    <li>
                        <i class="fa-solid fa-circle-check"></i>
                        <span>Smooth transitions</span>
                    </li>
                </ul>
            </div>
        </div>
    </div>

    <div class="content-card">
        <div class="card-header">
            <div class="card-header-left">
                <div class="card-icon bg-success">
                    <i class="fa-solid fa-list-check"></i>
                </div>
                <div>
                    <h3>Checkbox Groups</h3>
                    <p class="card-subtitle">Multiple selection options</p>
                </div>
            </div>
        </div>
        <div class="card-body">
            <div class="form-group">
                <label style="font-weight: 600; margin-bottom: 8px; display: block;">Select Hobbies:</label>
                <div class="cr-group">
                    <label class="custom-checkbox">
                        <input type="checkbox" name="hobby" value="reading">
                        <span>📚 Reading</span>
                    </label>
                    <label class="custom-checkbox">
                        <input type="checkbox" name="hobby" value="sports" checked>
                        <span>⚽ Sports</span>
                    </label>
                    <label class="custom-checkbox">
                        <input type="checkbox" name="hobby" value="music">
                        <span>🎵 Music</span>
                    </label>
                    <label class="custom-checkbox">
                        <input type="checkbox" name="hobby" value="cooking" checked>
                        <span>🍳 Cooking</span>
                    </label>
                    <label class="custom-checkbox">
                        <input type="checkbox" name="hobby" value="traveling">
                        <span>✈️ Traveling</span>
                    </label>
                </div>
            </div>

            <div class="divider"></div>

            <div class="form-group">
                <label style="font-weight: 600; margin-bottom: 8px; display: block;">Inline Checkboxes:</label>
                <div class="cr-inline-group">
                    <label class="custom-checkbox">
                        <input type="checkbox">
                        <span>Option 1</span>
                    </label>
                    <label class="custom-checkbox">
                        <input type="checkbox" checked>
                        <span>Option 2</span>
                    </label>
                    <label class="custom-checkbox">
                        <input type="checkbox">
                        <span>Option 3</span>
                    </label>
                </div>
            </div>
        </div>
    </div>
</div>

<!-- Basic Radio Buttons -->
<div class="cr-section-title">
    <i class="fa-solid fa-circle-dot"></i>
    Basic Radio Buttons <span class="badge badge-success">Popular</span>
</div>

<div class="checkbox-radio-grid">
    <div class="content-card">
        <div class="card-header">
            <div class="card-header-left">
                <div class="card-icon bg-info">
                    <i class="fa-solid fa-dot-circle"></i>
                </div>
                <div>
                    <h3>Default Radio Buttons</h3>
                    <p class="card-subtitle">Single selection from group</p>
                </div>
            </div>
        </div>
        <div class="card-body">
            <div class="cr-group">
                <label class="custom-radio">
                    <input type="radio" name="default_radio">
                    <span>Unselected</span>
                </label>

                <label class="custom-radio">
                    <input type="radio" name="default_radio" checked>
                    <span>Selected</span>
                </label>

                <label class="custom-radio">
                    <input type="radio" name="default_radio" disabled>
                    <span>Disabled</span>
                </label>

                <label class="custom-radio">
                    <input type="radio" name="default_radio" checked disabled>
                    <span>Selected & Disabled</span>
                </label>
            </div>

            <div class="divider"></div>

            <div class="form-group">
                <h4 style="font-size: 14px; margin-bottom: 12px;">Features:</h4>
                <ul class="feature-list">
                    <li>
                        <i class="fa-solid fa-circle-check"></i>
                        <span>Only one selection per group</span>
                    </li>
                    <li>
                        <i class="fa-solid fa-circle-check"></i>
                        <span>Visual feedback on hover</span>
                    </li>
                    <li>
                        <i class="fa-solid fa-circle-check"></i>
                        <span>Keyboard navigation</span>
                    </li>
                    <li>
                        <i class="fa-solid fa-circle-check"></i>
                        <span>Accessible labels</span>
                    </li>
                </ul>
            </div>
        </div>
    </div>

    <div class="content-card">
        <div class="card-header">
            <div class="card-header-left">
                <div class="card-icon bg-warning">
                    <i class="fa-solid fa-list-ul"></i>
                </div>
                <div>
                    <h3>Radio Groups</h3>
                    <p class="card-subtitle">Categorized radio options</p>
                </div>
            </div>
        </div>
        <div class="card-body">
            <div class="form-group">
                <label style="font-weight: 600; margin-bottom: 8px; display: block;">Gender:</label>
                <div class="cr-group">
                    <label class="custom-radio">
                        <input type="radio" name="gender" value="male">
                        <span>👨 Male</span>
                    </label>
                    <label class="custom-radio">
                        <input type="radio" name="gender" value="female" checked>
                        <span>👩 Female</span>
                    </label>
                    <label class="custom-radio">
                        <input type="radio" name="gender" value="other">
                        <span>🧑 Other</span>
                    </label>
                </div>
            </div>

            <div class="divider"></div>

            <div class="form-group">
                <label style="font-weight: 600; margin-bottom: 8px; display: block;">Experience Level:</label>
                <div class="cr-group">
                    <label class="custom-radio">
                        <input type="radio" name="experience" value="beginner">
                        <span>🌱 Beginner</span>
                    </label>
                    <label class="custom-radio">
                        <input type="radio" name="experience" value="intermediate" checked>
                        <span>🌿 Intermediate</span>
                    </label>
                    <label class="custom-radio">
                        <input type="radio" name="experience" value="advanced">
                        <span>🌳 Advanced</span>
                    </label>
                    <label class="custom-radio">
                        <input type="radio" name="experience" value="expert">
                        <span>🏆 Expert</span>
                    </label>
                </div>
            </div>
        </div>
    </div>
</div>

<!-- Toggle Switches -->
<div class="cr-section-title">
    <i class="fa-solid fa-toggle-on"></i>
    Toggle Switches <span class="badge badge-warning">Modern</span>
</div>

<div class="checkbox-radio-grid">
    <div class="content-card">
        <div class="card-header">
            <div class="card-header-left">
                <div class="card-icon bg-primary">
                    <i class="fa-solid fa-power-off"></i>
                </div>
                <div>
                    <h3>Basic Toggles</h3>
                    <p class="card-subtitle">On/Off switch controls</p>
                </div>
            </div>
        </div>
        <div class="card-body">
            <div class="cr-group">
                <label class="toggle-switch">
                    <input type="checkbox">
                    <div>
                        <div class="toggle-label">Wi-Fi</div>
                        <div class="toggle-description">Enable wireless connection</div>
                    </div>
                </label>

                <label class="toggle-switch">
                    <input type="checkbox" checked>
                    <div>
                        <div class="toggle-label">Bluetooth</div>
                        <div class="toggle-description">Enable Bluetooth devices</div>
                    </div>
                </label>

                <label class="toggle-switch">
                    <input type="checkbox">
                    <div>
                        <div class="toggle-label">Airplane Mode</div>
                        <div class="toggle-description">Disable all connections</div>
                    </div>
                </label>

                <label class="toggle-switch">
                    <input type="checkbox" disabled>
                    <div>
                        <div class="toggle-label">Disabled Toggle</div>
                        <div class="toggle-description">This toggle is disabled</div>
                    </div>
                </label>
            </div>
        </div>
    </div>

    <div class="content-card">
        <div class="card-header">
            <div class="card-header-left">
                <div class="card-icon bg-success">
                    <i class="fa-solid fa-gear"></i>
                </div>
                <div>
                    <h3>Settings Toggles</h3>
                    <p class="card-subtitle">Application preferences</p>
                </div>
            </div>
        </div>
        <div class="card-body">
            <div class="cr-group">
                <label class="toggle-switch">
                    <input type="checkbox" checked>
                    <div>
                        <div class="toggle-label">🌙 Dark Mode</div>
                        <div class="toggle-description">Use dark theme</div>
                    </div>
                </label>

                <label class="toggle-switch">
                    <input type="checkbox" checked>
                    <div>
                        <div class="toggle-label">🔔 Notifications</div>
                        <div class="toggle-description">Show push notifications</div>
                    </div>
                </label>

                <label class="toggle-switch">
                    <input type="checkbox">
                    <div>
                        <div class="toggle-label">🔒 Two-Factor Auth</div>
                        <div class="toggle-description">Enable 2FA security</div>
                    </div>
                </label>

                <label class="toggle-switch">
                    <input type="checkbox" checked>
                    <div>
                        <div class="toggle-label">📧 Email Alerts</div>
                        <div class="toggle-description">Receive email notifications</div>
                    </div>
                </label>
            </div>

            <div class="divider"></div>

            <div class="form-group">
                <h4 style="font-size: 14px; margin-bottom: 12px;">Toggle Features:</h4>
                <ul class="feature-list">
                    <li>
                        <i class="fa-solid fa-circle-check"></i>
                        <span>Smooth sliding animation</span>
                    </li>
                    <li>
                        <i class="fa-solid fa-circle-check"></i>
                        <span>Visual ON/OFF states</span>
                    </li>
                    <li>
                        <i class="fa-solid fa-circle-check"></i>
                        <span>Label + description layout</span>
                    </li>
                    <li>
                        <i class="fa-solid fa-circle-check"></i>
                        <span>Accessible keyboard control</span>
                    </li>
                </ul>
            </div>
        </div>
    </div>
</div>

<!-- Checkbox Cards -->
<div class="cr-section-title">
    <i class="fa-solid fa-credit-card"></i>
    Checkbox Cards
</div>

<div class="checkbox-radio-grid">
    <div class="content-card">
        <div class="card-header">
            <div class="card-header-left">
                <div class="card-icon bg-info">
                    <i class="fa-solid fa-layer-group"></i>
                </div>
                <div>
                    <h3>Selection Cards</h3>
                    <p class="card-subtitle">Visual card-based selection</p>
                </div>
            </div>
        </div>
        <div class="card-body">
            <div style="display: grid; grid-template-columns: 1fr 1fr; gap: 12px;">
                <div class="checkbox-card selected" onclick="this.classList.toggle('selected'); this.querySelector('input').checked = !this.querySelector('input').checked;">
                    <input type="checkbox" checked hidden>
                    <div class="checkbox-card-title">💎 Premium</div>
                    <div class="checkbox-card-desc">Full access to all features</div>
                    <div style="margin-top: 8px; font-size: 16px; font-weight: 700; color: var(--accent);">$99/mo</div>
                </div>

                <div class="checkbox-card" onclick="this.classList.toggle('selected'); this.querySelector('input[type=hidden]').checked = !this.querySelector('input[type=hidden]').checked;">
                    <input type="checkbox" hidden>
                    <div class="checkbox-card-title">🥈 Standard</div>
                    <div class="checkbox-card-desc">Basic features included</div>
                    <div style="margin-top: 8px; font-size: 16px; font-weight: 700; color: var(--text-secondary);">$49/mo</div>
                </div>

                <div class="checkbox-card" onclick="this.classList.toggle('selected'); this.querySelector('input[type=hidden]').checked = !this.querySelector('input[type=hidden]').checked;">
                    <input type="checkbox" hidden>
                    <div class="checkbox-card-title">🥉 Basic</div>
                    <div class="checkbox-card-desc">Limited features</div>
                    <div style="margin-top: 8px; font-size: 16px; font-weight: 700; color: var(--text-secondary);">$19/mo</div>
                </div>

                <div class="checkbox-card" onclick="this.classList.toggle('selected'); this.querySelector('input[type=hidden]').checked = !this.querySelector('input[type=hidden]').checked;">
                    <input type="checkbox" hidden>
                    <div class="checkbox-card-title">🆓 Free</div>
                    <div class="checkbox-card-desc">Try before you buy</div>
                    <div style="margin-top: 8px; font-size: 16px; font-weight: 700; color: var(--success);">Free</div>
                </div>
            </div>
        </div>
    </div>

    <div class="content-card">
        <div class="card-header">
            <div class="card-header-left">
                <div class="card-icon bg-warning">
                    <i class="fa-solid fa-calendar-check"></i>
                </div>
                <div>
                    <h3>Day Selector Cards</h3>
                    <p class="card-subtitle">Select multiple days</p>
                </div>
            </div>
        </div>
        <div class="card-body">
            <div style="display: grid; grid-template-columns: repeat(7, 1fr); gap: 8px;">
                <div class="checkbox-card selected" style="text-align: center; padding: 12px 8px;" onclick="this.classList.toggle('selected');">
                    <input type="checkbox" hidden checked>
                    <div style="font-weight: 600; font-size: 13px;">Mon</div>
                </div>
                <div class="checkbox-card selected" style="text-align: center; padding: 12px 8px;" onclick="this.classList.toggle('selected');">
                    <input type="checkbox" hidden checked>
                    <div style="font-weight: 600; font-size: 13px;">Tue</div>
                </div>
                <div class="checkbox-card" style="text-align: center; padding: 12px 8px;" onclick="this.classList.toggle('selected');">
                    <input type="checkbox" hidden>
                    <div style="font-weight: 600; font-size: 13px;">Wed</div>
                </div>
                <div class="checkbox-card selected" style="text-align: center; padding: 12px 8px;" onclick="this.classList.toggle('selected');">
                    <input type="checkbox" hidden checked>
                    <div style="font-weight: 600; font-size: 13px;">Thu</div>
                </div>
                <div class="checkbox-card" style="text-align: center; padding: 12px 8px;" onclick="this.classList.toggle('selected');">
                    <input type="checkbox" hidden>
                    <div style="font-weight: 600; font-size: 13px;">Fri</div>
                </div>
                <div class="checkbox-card" style="text-align: center; padding: 12px 8px;" onclick="this.classList.toggle('selected');">
                    <input type="checkbox" hidden>
                    <div style="font-weight: 600; font-size: 13px;">Sat</div>
                </div>
                <div class="checkbox-card" style="text-align: center; padding: 12px 8px;" onclick="this.classList.toggle('selected');">
                    <input type="checkbox" hidden>
                    <div style="font-weight: 600; font-size: 13px;">Sun</div>
                </div>
            </div>

            <div class="divider"></div>

            <div class="form-group">
                <label style="font-weight: 600; margin-bottom: 8px; display: block;">Select Payment Method:</label>
                <div class="cr-group">
                    <label class="custom-radio">
                        <input type="radio" name="payment" value="credit" checked>
                        <span>💳 Credit Card</span>
                    </label>
                    <label class="custom-radio">
                        <input type="radio" name="payment" value="paypal">
                        <span>🅿️ PayPal</span>
                    </label>
                    <label class="custom-radio">
                        <input type="radio" name="payment" value="bank">
                        <span>🏦 Bank Transfer</span>
                    </label>
                    <label class="custom-radio">
                        <input type="radio" name="payment" value="crypto">
                        <span>₿ Cryptocurrency</span>
                    </label>
                </div>
            </div>
        </div>
    </div>
</div>

<!-- Advanced Examples -->
<div class="cr-section-title">
    <i class="fa-solid fa-wand-magic-sparkles"></i>
    Advanced Examples
</div>

<div class="checkbox-radio-grid">
    <div class="content-card">
        <div class="card-header">
            <div class="card-header-left">
                <div class="card-icon bg-primary">
                    <i class="fa-solid fa-star"></i>
                </div>
                <div>
                    <h3>Rating System</h3>
                    <p class="card-subtitle">Star rating with radio buttons</p>
                </div>
            </div>
        </div>
        <div class="card-body">
            <div class="form-group">
                <label style="font-weight: 600; margin-bottom: 12px; display: block;">Rate this product:</label>
                <div style="display: flex; gap: 8px; font-size: 32px;">
                    <label style="cursor: pointer; transition: transform 0.2s;" onmouseover="this.style.transform='scale(1.2)'" onmouseout="this.style.transform='scale(1)'">
                        <input type="radio" name="rating" value="1" style="display: none;">
                        <span>⭐</span>
                    </label>
                    <label style="cursor: pointer; transition: transform 0.2s;" onmouseover="this.style.transform='scale(1.2)'" onmouseout="this.style.transform='scale(1)'">
                        <input type="radio" name="rating" value="2" style="display: none;">
                        <span>⭐</span>
                    </label>
                    <label style="cursor: pointer; transition: transform 0.2s;" onmouseover="this.style.transform='scale(1.2)'" onmouseout="this.style.transform='scale(1)'">
                        <input type="radio" name="rating" value="3" checked style="display: none;">
                        <span>⭐</span>
                    </label>
                    <label style="cursor: pointer; transition: transform 0.2s;" onmouseover="this.style.transform='scale(1.2)'" onmouseout="this.style.transform='scale(1)'">
                        <input type="radio" name="rating" value="4" style="display: none;">
                        <span>⭐</span>
                    </label>
                    <label style="cursor: pointer; transition: transform 0.2s;" onmouseover="this.style.transform='scale(1.2)'" onmouseout="this.style.transform='scale(1)'">
                        <input type="radio" name="rating" value="5" style="display: none;">
                        <span>⭐</span>
                    </label>
                </div>
                <div class="helper-text">
                    <i class="fa-solid fa-circle-info"></i>
                    Current rating: 3/5 stars
                </div>
            </div>

            <div class="divider"></div>

            <div class="form-group">
                <label style="font-weight: 600; margin-bottom: 12px; display: block;">Satisfaction Level:</label>
                <div class="cr-inline-group">
                    <label class="custom-radio">
                        <input type="radio" name="satisfaction" value="very-bad">
                        <span>😞 Very Bad</span>
                    </label>
                    <label class="custom-radio">
                        <input type="radio" name="satisfaction" value="bad">
                        <span>😕 Bad</span>
                    </label>
                    <label class="custom-radio">
                        <input type="radio" name="satisfaction" value="neutral" checked>
                        <span>😐 Neutral</span>
                    </label>
                    <label class="custom-radio">
                        <input type="radio" name="satisfaction" value="good">
                        <span>😊 Good</span>
                    </label>
                    <label class="custom-radio">
                        <input type="radio" name="satisfaction" value="excellent">
                        <span>😄 Excellent</span>
                    </label>
                </div>
            </div>
        </div>
    </div>

    <div class="content-card">
        <div class="card-header">
            <div class="card-header-left">
                <div class="card-icon bg-success">
                    <i class="fa-solid fa-file-signature"></i>
                </div>
                <div>
                    <h3>Agreement Checkboxes</h3>
                    <p class="card-subtitle">Terms and conditions</p>
                </div>
            </div>
        </div>
        <div class="card-body">
            <div class="form-group">
                <div class="cr-group">
                    <label class="custom-checkbox">
                        <input type="checkbox" required>
                        <span>I agree to the <a href="#" style="color: var(--accent);">Terms of Service</a></span>
                    </label>

                    <label class="custom-checkbox">
                        <input type="checkbox" required>
                        <span>I accept the <a href="#" style="color: var(--accent);">Privacy Policy</a></span>
                    </label>

                    <label class="custom-checkbox">
                        <input type="checkbox">
                        <span>Subscribe to newsletter</span>
                    </label>

                    <label class="custom-checkbox">
                        <input type="checkbox" checked>
                        <span>Enable two-factor authentication</span>
                    </label>
                </div>
            </div>

            <div class="divider"></div>

            <div class="form-group">
                <label style="font-weight: 600; margin-bottom: 12px; display: block;">Notification Preferences:</label>
                <div class="cr-group">
                    <label class="toggle-switch">
                        <input type="checkbox" checked>
                        <div>
                            <div class="toggle-label">📧 Email Notifications</div>
                        </div>
                    </label>

                    <label class="toggle-switch">
                        <input type="checkbox" checked>
                        <div>
                            <div class="toggle-label">📱 SMS Notifications</div>
                        </div>
                    </label>

                    <label class="toggle-switch">
                        <input type="checkbox">
                        <div>
                            <div class="toggle-label">🔔 Push Notifications</div>
                        </div>
                    </label>
                </div>
            </div>
        </div>
    </div>
</div>

<!-- Complete Form Example -->
<div class="cr-section-title">
    <i class="fa-solid fa-clipboard-list"></i>
    Complete Form Example
</div>

<div class="checkbox-radio-grid full-width">
    <div class="content-card">
        <div class="card-header">
            <div class="card-header-left">
                <div class="card-icon bg-primary">
                    <i class="fa-solid fa-user-plus"></i>
                </div>
                <div>
                    <h3>Registration Form</h3>
                    <p class="card-subtitle">Complete form with checkboxes, radios, and toggles</p>
                </div>
            </div>
        </div>
        <div class="card-body">
            <form style="max-width: 800px;">
                <div class="form-group">
                    <label style="font-weight: 600; margin-bottom: 8px; display: block;">Account Type <span style="color: var(--danger);">*</span></label>
                    <div class="cr-inline-group">
                        <label class="custom-radio">
                            <input type="radio" name="account_type" value="personal" checked>
                            <span>👤 Personal</span>
                        </label>
                        <label class="custom-radio">
                            <input type="radio" name="account_type" value="business">
                            <span>💼 Business</span>
                        </label>
                        <label class="custom-radio">
                            <input type="radio" name="account_type" value="enterprise">
                            <span>🏢 Enterprise</span>
                        </label>
                    </div>
                </div>

                <div class="form-group">
                    <label style="font-weight: 600; margin-bottom: 8px; display: block;">Interests (Select all that apply)</label>
                    <div style="display: grid; grid-template-columns: repeat(3, 1fr); gap: 8px;">
                        <label class="custom-checkbox">
                            <input type="checkbox" name="interest" value="technology">
                            <span>💻 Technology</span>
                        </label>
                        <label class="custom-checkbox">
                            <input type="checkbox" name="interest" value="design" checked>
                            <span>🎨 Design</span>
                        </label>
                        <label class="custom-checkbox">
                            <input type="checkbox" name="interest" value="marketing">
                            <span>📈 Marketing</span>
                        </label>
                        <label class="custom-checkbox">
                            <input type="checkbox" name="interest" value="finance" checked>
                            <span>💰 Finance</span>
                        </label>
                        <label class="custom-checkbox">
                            <input type="checkbox" name="interest" value="health">
                            <span>🏥 Health</span>
                        </label>
                        <label class="custom-checkbox">
                            <input type="checkbox" name="interest" value="education">
                            <span>📚 Education</span>
                        </label>
                    </div>
                </div>

                <div class="divider"></div>

                <div class="form-group">
                    <label style="font-weight: 600; margin-bottom: 12px; display: block;">Preferences</label>
                    <div class="cr-group">
                        <label class="toggle-switch">
                            <input type="checkbox" checked>
                            <div>
                                <div class="toggle-label">🌙 Dark Mode</div>
                                <div class="toggle-description">Use dark theme for the interface</div>
                            </div>
                        </label>

                        <label class="toggle-switch">
                            <input type="checkbox" checked>
                            <div>
                                <div class="toggle-label">🔔 Email Notifications</div>
                                <div class="toggle-description">Receive updates via email</div>
                            </div>
                        </label>

                        <label class="toggle-switch">
                            <input type="checkbox">
                            <div>
                                <div class="toggle-label">📱 SMS Notifications</div>
                                <div class="toggle-description">Receive updates via SMS</div>
                            </div>
                        </label>
                    </div>
                </div>

                <div class="divider"></div>

                <div class="form-group">
                    <label style="font-weight: 600; margin-bottom: 12px; display: block;">Agreements <span style="color: var(--danger);">*</span></label>
                    <div class="cr-group">
                        <label class="custom-checkbox">
                            <input type="checkbox" required>
                            <span>I agree to the <a href="#" style="color: var(--accent);">Terms of Service</a> and <a href="#" style="color: var(--accent);">Privacy Policy</a></span>
                        </label>

                        <label class="custom-checkbox">
                            <input type="checkbox">
                            <span>I want to receive marketing emails</span>
                        </label>

                        <label class="custom-checkbox">
                            <input type="checkbox" checked>
                            <span>I agree to share anonymous usage data to improve the service</span>
                        </label>
                    </div>
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

<!-- Comparison Summary -->
<div class="checkbox-radio-grid full-width">
    <div class="content-card">
        <div class="card-header">
            <div class="card-header-left">
                <div class="card-icon bg-info">
                    <i class="fa-solid fa-circle-question"></i>
                </div>
                <div>
                    <h3>Checkbox vs Radio vs Toggle</h3>
                    <p class="card-subtitle">When to use each control type</p>
                </div>
            </div>
        </div>
        <div class="card-body">
            <div class="checkbox-radio-grid three-cols">
                <div>
                    <h4 style="font-size: 14px; margin-bottom: 12px; display: flex; align-items: center; gap: 8px;">
                        <span style="font-size: 24px;">☑️</span>
                        Checkboxes
                    </h4>
                    <ul class="feature-list">
                        <li><i class="fa-solid fa-check"></i> Multiple selections allowed</li>
                        <li><i class="fa-solid fa-check"></i> Independent options</li>
                        <li><i class="fa-solid fa-check"></i> Can be unchecked</li>
                        <li><i class="fa-solid fa-check"></i> Use for "select all that apply"</li>
                    </ul>
                </div>
                <div>
                    <h4 style="font-size: 14px; margin-bottom: 12px; display: flex; align-items: center; gap: 8px;">
                        <span style="font-size: 24px;">🔘</span>
                        Radio Buttons
                    </h4>
                    <ul class="feature-list">
                        <li><i class="fa-solid fa-check"></i> Single selection only</li>
                        <li><i class="fa-solid fa-check"></i> Mutually exclusive options</li>
                        <li><i class="fa-solid fa-check"></i> Group with same name</li>
                        <li><i class="fa-solid fa-check"></i> Use for "choose one"</li>
                    </ul>
                </div>
                <div>
                    <h4 style="font-size: 14px; margin-bottom: 12px; display: flex; align-items: center; gap: 8px;">
                        <span style="font-size: 24px;">🔛</span>
                        Toggle Switches
                    </h4>
                    <ul class="feature-list">
                        <li><i class="fa-solid fa-check"></i> ON/OFF states</li>
                        <li><i class="fa-solid fa-check"></i> Immediate action</li>
                        <li><i class="fa-solid fa-check"></i> Settings & preferences</li>
                        <li><i class="fa-solid fa-check"></i> Use for "enable/disable"</li>
                    </ul>
                </div>
            </div>
        </div>
    </div>
</div>
@endsection
