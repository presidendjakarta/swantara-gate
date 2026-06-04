@extends('layouts.app')

@section('title', 'Form Elements')

@push('styles')
@endpush

@section('content')
<div class="page-header">
    <div>
        <h1>Form Elements</h1>
        <p>Komponen form lengkap dengan berbagai tipe input, button, dan controls</p>
    </div>
    <div style="display: flex; gap: 10px;">
        <button class="btn btn-secondary">
            <i class="fa-solid fa-book"></i>
            <span>Documentation</span>
        </button>
        <button class="btn btn-primary">
            <i class="fa-solid fa-eye"></i>
            <span>Preview All</span>
        </button>
    </div>
</div>

<!-- Info Alert -->
<div class="content-card" style="margin-bottom: 24px; border-left: 4px solid var(--accent);">
    <div class="card-body" style="padding: 16px 20px;">
        <div style="display: flex; gap: 12px; align-items: start;">
            <i class="fa-solid fa-circle-info" style="color: var(--accent); font-size: 20px; margin-top: 2px;"></i>
            <div style="flex: 1;">
                <h4 style="margin-bottom: 4px; font-size: 14px;">Form Components Library</h4>
                <p style="font-size: 13px; color: var(--text-secondary); margin: 0;">Halaman ini menampilkan semua elemen form yang tersedia. Setiap komponen sudah responsive dan mengikuti Fluent Design System.</p>
            </div>
        </div>
    </div>
</div>

<!-- Basic Input Types -->
<div class="form-section-title">
    <i class="fa-solid fa-keyboard"></i>
    Basic Input Types
</div>

<div class="form-demo-grid">
    <!-- Text Inputs -->
    <div class="content-card">
        <div class="card-header">
            <div class="card-header-left">
                <div class="card-icon bg-primary">
                    <i class="fa-solid fa-font"></i>
                </div>
                <div>
                    <h3>Text Inputs <span class="badge badge-primary">Popular</span></h3>
                    <p class="card-subtitle">Input fields untuk text, email, password</p>
                </div>
            </div>
        </div>
        <div class="card-body">
            <div class="form-group">
                <label>Text Input <span style="color: var(--danger);">*</span></label>
                <input type="text" class="form-control" placeholder="Enter your name">
                <div class="helper-text">
                    <i class="fa-solid fa-circle-info"></i>
                    Masukkan nama lengkap Anda
                </div>
            </div>

            <div class="form-group">
                <label>Email Address</label>
                <div class="input-with-icon">
                    <i class="fa-solid fa-envelope"></i>
                    <input type="email" class="form-control" placeholder="name@example.com">
                </div>
            </div>

            <div class="form-group">
                <label>Password</label>
                <div class="input-with-icon">
                    <i class="fa-solid fa-lock"></i>
                    <input type="password" class="form-control" placeholder="••••••••">
                </div>
                <div class="helper-text">
                    <i class="fa-solid fa-shield-halved"></i>
                    Minimal 8 karakter
                </div>
            </div>

            <div class="form-group">
                <label>Search Input</label>
                <div class="input-with-icon">
                    <i class="fa-solid fa-magnifying-glass"></i>
                    <input type="search" class="form-control" placeholder="Search...">
                </div>
            </div>
        </div>
    </div>

    <!-- Number & Date Inputs -->
    <div class="content-card">
        <div class="card-header">
            <div class="card-header-left">
                <div class="card-icon bg-success">
                    <i class="fa-solid fa-hashtag"></i>
                </div>
                <div>
                    <h3>Number & Date</h3>
                    <p class="card-subtitle">Numeric and date pickers</p>
                </div>
            </div>
        </div>
        <div class="card-body">
            <div class="form-group">
                <label>Number Input</label>
                <input type="number" class="form-control" placeholder="0" min="0" max="100">
            </div>

            <div class="form-group">
                <label>Phone Number</label>
                <div class="input-with-icon">
                    <i class="fa-solid fa-phone"></i>
                    <input type="tel" class="form-control" placeholder="+62 812-3456-7890">
                </div>
            </div>

            <div class="form-group">
                <label>Date Picker</label>
                <input type="date" class="form-control">
            </div>

            <div class="form-group">
                <label>Time Picker</label>
                <input type="time" class="form-control">
            </div>

            <div class="form-group">
                <label>URL Input</label>
                <div class="input-with-icon">
                    <i class="fa-solid fa-link"></i>
                    <input type="url" class="form-control" placeholder="https://example.com">
                </div>
            </div>
        </div>
    </div>
