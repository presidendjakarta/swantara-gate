@extends('layouts.app')

@section('title', 'Input Groups')

@push('styles')
@endpush

@section('content')
<div class="page-header">
    <div>
        <h1>Input Groups</h1>
        <p>Kombinasikan input dengan text, icon, button, dan dropdown</p>
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
                <h4 style="margin-bottom: 4px; font-size: 14px;">Input Groups Overview</h4>
                <p style="font-size: 13px; color: var(--text-secondary); margin: 0;">Input groups allow you to prepend or append text, icons, buttons, or dropdowns to form inputs. Perfect for creating enhanced form controls with contextual information.</p>
            </div>
        </div>
    </div>
</div>

<!-- Basic Input Groups -->
<div class="ig-section-title">
    <i class="fa-solid fa-object-group"></i>
    Basic Input Groups <span class="badge badge-primary">Essential</span>
</div>

<div class="input-groups-grid">
    <div class="content-card">
        <div class="card-header">
            <div class="card-header-left">
                <div class="card-icon bg-primary">
                    <i class="fa-solid fa-circle-plus"></i>
                </div>
                <div>
                    <h3>Prepend Addons</h3>
                    <p class="card-subtitle">Text/icon before input</p>
                </div>
            </div>
        </div>
        <div class="card-body">
            <div class="ig-example">
                <label class="ig-label">Username</label>
                <div class="input-group">
                    <span class="input-group-text">@</span>
                    <input type="text" class="form-control" placeholder="Username">
                </div>
            </div>

            <div class="ig-example">
                <label class="ig-label">Website URL</label>
                <div class="input-group">
                    <span class="input-group-text">
                        <i class="fa-solid fa-globe"></i>
                    </span>
                    <input type="url" class="form-control" placeholder="https://example.com">
                </div>
            </div>

            <div class="ig-example">
                <label class="ig-label">Email Address</label>
                <div class="input-group">
                    <span class="input-group-text">
                        <i class="fa-solid fa-envelope"></i>
                    </span>
                    <input type="email" class="form-control" placeholder="name@example.com">
                </div>
            </div>

            <div class="ig-example">
                <label class="ig-label">Phone Number</label>
                <div class="input-group">
                    <span class="input-group-text">
                        <i class="fa-solid fa-phone"></i>
                    </span>
                    <input type="tel" class="form-control" placeholder="+62 812-3456-7890">
                </div>
            </div>
        </div>
    </div>

    <div class="content-card">
        <div class="card-header">
            <div class="card-header-left">
                <div class="card-icon bg-success">
                    <i class="fa-solid fa-circle-minus"></i>
                </div>
                <div>
                    <h3>Append Addons</h3>
                    <p class="card-subtitle">Text/icon after input</p>
                </div>
            </div>
        </div>
        <div class="card-body">
            <div class="ig-example">
                <label class="ig-label">Website Domain</label>
                <div class="input-group">
                    <input type="text" class="form-control" placeholder="yourdomain">
                    <span class="input-group-text">.com</span>
                </div>
            </div>

            <div class="ig-example">
                <label class="ig-label">Price (USD)</label>
                <div class="input-group">
                    <input type="number" class="form-control" placeholder="0.00">
                    <span class="input-group-text">USD</span>
                </div>
            </div>

            <div class="ig-example">
                <label class="ig-label">Weight (kg)</label>
                <div class="input-group">
                    <input type="number" class="form-control" placeholder="0">
                    <span class="input-group-text">kg</span>
                </div>
            </div>

            <div class="ig-example">
                <label class="ig-label">Search</label>
                <div class="input-group">
                    <input type="search" class="form-control" placeholder="Search...">
                    <span class="input-group-text">
                        <i class="fa-solid fa-magnifying-glass"></i>
                    </span>
                </div>
            </div>
        </div>
    </div>
</div>

<!-- Both Sides Input Groups -->
<div class="ig-section-title">
    <i class="fa-solid fa-left-right"></i>
    Both Sides Input Groups
</div>

