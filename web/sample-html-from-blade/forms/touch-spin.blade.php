@extends('layouts.app')

@section('title', 'Touch Spin')

@push('styles')
@endpush

@section('content')
<div class="page-header">
    <div>
        <h1>Touch Spin</h1>
        <p>Number input dengan tombol increment/decrement untuk pemilihan angka interaktif</p>
    </div>
    <div style="display: flex; gap: 10px;">
        <button class="btn btn-secondary">
            <i class="fa-solid fa-book"></i>
            <span>Documentation</span>
        </button>
        <button class="btn btn-primary">
            <i class="fa-solid fa-arrow-up-9-1"></i>
            <span>Try Spin</span>
        </button>
    </div>
</div>

<!-- Info Alert -->
<div class="content-card" style="margin-bottom: 24px; border-left: 4px solid var(--accent);">
    <div class="card-body" style="padding: 16px 20px;">
        <div style="display: flex; gap: 12px; align-items: start;">
            <i class="fa-solid fa-circle-info" style="color: var(--accent); font-size: 20px; margin-top: 2px;"></i>
            <div style="flex: 1;">
                <h4 style="margin-bottom: 4px; font-size: 14px;">What is Touch Spin?</h4>
                <p style="font-size: 13px; color: var(--text-secondary); margin: 0;">Touch Spin adalah komponen number input yang menyediakan tombol + dan - untuk increment/decrement nilai. Sangat berguna untuk quantity selection, age input, dan numeric forms.</p>
            </div>
        </div>
    </div>
</div>

<!-- Basic Touch Spin -->
<div class="ts-section-title">
    <i class="fa-solid fa-arrow-up-9-1"></i>
    Basic Touch Spin <span class="badge badge-primary">Essential</span>
</div>

<div class="touchspin-grid">
    <div class="content-card">
        <div class="card-header">
            <div class="card-header-left">
                <div class="card-icon bg-primary">
                    <i class="fa-solid fa-hashtag"></i>
                </div>
                <div>
                    <h3>Default Spinners</h3>
                    <p class="card-subtitle">Basic increment/decrement</p>
                </div>
            </div>
        </div>
        <div class="card-body">
            <div class="ts-example">
                <label class="ts-label">
                    Quantity
                    <span class="ts-range">0 - 100</span>
                </label>
                <div class="touch-spin">
                    <button class="spin-btn">−</button>
                    <input type="number" class="spin-input" value="5" min="0" max="100">
                    <button class="spin-btn">+</button>
                </div>
                <div class="ts-helper">
                    <i class="fa-solid fa-circle-info"></i>
                    Step: 1
                </div>
            </div>

            <div class="ts-example">
                <label class="ts-label">
                    Age
                    <span class="ts-range">1 - 120</span>
                </label>
                <div class="touch-spin">
                    <button class="spin-btn">−</button>
                    <input type="number" class="spin-input" value="25" min="1" max="120">
                    <button class="spin-btn">+</button>
                </div>
                <div class="ts-helper">
                    <i class="fa-solid fa-circle-info"></i>
                    Step: 1 year
                </div>
            </div>

            <div class="ts-example">
                <label class="ts-label">
                    Rating
                    <span class="ts-range">1 - 5</span>
                </label>
                <div class="touch-spin">
                    <button class="spin-btn">−</button>
                    <input type="number" class="spin-input" value="4" min="1" max="5">
                    <button class="spin-btn">+</button>
                </div>
                <div class="ts-helper">
                    <i class="fa-solid fa-circle-info"></i>
                    Step: 1 star
                </div>
            </div>

            <div class="ts-example">
                <label class="ts-label">
                    Items Count
                    <span class="ts-range">0 - 1000</span>
                </label>
                <div class="touch-spin">
                    <button class="spin-btn">−</button>
                    <input type="number" class="spin-input" value="42" min="0" max="1000">
                    <button class="spin-btn">+</button>
                </div>
                <div class="step-indicator">Step: 1</div>
            </div>

            <div class="divider"></div>

            <div class="spin-preview">
                <div class="spin-preview-label">Current Quantity:</div>
                <div class="spin-preview-value">5 items</div>
            </div>
        </div>
    </div>

    <div class="content-card">
        <div class="card-header">
            <div class="card-header-left">
                <div class="card-icon bg-success">
                    <i class="fa-solid fa-sort-numeric-up"></i>
                </div>
                <div>
                    <h3>Decimal & Float</h3>
                    <p class="card-subtitle">Decimal number spinners</p>
                </div>
            </div>
        </div>
        <div class="card-body">
            <div class="ts-example">
                <label class="ts-label">
                    Price (USD)
                    <span class="ts-range">0.00 - 999.99</span>
                </label>
                <div class="touch-spin">
                    <button class="spin-btn">−</button>
                    <input type="number" class="spin-input" value="19.99" min="0" max="999.99" step="0.01">
                    <button class="spin-btn">+</button>
                </div>
                <div class="step-indicator">Step: 0.01</div>
            </div>

            <div class="ts-example">
                <label class="ts-label">
                    Weight (kg)
                    <span class="ts-range">0.0 - 200.0</span>
                </label>
                <div class="touch-spin">
                    <button class="spin-btn">−</button>
                    <input type="number" class="spin-input" value="65.5" min="0" max="200" step="0.1">
                    <button class="spin-btn">+</button>
                </div>
                <div class="step-indicator">Step: 0.1</div>
            </div>

            <div class="ts-example">
                <label class="ts-label">
                    Percentage
                    <span class="ts-range">0.0 - 100.0</span>
                </label>
                <div class="touch-spin">
                    <button class="spin-btn">−</button>
                    <input type="number" class="spin-input" value="75.5" min="0" max="100" step="0.5">
                    <button class="spin-btn">+</button>
                </div>
                <div class="step-indicator">Step: 0.5</div>
            </div>

            <div class="ts-example">
                <label class="ts-label">
                    Temperature (°C)
                    <span class="ts-range">-50.0 - 100.0</span>
                </label>
                <div class="touch-spin">
                    <button class="spin-btn">−</button>
                    <input type="number" class="spin-input" value="25.0" min="-50" max="100" step="0.5">
                    <button class="spin-btn">+</button>
                </div>
                <div class="step-indicator">Step: 0.5°C</div>
            </div>

            <div class="divider"></div>

            <div class="form-group">
                <h4 style="font-size: 14px; margin-bottom: 12px;">Features:</h4>
                <ul class="feature-list">
                    <li>
                        <i class="fa-solid fa-circle-check"></i>
                        <span>Integer & decimal support</span>
                    </li>
                    <li>
                        <i class="fa-solid fa-circle-check"></i>
                        <span>Customizable step size</span>
                    </li>
                    <li>
                        <i class="fa-solid fa-circle-check"></i>
                        <span>Min/max value limits</span>
                    </li>
                    <li>
                        <i class="fa-solid fa-circle-check"></i>
                        <span>Keyboard navigation</span>
                    </li>
                </ul>
            </div>
        </div>
    </div>