</div>

<!-- Select & Dropdown -->
<div class="form-section-title">
    <i class="fa-solid fa-list-check"></i>
    Select & Dropdown
</div>

<div class="form-demo-grid">
    <div class="content-card">
        <div class="card-header">
            <div class="card-header-left">
                <div class="card-icon bg-info">
                    <i class="fa-solid fa-caret-down"></i>
                </div>
                <div>
                    <h3>Dropdown Select</h3>
                    <p class="card-subtitle">Single and multiple selection</p>
                </div>
            </div>
        </div>
        <div class="card-body">
            <div class="form-group">
                <label>Basic Select</label>
                <select class="form-control">
                    <option value="">-- Pilih Opsi --</option>
                    <option value="1">Opsi 1</option>
                    <option value="2">Opsi 2</option>
                    <option value="3">Opsi 3</option>
                    <option value="4">Opsi 4</option>
                </select>
            </div>

            <div class="form-group">
                <label>Select with Optgroup</label>
                <select class="form-control">
                    <option value="">-- Pilih --</option>
                    <optgroup label="Buah">
                        <option value="apel">Apel</option>
                        <option value="jeruk">Jeruk</option>
                        <option value="mangga">Mangga</option>
                    </optgroup>
                    <optgroup label="Sayur">
                        <option value="bayam">Bayam</option>
                        <option value="wortel">Wortel</option>
                    </optgroup>
                </select>
            </div>

            <div class="form-group">
                <label>Multiple Select <span class="badge badge-success">Multi</span></label>
                <select class="form-control" multiple size="4">
                    <option value="1">Option 1</option>
                    <option value="2">Option 2</option>
                    <option value="3">Option 3</option>
                    <option value="4">Option 4</option>
                    <option value="5">Option 5</option>
                </select>
                <div class="helper-text">
                    <i class="fa-solid fa-hand-pointer"></i>
                    Tahan Ctrl untuk pilih multiple
                </div>
            </div>
        </div>
    </div>

    <!-- Textarea -->
    <div class="content-card">
        <div class="card-header">
            <div class="card-header-left">
                <div class="card-icon bg-warning">
                    <i class="fa-solid fa-align-left"></i>
                </div>
                <div>
                    <h3>Textarea</h3>
                    <p class="card-subtitle">Multi-line text input</p>
                </div>
            </div>
        </div>
        <div class="card-body">
            <div class="form-group">
                <label>Basic Textarea</label>
                <textarea class="form-control" rows="3" placeholder="Enter your message..."></textarea>
            </div>

            <div class="form-group">
                <label>Large Textarea</label>
                <textarea class="form-control" rows="6" placeholder="Enter detailed description..."></textarea>
                <div class="helper-text">
                    <i class="fa-solid fa-circle-info"></i>
                    Maksimal 500 karakter
                </div>
            </div>

            <div class="form-group">
                <label>Textarea with Resize</label>
                <textarea class="form-control" rows="4" placeholder="Resizable textarea..." style="resize: both;"></textarea>
            </div>
        </div>
    </div>
</div>

<!-- Buttons Section -->
<div class="form-section-title">
    <i class="fa-solid fa-hand-pointer"></i>
    Buttons
</div>