<div class="input-groups-grid">
    <div class="content-card">
        <div class="card-header">
            <div class="card-header-left">
                <div class="card-icon bg-info">
                    <i class="fa-solid fa-arrows-left-right"></i>
                </div>
                <div>
                    <h3>Prepend & Append</h3>
                    <p class="card-subtitle">Addons on both sides</p>
                </div>
            </div>
        </div>
        <div class="card-body">
            <div class="ig-example">
                <label class="ig-label">Full URL</label>
                <div class="input-group">
                    <span class="input-group-text">https://</span>
                    <input type="text" class="form-control" placeholder="www.example.com">
                    <span class="input-group-text">.com</span>
                </div>
            </div>

            <div class="ig-example">
                <label class="ig-label">Price Range</label>
                <div class="input-group">
                    <span class="input-group-text">$
                    </span>
                    <input type="number" class="form-control" placeholder="0">
                    <span class="input-group-text">.00</span>
                </div>
            </div>

            <div class="ig-example">
                <label class="ig-label">Email with Domain</label>
                <div class="input-group">
                    <span class="input-group-text">
                        <i class="fa-solid fa-envelope"></i>
                    </span>
                    <input type="email" class="form-control" placeholder="username">
                    <span class="input-group-text">@gmail.com</span>
                </div>
            </div>

            <div class="ig-example">
                <label class="ig-label">Phone with Country Code</label>
                <div class="input-group">
                    <span class="input-group-text">
                        <i class="fa-solid fa-phone"></i>
                    </span>
                    <input type="tel" class="form-control" placeholder="812-3456-7890">
                    <span class="input-group-text">+62</span>
                </div>
            </div>
        </div>
    </div>

    <div class="content-card">
        <div class="card-header">
            <div class="card-header-left">
                <div class="card-icon bg-warning">
                    <i class="fa-solid fa-coins"></i>
                </div>
                <div>
                    <h3>Multiple Addons</h3>
                    <p class="card-subtitle">Chain multiple elements</p>
                </div>
            </div>
        </div>
        <div class="card-body">
            <div class="ig-example">
                <label class="ig-label">Name</label>
                <div class="input-group">
                    <span class="input-group-text">
                        <i class="fa-solid fa-user"></i>
                    </span>
                    <input type="text" class="form-control" placeholder="First name">
                    <input type="text" class="form-control" placeholder="Last name">
                </div>
            </div>

            <div class="ig-example">
                <label class="ig-label">Amount with Currency</label>
                <div class="input-group">
                    <span class="input-group-text">$
                    </span>
                    <input type="number" class="form-control" placeholder="0.00">
                    <span class="input-group-text">.00</span>
                    <span class="input-group-text">USD</span>
                </div>
            </div>

            <div class="divider"></div>

            <div class="form-group">
                <h4 style="font-size: 14px; margin-bottom: 12px;">Features:</h4>
                <ul class="feature-list">
                    <li>
                        <i class="fa-solid fa-circle-check"></i>
                        <span>Prepend text/icons</span>
                    </li>
                    <li>
                        <i class="fa-solid fa-circle-check"></i>
                        <span>Append text/icons</span>
                    </li>
                    <li>
                        <i class="fa-solid fa-circle-check"></i>
                        <span>Multiple addons</span>
                    </li>
                    <li>
                        <i class="fa-solid fa-circle-check"></i>
                        <span>Seamless styling</span>
                    </li>
                </ul>
            </div>
        </div>
    </div>
</div>

<!-- Button Input Groups -->
<div class="ig-section-title">
    <i class="fa-solid fa-hand-pointer"></i>
    Button Input Groups <span class="badge badge-success">Popular</span>
</div>

