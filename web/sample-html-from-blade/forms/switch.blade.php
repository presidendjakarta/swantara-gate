@extends('layouts.app')

@section('title', 'Switch')

@push('styles')
@endpush

@section('content')
<div class="page-header">
    <div>
        <h1>Switch</h1>
        <p>Toggle switch components untuk on/off states dengan berbagai style dan ukuran</p>
    </div>
    <div style="display: flex; gap: 10px;">
        <button class="btn btn-secondary">
            <i class="fa-solid fa-book"></i>
            <span>Documentation</span>
        </button>
        <button class="btn btn-primary">
            <i class="fa-solid fa-toggle-on"></i>
            <span>Try Switch</span>
        </button>
    </div>
</div>

<!-- Info Alert -->
<div class="content-card" style="margin-bottom: 24px; border-left: 4px solid var(--accent);">
    <div class="card-body" style="padding: 16px 20px;">
        <div style="display: flex; gap: 12px; align-items: start;">
            <i class="fa-solid fa-circle-info" style="color: var(--accent); font-size: 20px; margin-top: 2px;"></i>
            <div style="flex: 1;">
                <h4 style="margin-bottom: 4px; font-size: 14px;">What is Toggle Switch?</h4>
                <p style="font-size: 13px; color: var(--text-secondary); margin: 0;">Toggle switches represent binary states (on/off, true/false). They provide immediate visual feedback and are commonly used for settings, preferences, and feature activation. Pure CSS implementation with smooth animations.</p>
            </div>
        </div>
    </div>
</div>

<!-- Basic Toggle Switches -->
<div class="sw-section-title">
    <i class="fa-solid fa-toggle-on"></i>
    Basic Toggle Switches <span class="badge badge-primary">Essential</span>
</div>

<div class="switch-grid">
    <div class="content-card">
        <div class="card-header">
            <div class="card-header-left">
                <div class="card-icon bg-primary">
                    <i class="fa-solid fa-toggle-on"></i>
                </div>
                <div>
                    <h3>Default Switches</h3>
                    <p class="card-subtitle">Basic on/off toggles</p>
                </div>
            </div>
        </div>
        <div class="card-body">
            <div class="sw-example">
                <label class="sw-label">Dark Mode</label>
                <div class="toggle-switch">
                    <input type="checkbox" id="darkMode" checked>
                    <label for="darkMode" class="toggle-switch-label">Enable Dark Mode</label>
                </div>
                <div class="sw-helper">
                    <i class="fa-solid fa-circle-info"></i>
                    Currently: ON
                </div>
            </div>

            <div class="sw-example">
                <label class="sw-label">Notifications</label>
                <div class="toggle-switch">
                    <input type="checkbox" id="notifications">
                    <label for="notifications" class="toggle-switch-label">Enable Notifications</label>
                </div>
                <div class="sw-helper">
                    <i class="fa-solid fa-circle-info"></i>
                    Currently: OFF
                </div>
            </div>

            <div class="sw-example">
                <label class="sw-label">Auto-Save</label>
                <div class="toggle-switch">
                    <input type="checkbox" id="autoSave" checked>
                    <label for="autoSave" class="toggle-switch-label">Auto-Save Documents</label>
                </div>
                <div class="sw-helper">
                    <i class="fa-solid fa-circle-info"></i>
                    Currently: ON
                </div>
            </div>

            <div class="sw-example">
                <label class="sw-label">Two-Factor Auth</label>
                <div class="toggle-switch">
                    <input type="checkbox" id="tfa" checked>
                    <label for="tfa" class="toggle-switch-label">Two-Factor Authentication</label>
                </div>
                <div class="status-indicator status-on">
                    <i class="fa-solid fa-check"></i>
                    Active
                </div>
            </div>

            <div class="divider"></div>

            <div class="helper-text">
                <i class="fa-solid fa-circle-info"></i>
                Click the switch to toggle on/off state
            </div>
        </div>
    </div>

    <div class="content-card">
        <div class="card-header">
            <div class="card-header-left">
                <div class="card-icon bg-success">
                    <i class="fa-solid fa-sliders"></i>
                </div>
                <div>
                    <h3>Size Variants</h3>
                    <p class="card-subtitle">Small, medium, and large</p>
                </div>
            </div>
        </div>
        <div class="card-body">
            <div class="sw-example">
                <label class="sw-label">Small Switch</label>
                <div class="toggle-switch toggle-switch-sm">
                    <input type="checkbox" id="small1" checked>
                    <label for="small1" class="toggle-switch-label">Small (40px)</label>
                </div>
                <div class="sw-helper">
                    <i class="fa-solid fa-circle-info"></i>
                    Size: 40x22px
                </div>
            </div>

            <div class="sw-example">
                <label class="sw-label">Medium Switch (Default)</label>
                <div class="toggle-switch">
                    <input type="checkbox" id="medium1" checked>
                    <label for="medium1" class="toggle-switch-label">Medium (52px)</label>
                </div>
                <div class="sw-helper">
                    <i class="fa-solid fa-circle-info"></i>
                    Size: 52x28px
                </div>
            </div>

            <div class="sw-example">
                <label class="sw-label">Large Switch</label>
                <div class="toggle-switch toggle-switch-lg">
                    <input type="checkbox" id="large1" checked>
                    <label for="large1" class="toggle-switch-label">Large (64px)</label>
                </div>
                <div class="sw-helper">
                    <i class="fa-solid fa-circle-info"></i>
                    Size: 64x34px
                </div>
            </div>

            <div class="divider"></div>

            <div class="form-group">
                <h4 style="font-size: 14px; margin-bottom: 12px;">Size Guide:</h4>
                <ul class="feature-list">
                    <li>
                        <i class="fa-solid fa-circle-check"></i>
                        <span>Small: Dense forms, tables</span>
                    </li>
                    <li>
                        <i class="fa-solid fa-circle-check"></i>
                        <span>Medium: Standard use</span>
                    </li>
                    <li>
                        <i class="fa-solid fa-circle-check"></i>
                        <span>Large: Touchscreens, accessibility</span>
                    </li>
                </ul>
            </div>
        </div>
    </div>