<div class="form-demo-grid">
    <!-- Button Variants -->
    <div class="content-card">
        <div class="card-header">
            <div class="card-header-left">
                <div class="card-icon bg-primary">
                    <i class="fa-solid fa-square"></i>
                </div>
                <div>
                    <h3>Button Variants <span class="badge badge-primary">Essential</span></h3>
                    <p class="card-subtitle">Different button styles</p>
                </div>
            </div>
        </div>
        <div class="card-body">
            <div class="form-group">
                <label style="margin-bottom: 12px;">Solid Buttons</label>
                <div class="button-group">
                    <button class="btn btn-primary">Primary</button>
                    <button class="btn btn-secondary">Secondary</button>
                    <button class="btn btn-success">Success</button>
                    <button class="btn btn-danger">Danger</button>
                    <button class="btn btn-warning">Warning</button>
                    <button class="btn btn-info">Info</button>
                </div>
            </div>

            <div class="divider"></div>

            <div class="form-group">
                <label style="margin-bottom: 12px;">Outline Buttons</label>
                <div class="button-group">
                    <button class="btn btn-outline-primary">Primary</button>
                    <button class="btn btn-outline-secondary">Secondary</button>
                    <button class="btn btn-outline-success">Success</button>
                    <button class="btn btn-outline-danger">Danger</button>
                    <button class="btn btn-outline-warning">Warning</button>
                    <button class="btn btn-outline-info">Info</button>
                </div>
            </div>

            <div class="divider"></div>

            <div class="form-group">
                <label style="margin-bottom: 12px;">Button States</label>
                <div class="button-group">
                    <button class="btn btn-primary">Normal</button>
                    <button class="btn btn-success" disabled>
                        <i class="fa-solid fa-spinner fa-spin"></i>
                        Loading
                    </button>
                    <button class="btn btn-secondary" disabled>Disabled</button>
                    <button class="btn btn-danger">
                        <i class="fa-solid fa-trash"></i>
                        Delete
                    </button>
                    <button class="btn btn-info">
                        <i class="fa-solid fa-download"></i>
                        Download
                    </button>
                </div>
            </div>
        </div>
    </div>

    <!-- Button Sizes -->
    <div class="content-card">
        <div class="card-header">
            <div class="card-header-left">
                <div class="card-icon bg-success">
                    <i class="fa-solid fa-up-right-and-down-left-from-center"></i>
                </div>
                <div>
                    <h3>Button Sizes</h3>
                    <p class="card-subtitle">Small, medium, and large</p>
                </div>
            </div>
        </div>
        <div class="card-body">
            <div class="size-demo">
                <span class="size-label">Small</span>
                <button class="btn btn-primary btn-sm">Small Button</button>
            </div>

            <div class="size-demo">
                <span class="size-label">Medium</span>
                <button class="btn btn-primary">Medium Button</button>
            </div>

            <div class="size-demo">
                <span class="size-label">Large</span>
                <button class="btn btn-primary btn-lg">Large Button</button>
            </div>

            <div class="divider"></div>

            <div class="form-group">
                <label style="margin-bottom: 12px;">Icon Buttons</label>
                <div class="button-group">
                    <button class="btn btn-primary" title="Save">
                        <i class="fa-solid fa-floppy-disk"></i>
                    </button>
                    <button class="btn btn-success" title="Edit">
                        <i class="fa-solid fa-pen-to-square"></i>
                    </button>
                    <button class="btn btn-danger" title="Delete">
                        <i class="fa-solid fa-trash"></i>
                    </button>
                    <button class="btn btn-warning" title="Settings">
                        <i class="fa-solid fa-gear"></i>
                    </button>
                    <button class="btn btn-info" title="Refresh">
                        <i class="fa-solid fa-rotate-right"></i>
                    </button>
                </div>
            </div>

            <div class="divider"></div>

            <div class="form-group">
                <label style="margin-bottom: 12px;">Block Button</label>
                <button class="btn btn-primary" style="width: 100%;">
                    <i class="fa-solid fa-paper-plane"></i>
                    Submit Form
                </button>
            </div>
        </div>
    </div>
</div>

<!-- Checkbox & Radio -->
<div class="form-section-title">
    <i class="fa-solid fa-square-check"></i>
    Checkbox & Radio
</div>