<div class="input-groups-grid">
    <div class="content-card">
        <div class="card-header">
            <div class="card-header-left">
                <div class="card-icon bg-primary">
                    <i class="fa-solid fa-square-plus"></i>
                </div>
                <div>
                    <h3>Button Append</h3>
                    <p class="card-subtitle">Button after input</p>
                </div>
            </div>
        </div>
        <div class="card-body">
            <div class="ig-example">
                <label class="ig-label">Search</label>
                <div class="input-group">
                    <input type="search" class="form-control" placeholder="Search...">
                    <button class="btn btn-primary">
                        <i class="fa-solid fa-magnifying-glass"></i>
                        Search
                    </button>
                </div>
            </div>

            <div class="ig-example">
                <label class="ig-label">Send Message</label>
                <div class="input-group">
                    <input type="text" class="form-control" placeholder="Type message...">
                    <button class="btn btn-success">
                        <i class="fa-solid fa-paper-plane"></i>
                        Send
                    </button>
                </div>
            </div>

            <div class="ig-example">
                <label class="ig-label">Subscribe</label>
                <div class="input-group">
                    <input type="email" class="form-control" placeholder="Enter your email">
                    <button class="btn btn-info">
                        Subscribe
                    </button>
                </div>
            </div>

            <div class="ig-example">
                <label class="ig-label">Copy to Clipboard</label>
                <div class="input-group">
                    <input type="text" class="form-control" value="https://example.com/abc123" readonly>
                    <button class="btn btn-secondary">
                        <i class="fa-solid fa-copy"></i>
                        Copy
                    </button>
                </div>
            </div>
        </div>
    </div>

    <div class="content-card">
        <div class="card-header">
            <div class="card-header-left">
                <div class="card-icon bg-success">
                    <i class="fa-solid fa-square-minus"></i>
                </div>
                <div>
                    <h3>Button Prepend</h3>
                    <p class="card-subtitle">Button before input</p>
                </div>
            </div>
        </div>
        <div class="card-body">
            <div class="ig-example">
                <label class="ig-label">Create New</label>
                <div class="input-group">
                    <button class="btn btn-primary">
                        <i class="fa-solid fa-plus"></i>
                        Add
                    </button>
                    <input type="text" class="form-control" placeholder="Item name">
                </div>
            </div>

            <div class="ig-example">
                <label class="ig-label">Upload File</label>
                <div class="input-group">
                    <button class="btn btn-warning">
                        <i class="fa-solid fa-upload"></i>
                        Upload
                    </button>
                    <input type="file" class="form-control">
                </div>
            </div>

            <div class="ig-example">
                <label class="ig-label">Download</label>
                <div class="input-group">
                    <button class="btn btn-success">
                        <i class="fa-solid fa-download"></i>
                        Get
                    </button>
                    <input type="text" class="form-control" placeholder="File URL" readonly>
                </div>
            </div>

            <div class="divider"></div>

            <div class="form-group">
                <h4 style="font-size: 14px; margin-bottom: 12px;">Button Variants:</h4>
                <ul class="feature-list">
                    <li>
                        <i class="fa-solid fa-circle-check"></i>
                        <span>Primary, Secondary, Success</span>
                    </li>
                    <li>
                        <i class="fa-solid fa-circle-check"></i>
                        <span>Warning, Danger, Info</span>
                    </li>
                    <li>
                        <i class="fa-solid fa-circle-check"></i>
                        <span>With icons</span>
                    </li>
                    <li>
                        <i class="fa-solid fa-circle-check"></i>
                        <span>Disabled state</span>
                    </li>
                </ul>
            </div>
        </div>
    </div>
</div>

<!-- Dropdown Input Groups -->
<div class="ig-section-title">
    <i class="fa-solid fa-caret-down"></i>
    Dropdown Input Groups
</div>

<div class="input-groups-grid">
    <div class="content-card">
        <div class="card-header">
            <div class="card-header-left">
                <div class="card-icon bg-info">
                    <i class="fa-solid fa-list"></i>
                </div>
                <div>
                    <h3>Dropdown Prepend</h3>
                    <p class="card-subtitle">Dropdown before input</p>
                </div>
            </div>
        </div>
        <div class="card-body">
            <div class="ig-example">
                <label class="ig-label">Protocol & URL</label>
                <div class="input-group">
                    <select class="form-control" style="max-width: 120px;">
                        <option>https://</option>
                        <option>http://</option>
                        <option>ftp://</option>
                    </select>
                    <input type="text" class="form-control" placeholder="www.example.com">
                </div>
            </div>

            <div class="ig-example">
                <label class="ig-label">Currency Converter</label>
                <div class="input-group">
                    <select class="form-control" style="max-width: 100px;">
                        <option>USD</option>
                        <option>EUR</option>
                        <option>GBP</option>
                        <option>IDR</option>
                    </select>
                    <input type="number" class="form-control" placeholder="0.00">
                </div>
            </div>

            <div class="ig-example">
                <label class="ig-label">Phone with Country</label>
                <div class="input-group">
                    <select class="form-control" style="max-width: 100px;">
                        <option>🇮🇩 +62</option>
                        <option>🇺🇸 +1</option>
                        <option>🇬🇧 +44</option>
                        <option>🇯🇵 +81</option>
                    </select>
                    <input type="tel" class="form-control" placeholder="812-3456-7890">
                </div>
            </div>
        </div>
    </div>

    <div class="content-card">
        <div class="card-header">
            <div class="card-header-left">
                <div class="card-icon bg-warning">
                    <i class="fa-solid fa-bars"></i>
                </div>
                <div>
                    <h3>Dropdown Append</h3>
                    <p class="card-subtitle">Dropdown after input</p>
                </div>
            </div>
        </div>
        <div class="card-body">
            <div class="ig-example">
                <label class="ig-label">Search Filter</label>
                <div class="input-group">
                    <input type="search" class="form-control" placeholder="Search...">
                    <select class="form-control" style="max-width: 130px;">
                        <option>All</option>
                        <option>Users</option>
                        <option>Products</option>
                        <option>Orders</option>
                    </select>
                </div>
            </div>

            <div class="ig-example">
                <label class="ig-label">Amount with Currency</label>
                <div class="input-group">
                    <input type="number" class="form-control" placeholder="0.00">
                    <select class="form-control" style="max-width: 100px;">
                        <option>USD</option>
                        <option>EUR</option>
                        <option>GBP</option>
                    </select>
                </div>
            </div>

            <div class="ig-example">
                <label class="ig-label">Email with Domain</label>
                <div class="input-group">
                    <input type="text" class="form-control" placeholder="username">
                    <select class="form-control" style="max-width: 150px;">
                        <option>@gmail.com</option>
                        <option>@yahoo.com</option>
                        <option>@outlook.com</option>
                        <option>@company.com</option>
                    </select>
                </div>
            </div>
        </div>
    </div>