</div>

<!-- Color Variants -->
<div class="sw-section-title">
    <i class="fa-solid fa-palette"></i>
    Color Variants
</div>

<div class="switch-grid four-cols">
    <div class="content-card">
        <div class="card-header">
            <div class="card-header-left">
                <div class="card-icon bg-primary">
                    <i class="fa-solid fa-circle"></i>
                </div>
                <div>
                    <h3>Primary</h3>
                    <p class="card-subtitle">Default accent color</p>
                </div>
            </div>
        </div>
        <div class="card-body">
            <div class="sw-example">
                <label class="sw-label">Primary Theme</label>
                <div class="toggle-switch">
                    <input type="checkbox" id="primary1" checked>
                    <label for="primary1" class="toggle-switch-label">Accent Blue</label>
                </div>
            </div>

            <div class="sw-example">
                <div class="toggle-switch">
                    <input type="checkbox" id="primary2" checked>
                    <label for="primary2" class="toggle-switch-label">Enable Feature</label>
                </div>
            </div>

            <div class="divider"></div>

            <div class="helper-text">
                <i class="fa-solid fa-circle-info"></i>
                Default switch color
            </div>
        </div>
    </div>

    <div class="content-card">
        <div class="card-header">
            <div class="card-header-left">
                <div class="card-icon bg-success">
                    <i class="fa-solid fa-circle"></i>
                </div>
                <div>
                    <h3>Success</h3>
                    <p class="card-subtitle">Green active state</p>
                </div>
            </div>
        </div>
        <div class="card-body">
            <div class="sw-example">
                <label class="sw-label">Success Theme</label>
                <div class="toggle-switch toggle-switch-success">
                    <input type="checkbox" id="success1" checked>
                    <label for="success1" class="toggle-switch-label">Published</label>
                </div>
            </div>

            <div class="sw-example">
                <div class="toggle-switch toggle-switch-success">
                    <input type="checkbox" id="success2" checked>
                    <label for="success2" class="toggle-switch-label">Verified</label>
                </div>
            </div>

            <div class="divider"></div>

            <div class="helper-text">
                <i class="fa-solid fa-circle-info"></i>
                Use for positive actions
            </div>
        </div>
    </div>

    <div class="content-card">
        <div class="card-header">
            <div class="card-header-left">
                <div class="card-icon bg-danger">
                    <i class="fa-solid fa-circle"></i>
                </div>
                <div>
                    <h3>Danger</h3>
                    <p class="card-subtitle">Red active state</p>
                </div>
            </div>
        </div>
        <div class="card-body">
            <div class="sw-example">
                <label class="sw-label">Danger Theme</label>
                <div class="toggle-switch toggle-switch-danger">
                    <input type="checkbox" id="danger1" checked>
                    <label for="danger1" class="toggle-switch-label">Block User</label>
                </div>
            </div>

            <div class="sw-example">
                <div class="toggle-switch toggle-switch-danger">
                    <input type="checkbox" id="danger2" checked>
                    <label for="danger2" class="toggle-switch-label">Delete Mode</label>
                </div>
            </div>

            <div class="divider"></div>

            <div class="helper-text">
                <i class="fa-solid fa-circle-info"></i>
                Use for destructive actions
            </div>
        </div>
    </div>

    <div class="content-card">
        <div class="card-header">
            <div class="card-header-left">
                <div class="card-icon bg-warning">
                    <i class="fa-solid fa-circle"></i>
                </div>
                <div>
                    <h3>Warning & Info</h3>
                    <p class="card-subtitle">Yellow and cyan</p>
                </div>
            </div>
        </div>
        <div class="card-body">
            <div class="sw-example">
                <label class="sw-label">Warning Theme</label>
                <div class="toggle-switch toggle-switch-warning">
                    <input type="checkbox" id="warning1" checked>
                    <label for="warning1" class="toggle-switch-label">Maintenance</label>
                </div>
            </div>

            <div class="sw-example">
                <label class="sw-label">Info Theme</label>
                <div class="toggle-switch toggle-switch-info">
                    <input type="checkbox" id="info1" checked>
                    <label for="info1" class="toggle-switch-label">Sync Mode</label>
                </div>
            </div>

            <div class="divider"></div>

            <div class="helper-text">
                <i class="fa-solid fa-circle-info"></i>
                Context-specific colors
            </div>
        </div>
    </div>