</div>

<!-- Size Variants -->
<div class="ts-section-title">
    <i class="fa-solid fa-up-down"></i>
    Size Variants
</div>

<div class="touchspin-grid three-cols">
    <div class="content-card">
        <div class="card-header">
            <div class="card-header-left">
                <div class="card-icon bg-info">
                    <i class="fa-solid fa-compress"></i>
                </div>
                <div>
                    <h3>Small Size</h3>
                    <p class="card-subtitle">Compact spinner</p>
                </div>
            </div>
        </div>
        <div class="card-body">
            <div class="ts-example">
                <label class="ts-label">Small Quantity</label>
                <div class="touch-spin touch-spin-sm">
                    <button class="spin-btn">−</button>
                    <input type="number" class="spin-input" value="3" min="0" max="50">
                    <button class="spin-btn">+</button>
                </div>
                <div class="ts-helper">
                    <i class="fa-solid fa-circle-info"></i>
                    Size: Small (32px buttons)
                </div>
            </div>

            <div class="ts-example">
                <label class="ts-label">Small Decimal</label>
                <div class="touch-spin touch-spin-sm">
                    <button class="spin-btn">−</button>
                    <input type="number" class="spin-input" value="9.99" min="0" max="100" step="0.01">
                    <button class="spin-btn">+</button>
                </div>
                <div class="step-indicator">Step: 0.01</div>
            </div>

            <div class="divider"></div>

            <div class="helper-text">
                <i class="fa-solid fa-circle-info"></i>
                Best for: Dense forms, tables, mobile
            </div>
        </div>
    </div>

    <div class="content-card">
        <div class="card-header">
            <div class="card-header-left">
                <div class="card-icon bg-primary">
                    <i class="fa-solid fa-expand"></i>
                </div>
                <div>
                    <h3>Medium Size</h3>
                    <p class="card-subtitle">Default spinner (Recommended)</p>
                </div>
            </div>
        </div>
        <div class="card-body">
            <div class="ts-example">
                <label class="ts-label">Medium Quantity</label>
                <div class="touch-spin">
                    <button class="spin-btn">−</button>
                    <input type="number" class="spin-input" value="10" min="0" max="100">
                    <button class="spin-btn">+</button>
                </div>
                <div class="ts-helper">
                    <i class="fa-solid fa-circle-info"></i>
                    Size: Medium (40px buttons)
                </div>
            </div>

            <div class="ts-example">
                <label class="ts-label">Medium Decimal</label>
                <div class="touch-spin">
                    <button class="spin-btn">−</button>
                    <input type="number" class="spin-input" value="29.99" min="0" max="500" step="0.01">
                    <button class="spin-btn">+</button>
                </div>
                <div class="step-indicator">Step: 0.01</div>
            </div>

            <div class="divider"></div>

            <div class="helper-text">
                <i class="fa-solid fa-circle-info"></i>
                Best for: Standard forms, most use cases
            </div>
        </div>
    </div>

    <div class="content-card">
        <div class="card-header">
            <div class="card-header-left">
                <div class="card-icon bg-warning">
                    <i class="fa-solid fa-maximize"></i>
                </div>
                <div>
                    <h3>Large Size</h3>
                    <p class="card-subtitle">Large spinner for visibility</p>
                </div>
            </div>
        </div>
        <div class="card-body">
            <div class="ts-example">
                <label class="ts-label">Large Quantity</label>
                <div class="touch-spin touch-spin-lg">
                    <button class="spin-btn">−</button>
                    <input type="number" class="spin-input" value="20" min="0" max="200">
                    <button class="spin-btn">+</button>
                </div>
                <div class="ts-helper">
                    <i class="fa-solid fa-circle-info"></i>
                    Size: Large (48px buttons)
                </div>
            </div>

            <div class="ts-example">
                <label class="ts-label">Large Decimal</label>
                <div class="touch-spin touch-spin-lg">
                    <button class="spin-btn">−</button>
                    <input type="number" class="spin-input" value="49.99" min="0" max="1000" step="0.01">
                    <button class="spin-btn">+</button>
                </div>
                <div class="step-indicator">Step: 0.01</div>
            </div>

            <div class="divider"></div>

            <div class="helper-text">
                <i class="fa-solid fa-circle-info"></i>
                Best for: Touchscreens, accessibility
            </div>
        </div>
    </div>