</div>

<!-- Checkbox & Radio Input Groups -->
<div class="ig-section-title">
    <i class="fa-solid fa-square-check"></i>
    Checkbox & Radio Input Groups
</div>

<div class="input-groups-grid">
    <div class="content-card">
        <div class="card-header">
            <div class="card-header-left">
                <div class="card-icon bg-primary">
                    <i class="fa-solid fa-square-check"></i>
                </div>
                <div>
                    <h3>Checkbox Prepend</h3>
                    <p class="card-subtitle">Checkbox before input</p>
                </div>
            </div>
        </div>
        <div class="card-body">
            <div class="ig-example">
                <label class="ig-label">Agree to Terms</label>
                <div class="input-group">
                    <div class="input-group-text">
                        <input type="checkbox">
                    </div>
                    <input type="text" class="form-control" placeholder="I agree to the terms and conditions" readonly>
                </div>
            </div>

            <div class="ig-example">
                <label class="ig-label">Remember Me</label>
                <div class="input-group">
                    <div class="input-group-text">
                        <input type="checkbox" checked>
                    </div>
                    <input type="text" class="form-control" value="Keep me logged in" readonly>
                </div>
            </div>

            <div class="ig-example">
                <label class="ig-label">Subscribe Newsletter</label>
                <div class="input-group">
                    <div class="input-group-text">
                        <input type="checkbox">
                    </div>
                    <input type="email" class="form-control" placeholder="Enter email for newsletter">
                </div>
            </div>
        </div>
    </div>

    <div class="content-card">
        <div class="card-header">
            <div class="card-header-left">
                <div class="card-icon bg-success">
                    <i class="fa-solid fa-circle-dot"></i>
                </div>
                <div>
                    <h3>Radio Prepend</h3>
                    <p class="card-subtitle">Radio button before input</p>
                </div>
            </div>
        </div>
        <div class="card-body">
            <div class="ig-example">
                <label class="ig-label">Payment Method</label>
                <div class="input-group">
                    <div class="input-group-text">
                        <input type="radio" name="payment" checked>
                    </div>
                    <input type="text" class="form-control" value="Credit Card" readonly>
                </div>
            </div>

            <div class="ig-example">
                <label class="ig-label">Shipping Option</label>
                <div class="input-group">
                    <div class="input-group-text">
                        <input type="radio" name="shipping">
                    </div>
                    <input type="text" class="form-control" placeholder="Standard Shipping (5-7 days)">
                </div>
            </div>

            <div class="ig-example">
                <label class="ig-label">Contact Preference</label>
                <div class="input-group">
                    <div class="input-group-text">
                        <input type="radio" name="contact" checked>
                    </div>
                    <input type="email" class="form-control" placeholder="Email contact">
                </div>
            </div>
        </div>
    </div>
</div>

<!-- Segmented Buttons -->
<div class="ig-section-title">
    <i class="fa-solid fa-grip"></i>
    Segmented Buttons
</div>