</div>

<!-- Switches with Icons -->
<div class="sw-section-title">
    <i class="fa-solid fa-icons"></i>
    Switches with Icons
</div>

<div class="switch-grid">
    <div class="content-card">
        <div class="card-header">
            <div class="card-header-left">
                <div class="card-icon bg-info">
                    <i class="fa-solid fa-icons"></i>
                </div>
                <div>
                    <h3>Check/Cross Icons</h3>
                    <p class="card-subtitle">✓ and ✕ inside switch</p>
                </div>
            </div>
        </div>
        <div class="card-body">
            <div class="sw-example">
                <label class="sw-label">With Icons (ON)</label>
                <div class="toggle-switch toggle-with-icons">
                    <input type="checkbox" id="iconOn" checked>
                    <label for="iconOn" class="toggle-switch-label">Feature Enabled</label>
                </div>
                <div class="status-indicator status-on">
                    <i class="fa-solid fa-check"></i>
                    ON
                </div>
            </div>

            <div class="sw-example">
                <label class="sw-label">With Icons (OFF)</label>
                <div class="toggle-switch toggle-with-icons">
                    <input type="checkbox" id="iconOff">
                    <label for="iconOff" class="toggle-switch-label">Feature Disabled</label>
                </div>
                <div class="status-indicator status-off">
                    <i class="fa-solid fa-xmark"></i>
                    OFF
                </div>
            </div>

            <div class="divider"></div>

            <div class="helper-text">
                <i class="fa-solid fa-circle-info"></i>
                Icons provide clear visual feedback
            </div>
        </div>
    </div>

    <div class="content-card">
        <div class="card-header">
            <div class="card-header-left">
                <div class="card-icon bg-warning">
                    <i class="fa-solid fa-align-left"></i>
                </div>
                <div>
                    <h3>Label Position</h3>
                    <p class="card-subtitle">Left and right labels</p>
                </div>
            </div>
        </div>
        <div class="card-body">
            <div class="sw-example">
                <label class="sw-label">Label on Right (Default)</label>
                <div class="toggle-switch">
                    <input type="checkbox" id="rightLabel" checked>
                    <label for="rightLabel" class="toggle-switch-label">Label Right</label>
                </div>
            </div>

            <div class="sw-example">
                <label class="sw-label">Label on Left</label>
                <div class="toggle-switch toggle-label-left">
                    <input type="checkbox" id="leftLabel" checked>
                    <label for="leftLabel" class="toggle-switch-label">Label Left</label>
                </div>
            </div>

            <div class="divider"></div>

            <div class="form-group">
                <h4 style="font-size: 14px; margin-bottom: 12px;">Position Options:</h4>
                <ul class="feature-list">
                    <li>
                        <i class="fa-solid fa-circle-check"></i>
                        <span>Right label (default)</span>
                    </li>
                    <li>
                        <i class="fa-solid fa-circle-check"></i>
                        <span>Left label (RTL support)</span>
                    </li>
                </ul>
            </div>
        </div>
    </div>