</div>

<!-- Custom Styles -->
<div class="ts-section-title">
    <i class="fa-solid fa-palette"></i>
    Custom Styles & Colors
</div>

<div class="touchspin-grid">
    <div class="content-card">
        <div class="card-header">
            <div class="card-header-left">
                <div class="card-icon bg-primary">
                    <i class="fa-solid fa-paintbrush"></i>
                </div>
                <div>
                    <h3>Colored Buttons</h3>
                    <p class="card-subtitle">Accent-colored spinners</p>
                </div>
            </div>
        </div>
        <div class="card-body">
            <div class="ts-example">
                <label class="ts-label">Primary Theme</label>
                <div class="touch-spin">
                    <button class="spin-btn spin-btn-primary">−</button>
                    <input type="number" class="spin-input" value="10" min="0" max="100">
                    <button class="spin-btn spin-btn-primary">+</button>
                </div>
            </div>

            <div class="ts-example">
                <label class="ts-label">Success Theme</label>
                <div class="touch-spin">
                    <button class="spin-btn spin-btn-success">−</button>
                    <input type="number" class="spin-input" value="25" min="0" max="50">
                    <button class="spin-btn spin-btn-success">+</button>
                </div>
            </div>

            <div class="ts-example">
                <label class="ts-label">Danger Theme</label>
                <div class="touch-spin">
                    <button class="spin-btn spin-btn-danger">−</button>
                    <input type="number" class="spin-input" value="5" min="0" max="10">
                    <button class="spin-btn spin-btn-danger">+</button>
                </div>
            </div>

            <div class="ts-example">
                <label class="ts-label">Warning Theme</label>
                <div class="touch-spin">
                    <button class="spin-btn spin-btn-warning">−</button>
                    <input type="number" class="spin-input" value="15" min="0" max="30">
                    <button class="spin-btn spin-btn-warning">+</button>
                </div>
            </div>

            <div class="divider"></div>

            <div class="helper-text">
                <i class="fa-solid fa-circle-info"></i>
                Use color variants for context: success (add), danger (remove)
            </div>
        </div>
    </div>

    <div class="content-card">
        <div class="card-header">
            <div class="card-header-left">
                <div class="card-icon bg-success">
                    <i class="fa-solid fa-border-style"></i>
                </div>
                <div>
                    <h3>Custom Border Style</h3>
                    <p class="card-subtitle">Accent border spinners</p>
                </div>
            </div>
        </div>
        <div class="card-body">
            <div class="ts-example">
                <label class="ts-label">Custom Accent Border</label>
                <div class="touch-spin touch-spin-custom">
                    <button class="spin-btn">−</button>
                    <input type="number" class="spin-input" value="8" min="0" max="20">
                    <button class="spin-btn">+</button>
                </div>
            </div>

            <div class="ts-example">
                <label class="ts-label">Custom with Decimal</label>
                <div class="touch-spin touch-spin-custom">
                    <button class="spin-btn">−</button>
                    <input type="number" class="spin-input" value="15.5" min="0" max="100" step="0.5">
                    <button class="spin-btn">+</button>
                </div>
            </div>

            <div class="divider"></div>

            <div class="form-group">
                <h4 style="font-size: 14px; margin-bottom: 12px;">Customization:</h4>
                <ul class="feature-list">
                    <li>
                        <i class="fa-solid fa-circle-check"></i>
                        <span>Custom border color</span>
                    </li>
                    <li>
                        <i class="fa-solid fa-circle-check"></i>
                        <span>Hover accent background</span>
                    </li>
                    <li>
                        <i class="fa-solid fa-circle-check"></i>
                        <span>Bold button borders</span>
                    </li>
                    <li>
                        <i class="fa-solid fa-circle-check"></i>
                        <span>Accent text on hover</span>
                    </li>
                </ul>
            </div>
        </div>
    </div>