<div class="input-groups-grid full-width">
    <div class="content-card">
        <div class="card-header">
            <div class="card-header-left">
                <div class="card-icon bg-info">
                    <i class="fa-solid fa-rectangle-list"></i>
                </div>
                <div>
                    <h3>Dropdown with Split Button</h3>
                    <p class="card-subtitle">Button + dropdown toggle</p>
                </div>
            </div>
        </div>
        <div class="card-body">
            <div style="display: grid; grid-template-columns: repeat(2, 1fr); gap: 20px;">
                <div>
                    <div class="ig-example">
                        <label class="ig-label">Action Dropdown</label>
                        <div class="input-group">
                            <input type="text" class="form-control" placeholder="Select action...">
                            <button class="btn btn-primary">Action</button>
                            <button class="btn btn-primary dropdown-toggle" style="border-radius: 0 var(--radius-sm) var(--radius-sm) 0;">
                                <i class="fa-solid fa-caret-down"></i>
                            </button>
                        </div>
                    </div>

                    <div class="ig-example">
                        <label class="ig-label">Search with Filter</label>
                        <div class="input-group">
                            <input type="search" class="form-control" placeholder="Search items...">
                            <button class="btn btn-success">
                                <i class="fa-solid fa-magnifying-glass"></i>
                            </button>
                            <button class="btn btn-success dropdown-toggle" style="border-radius: 0 var(--radius-sm) var(--radius-sm) 0;">
                                <i class="fa-solid fa-caret-down"></i>
                            </button>
                        </div>
                    </div>
                </div>

                <div>
                    <div class="form-group">
                        <h4 style="font-size: 14px; margin-bottom: 12px;">Input Group Features:</h4>
                        <ul class="feature-list">
                            <li>
                                <i class="fa-solid fa-circle-check"></i>
                                <span>Text addons (prepend/append)</span>
                            </li>
                            <li>
                                <i class="fa-solid fa-circle-check"></i>
                                <span>Icon addons</span>
                            </li>
                            <li>
                                <i class="fa-solid fa-circle-check"></i>
                                <span>Button integration</span>
                            </li>
                            <li>
                                <i class="fa-solid fa-circle-check"></i>
                                <span>Dropdown menus</span>
                            </li>
                            <li>
                                <i class="fa-solid fa-circle-check"></i>
                                <span>Checkbox/Radio addons</span>
                            </li>
                            <li>
                                <i class="fa-solid fa-circle-check"></i>
                                <span>Multiple addons support</span>
                            </li>
                        </ul>
                    </div>
                </div>
            </div>
        </div>
    </div>
</div>

<!-- Complete Form Example -->
<div class="ig-section-title">
    <i class="fa-solid fa-clipboard-list"></i>
    Complete Form Example
</div>