</div>

<!-- List Switches -->
<div class="sw-section-title">
    <i class="fa-solid fa-list"></i>
    List Switches <span class="badge badge-success">Popular</span>
</div>

<div class="switch-grid full-width">
    <div class="content-card">
        <div class="card-header">
            <div class="card-header-left">
                <div class="card-icon bg-primary">
                    <i class="fa-solid fa-list-check"></i>
                </div>
                <div>
                    <h3>Settings List</h3>
                    <p class="card-subtitle">Row-based toggle switches</p>
                </div>
            </div>
        </div>
        <div class="card-body">
            <div style="display: flex; flex-direction: column; gap: 12px;">
                <div class="list-switch">
                    <div class="list-switch-content">
                        <div class="list-switch-title">🌙 Dark Mode</div>
                        <div class="list-switch-desc">Switch to dark theme</div>
                    </div>
                    <div class="toggle-switch">
                        <input type="checkbox" id="list1" checked>
                    </div>
                </div>

                <div class="list-switch">
                    <div class="list-switch-content">
                        <div class="list-switch-title">🔔 Notifications</div>
                        <div class="list-switch-desc">Receive push notifications</div>
                    </div>
                    <div class="toggle-switch">
                        <input type="checkbox" id="list2" checked>
                    </div>
                </div>

                <div class="list-switch">
                    <div class="list-switch-content">
                        <div class="list-switch-title">📧 Email Alerts</div>
                        <div class="list-switch-desc">Get email notifications</div>
                    </div>
                    <div class="toggle-switch">
                        <input type="checkbox" id="list3">
                    </div>
                </div>

                <div class="list-switch">
                    <div class="list-switch-content">
                        <div class="list-switch-title">💾 Auto-Save</div>
                        <div class="list-switch-desc">Automatically save changes</div>
                    </div>
                    <div class="toggle-switch">
                        <input type="checkbox" id="list4" checked>
                    </div>
                </div>

                <div class="list-switch">
                    <div class="list-switch-content">
                        <div class="list-switch-title">🔒 Two-Factor Auth</div>
                        <div class="list-switch-desc">Enhanced security</div>
                    </div>
                    <div class="toggle-switch">
                        <input type="checkbox" id="list5" checked>
                    </div>
                </div>

                <div class="list-switch">
                    <div class="list-switch-content">
                        <div class="list-switch-title">🌐 Language Detection</div>
                        <div class="list-switch-desc">Auto-detect browser language</div>
                    </div>
                    <div class="toggle-switch">
                        <input type="checkbox" id="list6">
                    </div>
                </div>
            </div>

            <div class="divider"></div>

            <div class="helper-text">
                <i class="fa-solid fa-circle-info"></i>
                Hover effect highlights the active row
            </div>
        </div>
    </div>
</div>

<!-- Card Switches -->
<div class="sw-section-title">
    <i class="fa-solid fa-id-card"></i>
    Card Switches
</div>