<div class="form-demo-grid">
    <div class="content-card">
        <div class="card-header">
            <div class="card-header-left">
                <div class="card-icon bg-info">
                    <i class="fa-solid fa-check-square"></i>
                </div>
                <div>
                    <h3>Checkboxes</h3>
                    <p class="card-subtitle">Multiple selection options</p>
                </div>
            </div>
        </div>
        <div class="card-body">
            <div class="state-demo">
                <label style="display: flex; align-items: center; gap: 10px; cursor: pointer; padding: 8px 0;">
                    <input type="checkbox">
                    <span>Default checkbox</span>
                </label>
            </div>

            <div class="state-demo">
                <label style="display: flex; align-items: center; gap: 10px; cursor: pointer; padding: 8px 0;">
                    <input type="checkbox" checked>
                    <span>Checked checkbox</span>
                </label>
            </div>

            <div class="state-demo">
                <label style="display: flex; align-items: center; gap: 10px; cursor: pointer; padding: 8px 0; opacity: 0.5;">
                    <input type="checkbox" disabled>
                    <span>Disabled checkbox</span>
                </label>
            </div>

            <div class="state-demo">
                <label style="display: flex; align-items: center; gap: 10px; cursor: pointer; padding: 8px 0;">
                    <input type="checkbox" checked disabled>
                    <span>Checked & disabled</span>
                </label>
            </div>

            <div class="divider"></div>

            <div class="form-group">
                <label>Inline Checkboxes</label>
                <div style="display: flex; gap: 16px; flex-wrap: wrap;">
                    <label style="display: flex; align-items: center; gap: 6px; cursor: pointer;">
                        <input type="checkbox"> Option 1
                    </label>
                    <label style="display: flex; align-items: center; gap: 6px; cursor: pointer;">
                        <input type="checkbox"> Option 2
                    </label>
                    <label style="display: flex; align-items: center; gap: 6px; cursor: pointer;">
                        <input type="checkbox"> Option 3
                    </label>
                </div>
            </div>
        </div>
    </div>

    <div class="content-card">
        <div class="card-header">
            <div class="card-header-left">
                <div class="card-icon bg-warning">
                    <i class="fa-solid fa-circle-dot"></i>
                </div>
                <div>
                    <h3>Radio Buttons</h3>
                    <p class="card-subtitle">Single selection from group</p>
                </div>
            </div>
        </div>
        <div class="card-body">
            <div class="form-group">
                <label>Gender Selection</label>
                <div class="state-demo">
                    <label style="display: flex; align-items: center; gap: 10px; cursor: pointer; padding: 8px 0;">
                        <input type="radio" name="gender" value="male">
                        <span>Male</span>
                    </label>
                </div>
                <div class="state-demo">
                    <label style="display: flex; align-items: center; gap: 10px; cursor: pointer; padding: 8px 0;">
                        <input type="radio" name="gender" value="female" checked>
                        <span>Female (Selected)</span>
                    </label>
                </div>
                <div class="state-demo">
                    <label style="display: flex; align-items: center; gap: 10px; cursor: pointer; padding: 8px 0; opacity: 0.5;">
                        <input type="radio" name="gender" value="other" disabled>
                        <span>Disabled option</span>
                    </label>
                </div>
            </div>

            <div class="divider"></div>

            <div class="form-group">
                <label>Rating (1-5)</label>
                <div style="display: flex; gap: 12px;">
                    <label style="display: flex; align-items: center; gap: 6px; cursor: pointer;">
                        <input type="radio" name="rating" value="1"> 1
                    </label>
                    <label style="display: flex; align-items: center; gap: 6px; cursor: pointer;">
                        <input type="radio" name="rating" value="2"> 2
                    </label>
                    <label style="display: flex; align-items: center; gap: 6px; cursor: pointer;">
                        <input type="radio" name="rating" value="3" checked> 3
                    </label>
                    <label style="display: flex; align-items: center; gap: 6px; cursor: pointer;">
                        <input type="radio" name="rating" value="4"> 4
                    </label>
                    <label style="display: flex; align-items: center; gap: 6px; cursor: pointer;">
                        <input type="radio" name="rating" value="5"> 5
                    </label>
                </div>
            </div>
        </div>
    </div>
</div>

<!-- Input States & Validation -->
<div class="form-section-title">
    <i class="fa-solid fa-triangle-exclamation"></i>
    Input States & Validation
</div>