<div class="input-groups-grid full-width">
    <div class="content-card">
        <div class="card-header">
            <div class="card-header-left">
                <div class="card-icon bg-primary">
                    <i class="fa-solid fa-cart-shopping"></i>
                </div>
                <div>
                    <h3>E-Commerce Checkout Form</h3>
                    <p class="card-subtitle">Real-world example with various input groups</p>
                </div>
            </div>
        </div>
        <div class="card-body">
            <form style="max-width: 800px;">
                <div class="form-group">
                    <label style="font-weight: 600; margin-bottom: 8px; display: block;">Email Address <span style="color: var(--danger);">*</span></label>
                    <div class="input-group">
                        <span class="input-group-text">
                            <i class="fa-solid fa-envelope"></i>
                        </span>
                        <input type="email" class="form-control" placeholder="your@email.com" required>
                    </div>
                </div>

                <div class="form-group">
                    <label style="font-weight: 600; margin-bottom: 8px; display: block;">Phone Number</label>
                    <div class="input-group">
                        <select class="form-control" style="max-width: 120px;">
                            <option>🇮🇩 +62</option>
                            <option>🇺🇸 +1</option>
                            <option>🇬🇧 +44</option>
                        </select>
                        <input type="tel" class="form-control" placeholder="812-3456-7890">
                    </div>
                </div>

                <div class="form-group">
                    <label style="font-weight: 600; margin-bottom: 8px; display: block;">Website (Optional)</label>
                    <div class="input-group">
                        <span class="input-group-text">https://</span>
                        <input type="text" class="form-control" placeholder="www.yoursite.com">
                        <span class="input-group-text">.com</span>
                    </div>
                </div>

                <div class="divider"></div>

                <div class="form-group">
                    <label style="font-weight: 600; margin-bottom: 8px; display: block;">Payment Information <span style="color: var(--danger);">*</span></label>
                    <div class="input-group">
                        <span class="input-group-text">
                            <i class="fa-solid fa-credit-card"></i>
                        </span>
                        <input type="text" class="form-control" placeholder="Card number" maxlength="19" required>
                    </div>
                </div>

                <div style="display: grid; grid-template-columns: 1fr 1fr; gap: 16px;">
                    <div class="form-group">
                        <div class="input-group">
                            <input type="text" class="form-control" placeholder="MM/YY" maxlength="5" required>
                            <span class="input-group-text">
                                <i class="fa-solid fa-calendar"></i>
                            </span>
                        </div>
                    </div>
                    <div class="form-group">
                        <div class="input-group">
                            <input type="text" class="form-control" placeholder="CVC" maxlength="4" required>
                            <span class="input-group-text">
                                <i class="fa-solid fa-lock"></i>
                            </span>
                        </div>
                    </div>
                </div>

                <div class="form-group">
                    <label style="font-weight: 600; margin-bottom: 8px; display: block;">Promo Code</label>
                    <div class="input-group">
                        <input type="text" class="form-control" placeholder="Enter promo code">
                        <button class="btn btn-success">
                            Apply
                        </button>
                    </div>
                </div>

                <div class="form-group">
                    <label style="font-weight: 600; margin-bottom: 8px; display: block;">Donation Amount (Optional)</label>
                    <div class="input-group">
                        <span class="input-group-text">$
                        </span>
                        <input type="number" class="form-control" placeholder="0.00">
                        <span class="input-group-text">USD</span>
                    </div>
                </div>

                <div class="divider"></div>

                <div style="display: flex; gap: 12px; margin-top: 24px;">
                    <button type="submit" class="btn btn-primary" style="flex: 1;">
                        <i class="fa-solid fa-lock"></i>
                        Complete Purchase
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

<!-- Usage Guidelines -->
<div class="input-groups-grid full-width">
    <div class="content-card">
        <div class="card-header">
            <div class="card-header-left">
                <div class="card-icon bg-success">
                    <i class="fa-solid fa-lightbulb"></i>
                </div>
                <div>
                    <h3>When to Use Input Groups</h3>
                    <p class="card-subtitle">Best practices and use cases</p>
                </div>
            </div>
        </div>
        <div class="card-body">
            <div class="input-groups-grid three-cols">
                <div>
                    <h4 style="font-size: 14px; margin-bottom: 12px; display: flex; align-items: center; gap: 8px;">
                        <span style="font-size: 24px;">🔤</span>
                        Text Addons
                    </h4>
                    <ul class="feature-list">
                        <li><i class="fa-solid fa-check"></i> Email @ symbol</li>
                        <li><i class="fa-solid fa-check"></i> Currency symbols</li>
                        <li><i class="fa-solid fa-check"></i> URL protocols</li>
                        <li><i class="fa-solid fa-check"></i> File extensions</li>
                        <li><i class="fa-solid fa-check"></i> Units of measurement</li>
                    </ul>
                </div>
                <div>
                    <h4 style="font-size: 14px; margin-bottom: 12px; display: flex; align-items: center; gap: 8px;">
                        <span style="font-size: 24px;">🔘</span>
                        Button Addons
                    </h4>
                    <ul class="feature-list">
                        <li><i class="fa-solid fa-check"></i> Search buttons</li>
                        <li><i class="fa-solid fa-check"></i> Submit actions</li>
                        <li><i class="fa-solid fa-check"></i> Copy to clipboard</li>
                        <li><i class="fa-solid fa-check"></i> Upload triggers</li>
                        <li><i class="fa-solid fa-check"></i> Download links</li>
                    </ul>
                </div>
                <div>
                    <h4 style="font-size: 14px; margin-bottom: 12px; display: flex; align-items: center; gap: 8px;">
                        <span style="font-size: 24px;">📋</span>
                        Dropdown Addons
                    </h4>
                    <ul class="feature-list">
                        <li><i class="fa-solid fa-check"></i> Currency selection</li>
                        <li><i class="fa-solid fa-check"></i> Country codes</li>
                        <li><i class="fa-solid fa-check"></i> Protocol choice</li>
                        <li><i class="fa-solid fa-check"></i> Category filter</li>
                        <li><i class="fa-solid fa-check"></i> Domain selection</li>
                    </ul>
                </div>
            </div>
        </div>
    </div>
</div>
@endsection