<div class="switch-grid three-cols">
    <div class="content-card">
        <div class="card-header">
            <div class="card-header-left">
                <div class="card-icon bg-info">
                    <i class="fa-solid fa-cloud"></i>
                </div>
                <div>
                    <h3>Cloud Sync</h3>
                    <p class="card-subtitle">Cloud storage toggle</p>
                </div>
            </div>
        </div>
        <div class="card-body">
            <div class="card-switch active">
                <div class="card-switch-icon">☁️</div>
                <div class="card-switch-title">Cloud Sync</div>
                <div class="card-switch-desc">Automatically sync to cloud</div>
            </div>

            <div class="divider"></div>

            <div class="status-indicator status-on" style="width: 100%; justify-content: center;">
                <i class="fa-solid fa-check"></i>
                Enabled
            </div>
        </div>
    </div>

    <div class="content-card">
        <div class="card-header">
            <div class="card-header-left">
                <div class="card-icon bg-success">
                    <i class="fa-solid fa-shield-halved"></i>
                </div>
                <div>
                    <h3>Security Mode</h3>
                    <p class="card-subtitle">Enhanced security toggle</p>
                </div>
            </div>
        </div>
        <div class="card-body">
            <div class="card-switch active">
                <div class="card-switch-icon">🛡️</div>
                <div class="card-switch-title">Security Mode</div>
                <div class="card-switch-desc">Enhanced protection active</div>
            </div>

            <div class="divider"></div>

            <div class="status-indicator status-on" style="width: 100%; justify-content: center;">
                <i class="fa-solid fa-check"></i>
                Enabled
            </div>
        </div>
    </div>

    <div class="content-card">
        <div class="card-header">
            <div class="card-header-left">
                <div class="card-icon bg-warning">
                    <i class="fa-solid fa-bell"></i>
                </div>
                <div>
                    <h3>Do Not Disturb</h3>
                    <p class="card-subtitle">Silence notifications</p>
                </div>
            </div>
        </div>
        <div class="card-body">
            <div class="card-switch">
                <div class="card-switch-icon">🔕</div>
                <div class="card-switch-title">Do Not Disturb</div>
                <div class="card-switch-desc">Silence all notifications</div>
            </div>

            <div class="divider"></div>

            <div class="status-indicator status-off" style="width: 100%; justify-content: center;">
                <i class="fa-solid fa-xmark"></i>
                Disabled
            </div>
        </div>
    </div>
</div>

<!-- Disabled & States -->
<div class="sw-section-title">
    <i class="fa-solid fa-circle-check"></i>
    Disabled & States
</div>

<div class="switch-grid">
    <div class="content-card">
        <div class="card-header">
            <div class="card-header-left">
                <div class="card-icon bg-warning">
                    <i class="fa-solid fa-ban"></i>
                </div>
                <div>
                    <h3>Disabled Switches</h3>
                    <p class="card-subtitle">Non-interactive toggles</p>
                </div>
            </div>
        </div>
        <div class="card-body">
            <div class="sw-example">
                <label class="sw-label">Disabled (ON)</label>
                <div class="toggle-switch">
                    <input type="checkbox" id="disabledOn" checked disabled>
                    <label for="disabledOn" class="toggle-switch-label">Locked ON</label>
                </div>
                <div class="sw-helper">
                    <i class="fa-solid fa-circle-info"></i>
                    Cannot be turned off
                </div>
            </div>

            <div class="sw-example">
                <label class="sw-label">Disabled (OFF)</label>
                <div class="toggle-switch">
                    <input type="checkbox" id="disabledOff" disabled>
                    <label for="disabledOff" class="toggle-switch-label">Locked OFF</label>
                </div>
                <div class="sw-helper">
                    <i class="fa-solid fa-circle-info"></i>
                    Cannot be turned on
                </div>
            </div>

            <div class="divider"></div>

            <div class="helper-text">
                <i class="fa-solid fa-circle-info"></i>
                Disabled switches have 50% opacity and no pointer events
            </div>
        </div>
    </div>

    <div class="content-card">
        <div class="card-header">
            <div class="card-header-left">
                <div class="card-icon bg-info">
                    <i class="fa-solid fa-circle-info"></i>
                </div>
                <div>
                    <h3>Switch States</h3>
                    <p class="card-subtitle">Different visual states</p>
                </div>
            </div>
        </div>
        <div class="card-body">
            <div class="sw-example">
                <label class="sw-label">Default State</label>
                <div class="toggle-switch">
                    <input type="checkbox" id="state1">
                    <label for="state1" class="toggle-switch-label">Default OFF</label>
                </div>
            </div>

            <div class="sw-example">
                <label class="sw-label">Checked State</label>
                <div class="toggle-switch">
                    <input type="checkbox" id="state2" checked>
                    <label for="state2" class="toggle-switch-label">Checked ON</label>
                </div>
            </div>

            <div class="sw-example">
                <label class="sw-label">Focus State</label>
                <div class="toggle-switch">
                    <input type="checkbox" id="state3">
                    <label for="state3" class="toggle-switch-label">Focus (tab to see)</label>
                </div>
                <div class="sw-helper">
                    <i class="fa-solid fa-circle-info"></i>
                    Tab to see focus ring
                </div>
            </div>

            <div class="divider"></div>

            <div class="helper-text">
                <i class="fa-solid fa-circle-info"></i>
                Focus state shows accent-colored ring
            </div>
        </div>
    </div>