</div>

<!-- Disabled & Validation -->
<div class="ts-section-title">
    <i class="fa-solid fa-circle-check"></i>
    States & Validation
</div>

<div class="touchspin-grid">
    <div class="content-card">
        <div class="card-header">
            <div class="card-header-left">
                <div class="card-icon bg-warning">
                    <i class="fa-solid fa-ban"></i>
                </div>
                <div>
                    <h3>Disabled State</h3>
                    <p class="card-subtitle">Non-interactive spinners</p>
                </div>
            </div>
        </div>
        <div class="card-body">
            <div class="ts-example">
                <label class="ts-label">Disabled Quantity</label>
                <div class="touch-spin disabled">
                    <button class="spin-btn" disabled>−</button>
                    <input type="number" class="spin-input" value="5" disabled>
                    <button class="spin-btn" disabled>+</button>
                </div>
                <div class="ts-helper">
                    <i class="fa-solid fa-circle-info"></i>
                    Cannot be modified
                </div>
            </div>

            <div class="ts-example">
                <label class="ts-label">Disabled Decimal</label>
                <div class="touch-spin disabled">
                    <button class="spin-btn" disabled>−</button>
                    <input type="number" class="spin-input" value="19.99" disabled>
                    <button class="spin-btn" disabled>+</button>
                </div>
                <div class="ts-helper">
                    <i class="fa-solid fa-circle-info"></i>
                    Read-only value
                </div>
            </div>

            <div class="divider"></div>

            <div class="helper-text">
                <i class="fa-solid fa-circle-info"></i>
                Disabled spinners have 50% opacity and no pointer events
            </div>
        </div>
    </div>

    <div class="content-card">
        <div class="card-header">
            <div class="card-header-left">
                <div class="card-icon bg-danger">
                    <i class="fa-solid fa-triangle-exclamation"></i>
                </div>
                <div>
                    <h3>Boundary Conditions</h3>
                    <p class="card-subtitle">Min/max limit reached</p>
                </div>
            </div>
        </div>
        <div class="card-body">
            <div class="ts-example">
                <label class="ts-label">At Minimum (0)</label>
                <div class="touch-spin">
                    <button class="spin-btn">−</button>
                    <input type="number" class="spin-input" value="0" min="0" max="100">
                    <button class="spin-btn">+</button>
                </div>
                <div class="ts-helper" style="color: var(--warning);">
                    <i class="fa-solid fa-circle-exclamation"></i>
                    Cannot decrement further
                </div>
            </div>

            <div class="ts-example">
                <label class="ts-label">At Maximum (100)</label>
                <div class="touch-spin">
                    <button class="spin-btn">−</button>
                    <input type="number" class="spin-input" value="100" min="0" max="100">
                    <button class="spin-btn">+</button>
                </div>
                <div class="ts-helper" style="color: var(--warning);">
                    <i class="fa-solid fa-circle-exclamation"></i>
                    Cannot increment further
                </div>
            </div>

            <div class="divider"></div>

            <div class="helper-text">
                <i class="fa-solid fa-circle-info"></i>
                Boundary checks prevent values outside min/max range
            </div>
        </div>
    </div>