<div class="form-demo-grid full-width">
    <div class="content-card">
        <div class="card-header">
            <div class="card-header-left">
                <div class="card-icon bg-danger">
                    <i class="fa-solid fa-circle-exclamation"></i>
                </div>
                <div>
                    <h3>Validation States</h3>
                    <p class="card-subtitle">Error, success, and warning states</p>
                </div>
            </div>
        </div>
        <div class="card-body">
            <div class="form-demo-grid three-cols">
                <div>
                    <div class="state-demo">
                        <label>Default State</label>
                        <input type="text" class="form-control" placeholder="Normal input">
                        <div class="helper-text">
                            <i class="fa-solid fa-circle-info"></i>
                            This is the default state
                        </div>
                    </div>
                </div>

                <div>
                    <div class="state-demo">
                        <label style="color: var(--success);">Success State</label>
                        <input type="text" class="form-control" value="valid value" style="border-color: var(--success); background-color: var(--success-bg);">
                        <div class="helper-text success-text">
                            <i class="fa-solid fa-circle-check"></i>
                            Looks good!
                        </div>
                    </div>
                </div>

                <div>
                    <div class="state-demo">
                        <label style="color: var(--danger);">Error State</label>
                        <input type="text" class="form-control" value="invalid" style="border-color: var(--danger); background-color: var(--danger-bg);">
                        <div class="helper-text error-text">
                            <i class="fa-solid fa-circle-xmark"></i>
                            Please enter a valid value
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
                        <span>Real-time validation feedback</span>
                    </li>
                    <li>
                        <i class="fa-solid fa-circle-check"></i>
                        <span>Color-coded states (success, error, warning)</span>
                    </li>
                    <li>
                        <i class="fa-solid fa-circle-check"></i>
                        <span>Helper text and error messages</span>
                    </li>
                    <li>
                        <i class="fa-solid fa-circle-check"></i>
                        <span>Accessible form controls</span>
                    </li>
                    <li>
                        <i class="fa-solid fa-circle-check"></i>
                        <span>Responsive and mobile-friendly</span>
                    </li>
                </ul>
            </div>
        </div>
    </div>
</div>

<!-- Complete Form Example -->
<div class="form-section-title">
    <i class="fa-solid fa-file-pen"></i>
    Complete Form Example
</div>

<div class="content-card">
    <div class="card-header">
        <div class="card-header-left">
            <div class="card-icon bg-primary">
                <i class="fa-solid fa-user-plus"></i>
            </div>
            <div>
                <h3>Registration Form</h3>
                <p class="card-subtitle">Contoh form lengkap dengan berbagai elemen</p>
            </div>
        </div>
    </div>
    <div class="card-body">
        <form style="max-width: 800px;">
            <div style="display: grid; grid-template-columns: 1fr 1fr; gap: 20px;">
                <div class="form-group">
                    <label>Nama Depan <span style="color: var(--danger);">*</span></label>
                    <input type="text" class="form-control" placeholder="John" required>
                </div>
                <div class="form-group">
                    <label>Nama Belakang <span style="color: var(--danger);">*</span></label>
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
                <label>Password <span style="color: var(--danger);">*</span></label>
                <div class="input-with-icon">
                    <i class="fa-solid fa-lock"></i>
                    <input type="password" class="form-control" placeholder="Min. 8 characters" required>
                </div>
            </div>

            <div class="form-group">
                <label>Jenis Kelamin</label>
                <div style="display: flex; gap: 20px;">
                    <label style="display: flex; align-items: center; gap: 8px; cursor: pointer;">
                        <input type="radio" name="gender" value="male">
                        <span>Male</span>
                    </label>
                    <label style="display: flex; align-items: center; gap: 8px; cursor: pointer;">
                        <input type="radio" name="gender" value="female">
                        <span>Female</span>
                    </label>
                </div>
            </div>

            <div class="form-group">
                <label>Kota</label>
                <select class="form-control">
                    <option value="">-- Pilih Kota --</option>
                    <option value="jakarta">Jakarta</option>
                    <option value="surabaya">Surabaya</option>
                    <option value="bandung">Bandung</option>
                    <option value="medan">Medan</option>
                </select>
            </div>

            <div class="form-group">
                <label>Bio</label>
                <textarea class="form-control" rows="3" placeholder="Tell us about yourself..."></textarea>
            </div>

            <div class="form-group">
                <label style="display: flex; align-items: center; gap: 8px; cursor: pointer;">
                    <input type="checkbox" required>
                    <span>Saya setuju dengan <a href="#" style="color: var(--accent);">syarat dan ketentuan</a></span>
                </label>
            </div>

            <div style="display: flex; gap: 12px; margin-top: 24px;">
                <button type="submit" class="btn btn-primary">
                    <i class="fa-solid fa-paper-plane"></i>
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
@endsection