</div>

<!-- Settings Panel Example -->
<div class="sw-section-title">
    <i class="fa-solid fa-clipboard-list"></i>
    Settings Panel Example
</div>

<div class="switch-grid full-width">
    <div class="content-card">
        <div class="card-header">
            <div class="card-header-left">
                <div class="card-icon bg-primary">
                    <i class="fa-solid fa-gear"></i>
                </div>
                <div>
                    <h3>Application Settings</h3>
                    <p class="card-subtitle">Complete settings panel with switches</p>
                </div>
            </div>
        </div>
        <div class="card-body">
            <h4 style="font-size: 14px; margin-bottom: 16px; color: var(--text-secondary);">🎨 Appearance</h4>
            <div style="display: flex; flex-direction: column; gap: 12px; margin-bottom: 24px;">
                <div class="list-switch">
                    <div class="list-switch-content">
                        <div class="list-switch-title">🌙 Dark Mode</div>
                        <div class="list-switch-desc">Use dark theme</div>
                    </div>
                    <div class="toggle-switch">
                        <input type="checkbox" id="settings1" checked>
                    </div>
                </div>

                <div class="list-switch">
                    <div class="list-switch-content">
                        <div class="list-switch-title">✨ Animations</div>
                        <div class="list-switch-desc">Enable UI animations</div>
                    </div>
                    <div class="toggle-switch">
                        <input type="checkbox" id="settings2" checked>
                    </div>
                </div>

                <div class="list-switch">
                    <div class="list-switch-content">
                        <div class="list-switch-title">🔤 Compact Mode</div>
                        <div class="list-switch-desc">Reduce spacing and padding</div>
                    </div>
                    <div class="toggle-switch">
                        <input type="checkbox" id="settings3">
                    </div>
                </div>
            </div>

            <h4 style="font-size: 14px; margin-bottom: 16px; color: var(--text-secondary);">🔔 Notifications</h4>
            <div style="display: flex; flex-direction: column; gap: 12px; margin-bottom: 24px;">
                <div class="list-switch">
                    <div class="list-switch-content">
                        <div class="list-switch-title">📱 Push Notifications</div>
                        <div class="list-switch-desc">Receive push notifications</div>
                    </div>
                    <div class="toggle-switch">
                        <input type="checkbox" id="settings4" checked>
                    </div>
                </div>

                <div class="list-switch">
                    <div class="list-switch-content">
                        <div class="list-switch-title">📧 Email Notifications</div>
                        <div class="list-switch-desc">Get email alerts</div>
                    </div>
                    <div class="toggle-switch">
                        <input type="checkbox" id="settings5">
                    </div>
                </div>

                <div class="list-switch">
                    <div class="list-switch-content">
                        <div class="list-switch-title">🔊 Sound Effects</div>
                        <div class="list-switch-desc">Play notification sounds</div>
                    </div>
                    <div class="toggle-switch">
                        <input type="checkbox" id="settings6" checked>
                    </div>
                </div>
            </div>

            <h4 style="font-size: 14px; margin-bottom: 16px; color: var(--text-secondary);">🔒 Privacy & Security</h4>
            <div style="display: flex; flex-direction: column; gap: 12px;">
                <div class="list-switch">
                    <div class="list-switch-content">
                        <div class="list-switch-title">🔐 Two-Factor Auth</div>
                        <div class="list-switch-desc">Enhanced login security</div>
                    </div>
                    <div class="toggle-switch">
                        <input type="checkbox" id="settings7" checked>
                    </div>
                </div>

                <div class="list-switch">
                    <div class="list-switch-content">
                        <div class="list-switch-title">👤 Profile Visibility</div>
                        <div class="list-switch-desc">Show profile to public</div>
                    </div>
                    <div class="toggle-switch">
                        <input type="checkbox" id="settings8">
                    </div>
                </div>

                <div class="list-switch">
                    <div class="list-switch-content">
                        <div class="list-switch-title">📊 Analytics Sharing</div>
                        <div class="list-switch-desc">Share usage analytics</div>
                    </div>
                    <div class="toggle-switch">
                        <input type="checkbox" id="settings9">
                    </div>
                </div>
            </div>

            <div class="divider"></div>

            <div style="display: flex; gap: 12px; margin-top: 16px;">
                <button class="btn btn-primary" style="flex: 1;">
                    <i class="fa-solid fa-floppy-disk"></i>
                    Save Settings
                </button>
                <button class="btn btn-secondary">
                    <i class="fa-solid fa-rotate-left"></i>
                    Reset Defaults
                </button>
            </div>
        </div>
    </div>