</div>

<!-- E-commerce Cart Example -->
<div class="ts-section-title">
    <i class="fa-solid fa-clipboard-list"></i>
    E-commerce Cart Example
</div>

<div class="touchspin-grid full-width">
    <div class="content-card">
        <div class="card-header">
            <div class="card-header-left">
                <div class="card-icon bg-primary">
                    <i class="fa-solid fa-cart-shopping"></i>
                </div>
                <div>
                    <h3>Shopping Cart</h3>
                    <p class="card-subtitle">Product quantity management</p>
                </div>
            </div>
        </div>
        <div class="card-body">
            <div style="display: flex; flex-direction: column; gap: 16px;">
                <div style="display: flex; align-items: center; gap: 16px; padding: 16px; border: 1px solid var(--border); border-radius: var(--radius-md);">
                    <div style="font-size: 32px; flex: 0 0 60px; text-align: center;">👕</div>
                    <div style="flex: 1;">
                        <div style="font-weight: 600; margin-bottom: 4px;">T-Shirt Basic</div>
                        <div style="font-size: 12px; color: var(--text-tertiary);">Size: M | Color: Blue</div>
                        <div style="font-weight: 600; color: var(--accent); margin-top: 4px;">$29.99</div>
                    </div>
                    <div>
                        <label class="ts-label">Quantity</label>
                        <div class="touch-spin">
                            <button class="spin-btn">−</button>
                            <input type="number" class="spin-input" value="2" min="1" max="10">
                            <button class="spin-btn">+</button>
                        </div>
                    </div>
                    <div style="font-weight: 700; font-size: 18px; color: var(--accent);">$59.98</div>
                </div>

                <div style="display: flex; align-items: center; gap: 16px; padding: 16px; border: 1px solid var(--border); border-radius: var(--radius-md);">
                    <div style="font-size: 32px; flex: 0 0 60px; text-align: center;">👖</div>
                    <div style="flex: 1;">
                        <div style="font-weight: 600; margin-bottom: 4px;">Jeans Slim Fit</div>
                        <div style="font-size: 12px; color: var(--text-tertiary);">Size: 32 | Color: Black</div>
                        <div style="font-weight: 600; color: var(--accent); margin-top: 4px;">$49.99</div>
                    </div>
                    <div>
                        <label class="ts-label">Quantity</label>
                        <div class="touch-spin">
                            <button class="spin-btn">−</button>
                            <input type="number" class="spin-input" value="1" min="1" max="5">
                            <button class="spin-btn">+</button>
                        </div>
                    </div>
                    <div style="font-weight: 700; font-size: 18px; color: var(--accent);">$49.99</div>
                </div>

                <div style="display: flex; align-items: center; gap: 16px; padding: 16px; border: 1px solid var(--border); border-radius: var(--radius-md);">
                    <div style="font-size: 32px; flex: 0 0 60px; text-align: center;">👟</div>
                    <div style="flex: 1;">
                        <div style="font-weight: 600; margin-bottom: 4px;">Sneakers Sport</div>
                        <div style="font-size: 12px; color: var(--text-tertiary);">Size: 42 | Color: White</div>
                        <div style="font-weight: 600; color: var(--accent); margin-top: 4px;">$89.99</div>
                    </div>
                    <div>
                        <label class="ts-label">Quantity</label>
                        <div class="touch-spin">
                            <button class="spin-btn">−</button>
                            <input type="number" class="spin-input" value="1" min="0" max="3">
                            <button class="spin-btn">+</button>
                        </div>
                    </div>
                    <div style="font-weight: 700; font-size: 18px; color: var(--accent);">$89.99</div>
                </div>
            </div>

            <div class="divider"></div>

            <div style="display: flex; justify-content: space-between; align-items: center; padding: 16px; background: var(--bg-secondary); border-radius: var(--radius-md);">
                <div>
                    <div style="font-size: 12px; color: var(--text-tertiary);">Total Items:</div>
                    <div style="font-size: 18px; font-weight: 700;">4 items</div>
                </div>
                <div style="text-align: right;">
                    <div style="font-size: 12px; color: var(--text-tertiary);">Total Price:</div>
                    <div style="font-size: 24px; font-weight: 700; color: var(--accent);">$199.96</div>
                </div>
            </div>

            <div style="display: flex; gap: 12px; margin-top: 16px;">
                <button class="btn btn-primary" style="flex: 1;">
                    <i class="fa-solid fa-credit-card"></i>
                    Checkout
                </button>
                <button class="btn btn-secondary">
                    <i class="fa-solid fa-rotate-left"></i>
                    Update Cart
                </button>
            </div>
        </div>
    </div>
</div>

<!-- Implementation Guide -->
<div class="touchspin-grid full-width">
    <div class="content-card">
        <div class="card-header">
            <div class="card-header-left">
                <div class="card-icon bg-success">
                    <i class="fa-solid fa-code"></i>
                </div>
                <div>
                    <h3>Implementation Guide</h3>
                    <p class="card-subtitle">How to implement touch spin with JavaScript</p>
                </div>
            </div>
        </div>
        <div class="card-body">
            <div class="code-block">
                <div style="color: var(--text-tertiary); margin-bottom: 8px;">1. HTML Structure:</div>
                <code style="color: var(--accent);">
                    &lt;div class="touch-spin"&gt;<br>
                    &nbsp;&nbsp;&lt;button class="spin-btn" onclick="decrement(this)"&gt;−&lt;/button&gt;<br>
                    &nbsp;&nbsp;&lt;input type="number" class="spin-input" value="5" min="0" max="100"&gt;<br>
                    &nbsp;&nbsp;&lt;button class="spin-btn" onclick="increment(this)"&gt;+&lt;/button&gt;<br>
                    &lt;/div&gt;
                </code>
            </div>

            <div class="code-block">
                <div style="color: var(--text-tertiary); margin-bottom: 8px; margin-top: 16px;">2. JavaScript Functions:</div>
                <code style="color: var(--success);">
                    function increment(btn) {<br>
                    &nbsp;&nbsp;const input = btn.parentElement.querySelector('.spin-input');<br>
                    &nbsp;&nbsp;const max = parseInt(input.max) || Infinity;<br>
                    &nbsp;&nbsp;const step = parseFloat(input.step) || 1;<br>
                    &nbsp;&nbsp;const newVal = parseFloat(input.value) + step;<br>
                    &nbsp;&nbsp;if (newVal <= max) input.value = newVal;<br>
                    }<br>
                    <br>
                    function decrement(btn) {<br>
                    &nbsp;&nbsp;const input = btn.parentElement.querySelector('.spin-input');<br>
                    &nbsp;&nbsp;const min = parseInt(input.min) || 0;<br>
                    &nbsp;&nbsp;const step = parseFloat(input.step) || 1;<br>
                    &nbsp;&nbsp;const newVal = parseFloat(input.value) - step;<br>
                    &nbsp;&nbsp;if (newVal >= min) input.value = newVal;<br>
                    }
                </code>
            </div>

            <div class="code-block">
                <div style="color: var(--text-tertiary); margin-bottom: 8px; margin-top: 16px;">3. With Bootstrap TouchSpin Library:</div>
                <code style="color: var(--warning);">
                    $('input.touchspin').TouchSpin({<br>
                    &nbsp;&nbsp;min: 0,<br>
                    &nbsp;&nbsp;max: 100,<br>
                    &nbsp;&nbsp;step: 1,<br>
                    &nbsp;&nbsp;decimals: 2,<br>
                    &nbsp;&nbsp;boostat: 5,<br>
                    &nbsp;&nbsp;maxboostedstep: 10<br>
                    });
                </code>
            </div>

            <div class="divider"></div>

            <div class="form-group">
                <h4 style="font-size: 14px; margin-bottom: 12px;">Popular Libraries:</h4>
                <ul class="feature-list">
                    <li>
                        <i class="fa-solid fa-circle-check"></i>
                        <span><strong>Bootstrap TouchSpin</strong> - Bootstrap integration</span>
                    </li>
                    <li>
                        <i class="fa-solid fa-circle-check"></i>
                        <span><strong>Custom Implementation</strong> - No dependencies</span>
                    </li>
                    <li>
                        <i class="fa-solid fa-circle-check"></i>
                        <span><strong>HTML5 Native</strong> - Built-in number input</span>
                    </li>
                    <li>
                        <i class="fa-solid fa-circle-check"></i>
                        <span><strong>React Number Input</strong> - React component</span>
                    </li>
                </ul>
            </div>
        </div>
    </div>
</div>
@endsection