</div>

<!-- Implementation Guide -->
<div class="switch-grid full-width">
    <div class="content-card">
        <div class="card-header">
            <div class="card-header-left">
                <div class="card-icon bg-success">
                    <i class="fa-solid fa-code"></i>
                </div>
                <div>
                    <h3>Implementation Guide</h3>
                    <p class="card-subtitle">How to implement toggle switches with pure CSS</p>
                </div>
            </div>
        </div>
        <div class="card-body">
            <div class="code-block">
                <div style="color: var(--text-tertiary); margin-bottom: 8px;">1. HTML Structure:</div>
                <code style="color: var(--accent);">
                    &lt;div class="toggle-switch"&gt;<br>
                    &nbsp;&nbsp;&lt;input type="checkbox" id="mySwitch"&gt;<br>
                    &nbsp;&nbsp;&lt;label for="mySwitch" class="toggle-switch-label"&gt;Enable Feature&lt;/label&gt;<br>
                    &lt;/div&gt;
                </code>
            </div>

            <div class="code-block">
                <div style="color: var(--text-tertiary); margin-bottom: 8px; margin-top: 16px;">2. CSS Base Styles:</div>
                <code style="color: var(--success);">
                    .toggle-switch input[type="checkbox"] {<br>
                    &nbsp;&nbsp;appearance: none;<br>
                    &nbsp;&nbsp;width: 52px;<br>
                    &nbsp;&nbsp;height: 28px;<br>
                    &nbsp;&nbsp;background: var(--border);<br>
                    &nbsp;&nbsp;border-radius: 14px;<br>
                    &nbsp;&nbsp;position: relative;<br>
                    &nbsp;&nbsp;cursor: pointer;<br>
                    &nbsp;&nbsp;transition: all 0.3s ease;<br>
                    }<br>
                    <br>
                    .toggle-switch input[type="checkbox"]::after {<br>
                    &nbsp;&nbsp;content: '';<br>
                    &nbsp;&nbsp;position: absolute;<br>
                    &nbsp;&nbsp;top: 3px;<br>
                    &nbsp;&nbsp;left: 3px;<br>
                    &nbsp;&nbsp;width: 22px;<br>
                    &nbsp;&nbsp;height: 22px;<br>
                    &nbsp;&nbsp;background: white;<br>
                    &nbsp;&nbsp;border-radius: 50%;<br>
                    &nbsp;&nbsp;transition: all 0.3s ease;<br>
                    }
                </code>
            </div>

            <div class="code-block">
                <div style="color: var(--text-tertiary); margin-bottom: 8px; margin-top: 16px;">3. Checked State:</div>
                <code style="color: var(--warning);">
                    .toggle-switch input[type="checkbox"]:checked {<br>
                    &nbsp;&nbsp;background: var(--accent);<br>
                    }<br>
                    <br>
                    .toggle-switch input[type="checkbox"]:checked::after {<br>
                    &nbsp;&nbsp;left: 27px;<br>
                    }
                </code>
            </div>

            <div class="code-block">
                <div style="color: var(--text-tertiary); margin-bottom: 8px; margin-top: 16px;">4. JavaScript Event Listener:</div>
                <code style="color: var(--info);">
                    document.getElementById('mySwitch').addEventListener('change', function() {<br>
                    &nbsp;&nbsp;if (this.checked) {<br>
                    &nbsp;&nbsp;&nbsp;&nbsp;console.log('Switch is ON');<br>
                    &nbsp;&nbsp;} else {<br>
                    &nbsp;&nbsp;&nbsp;&nbsp;console.log('Switch is OFF');<br>
                    &nbsp;&nbsp;}<br>
                    });
                </code>
            </div>

            <div class="divider"></div>

            <div class="form-group">
                <h4 style="font-size: 14px; margin-bottom: 12px;">Features:</h4>
                <ul class="feature-list">
                    <li>
                        <i class="fa-solid fa-circle-check"></i>
                        <span><strong>Pure CSS</strong> - No JavaScript required</span>
                    </li>
                    <li>
                        <i class="fa-solid fa-circle-check"></i>
                        <span><strong>Smooth Animation</strong> - CSS transitions</span>
                    </li>
                    <li>
                        <i class="fa-solid fa-circle-check"></i>
                        <span><strong>Accessible</strong> - Keyboard navigation</span>
                    </li>
                    <li>
                        <i class="fa-solid fa-circle-check"></i>
                        <span><strong>Customizable</strong> - Easy to theme</span>
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
    // Update status indicators when switches change
    $('.toggle-switch input[type="checkbox"]').on('change', function() {
        const $switch = $(this).closest('.toggle-switch');
        const $example = $switch.closest('.sw-example');
        const $statusIndicator = $example.find('.status-indicator');
        const $helper = $example.find('.sw-helper');
        
        if ($(this).is(':checked')) {
            // Update status indicator
            if ($statusIndicator.length) {
                $statusIndicator.removeClass('status-off').addClass('status-on');
                $statusIndicator.html('<i class="fa-solid fa-check"></i> ON');
            }
            
            // Update helper text
            if ($helper.length) {
                $helper.html('<i class="fa-solid fa-circle-info"></i> Currently: ON');
            }
        } else {
            // Update status indicator
            if ($statusIndicator.length) {
                $statusIndicator.removeClass('status-on').addClass('status-off');
                $statusIndicator.html('<i class="fa-solid fa-xmark"></i> OFF');
            }
            
            // Update helper text
            if ($helper.length) {
                $helper.html('<i class="fa-solid fa-circle-info"></i> Currently: OFF');
            }
        }
    });
    
    // Card switch click to toggle
    $('.card-switch').on('click', function() {
        $(this).toggleClass('active');
        
        const $card = $(this);
        const $statusIndicator = $card.closest('.card-body').find('.status-indicator');
        
        if ($card.hasClass('active')) {
            $statusIndicator.removeClass('status-off').addClass('status-on');
            $statusIndicator.html('<i class="fa-solid fa-check"></i> Enabled');
        } else {
            $statusIndicator.removeClass('status-on').addClass('status-off');
            $statusIndicator.html('<i class="fa-solid fa-xmark"></i> Disabled');
        }
    });
    
    // Settings panel save button
    $('.btn-primary:contains("Save Settings")').on('click', function(e) {
        e.preventDefault();
        
        const switches = [];
        $('.toggle-switch input[type="checkbox"]').each(function() {
            const id = $(this).attr('id');
            const checked = $(this).is(':checked');
            switches.push({ id, checked });
        });
        
        // Show success message
        if (typeof MetroAdmin !== 'undefined' && MetroAdmin.showToast) {
            MetroAdmin.showToast('Settings saved successfully!', 'success');
        } else {
            Swal.fire({
                icon: 'success',
                title: 'Saved!',
                text: 'Settings have been saved successfully.',
                timer: 2000,
                showConfirmButton: false
            });
        }
        
        console.log('Settings:', switches);
    });
    
    // Reset button
    $('.btn-secondary:contains("Reset Defaults")').on('click', function(e) {
        e.preventDefault();
        
        Swal.fire({
            title: 'Reset Settings?',
            text: 'All settings will be reset to default values.',
            icon: 'warning',
            showCancelButton: true,
            confirmButtonColor: 'var(--accent)',
            cancelButtonColor: 'var(--danger)',
            confirmButtonText: 'Yes, reset it!',
            cancelButtonText: 'Cancel'
        }).then((result) => {
            if (result.isConfirmed) {
                // Reset all switches to default
                $('.toggle-switch input[type="checkbox"]').prop('checked', false);
                
                // Set default checked switches
                $('#settings1, #settings2, #settings4, #settings6, #settings7').prop('checked', true);
                
                // Trigger change event
                $('.toggle-switch input[type="checkbox"]').trigger('change');
                
                Swal.fire({
                    icon: 'success',
                    title: 'Reset!',
                    text: 'Settings have been reset to defaults.',
                    timer: 2000,
                    showConfirmButton: false
                });
            }
        });
    });
    
    // Log switch changes for debugging
    $('.toggle-switch input[type="checkbox"]').on('change', function() {
        const id = $(this).attr('id');
        const state = $(this).is(':checked') ? 'ON' : 'OFF';
        console.log(`Switch ${id}: ${state}`);
    });
});
</script>
@endpush
