@extends('layouts.app')

@section('title', 'Default Forms')

@push('styles')
@endpush

@section('content')
<div class="page-header">
    <div>
        <h1>Default Forms</h1>
        <p>Complete form layouts with validation, horizontal, inline, and multi-column designs</p>
    </div>
    <div style="display: flex; gap: 10px;">
        <button class="btn btn-secondary">
            <i class="fa-solid fa-book"></i>
            <span>Documentation</span>
        </button>
        <button class="btn btn-primary">
            <i class="fa-solid fa-file-pen"></i>
            <span>Try Forms</span>
        </button>
    </div>
</div>

<!-- Info Alert -->
<div class="content-card" style="margin-bottom: 24px; border-left: 4px solid var(--accent);">
    <div class="card-body" style="padding: 16px 20px;">
        <div style="display: flex; gap: 12px; align-items: start;">
            <i class="fa-solid fa-circle-info" style="color: var(--accent); font-size: 20px; margin-top: 2px;"></i>
            <div style="flex: 1;">
                <h4 style="margin-bottom: 4px; font-size: 14px;">Default Forms Guide</h4>
                <p style="font-size: 13px; color: var(--text-secondary); margin: 0;">Default forms provide standard form layouts including vertical, horizontal, inline, and multi-column designs. Includes validation states, required fields, helper text, and form sections for better UX.</p>
            </div>
        </div>
    </div>
</div>

<!-- Basic Form -->
<div class="df-section-title">
    <i class="fa-solid fa-file-pen"></i>
    Basic Forms <span class="badge badge-primary">Essential</span>
</div>

<div class="df-grid">
    <div class="content-card">
        <div class="card-header">
            <div class="card-header-left">
                <div class="card-icon bg-primary">
                    <i class="fa-solid fa-user-plus"></i>
                </div>
                <div>
                    <h3>Registration Form</h3>
                    <p class="card-subtitle">Vertical layout form</p>
                </div>
            </div>
        </div>
        <div class="card-body">
            <form style="max-width: 600px;">
                <div class="df-example">
                    <div class="form-row form-row-2">
                        <div class="form-group">
                            <label class="df-label">
                                First Name
                                <span style="color: var(--danger);">*</span>
                            </label>
                            <input type="text" class="form-control" placeholder="John" required>
                        </div>
                        <div class="form-group">
                            <label class="df-label">
                                Last Name
                                <span style="color: var(--danger);">*</span>
                            </label>
                            <input type="text" class="form-control" placeholder="Doe" required>
                        </div>
                    </div>

                    <div class="form-group">
                        <label class="df-label">
                            Email Address
                            <span style="color: var(--danger);">*</span>
                        </label>
                        <input type="email" class="form-control" placeholder="john@example.com" required>
                        <div class="df-helper">
                            <i class="fa-solid fa-circle-info"></i>
                            We'll never share your email
                        </div>
                    </div>

                    <div class="form-group">
                        <label class="df-label">
                            Password
                            <span style="color: var(--danger);">*</span>
                        </label>
                        <input type="password" class="form-control" placeholder="••••••••" required>
                    </div>

                    <div class="form-group">
                        <label class="df-label">
                            Confirm Password
                            <span style="color: var(--danger);">*</span>
                        </label>
                        <input type="password" class="form-control" placeholder="••••••••" required>
                    </div>

                    <div class="form-group">
                        <label class="df-label">Gender</label>
                        <select class="form-control">
                            <option value="">Select gender...</option>
                            <option value="male">Male</option>
                            <option value="female">Female</option>
                            <option value="other">Other</option>
                        </select>
                    </div>

                    <div class="form-group">
                        <label style="display: flex; align-items: start; gap: 8px; cursor: pointer;">
                            <input type="checkbox" style="margin-top: 2px;">
                            <span style="font-size: 13px;">I agree to the <a href="#" style="color: var(--accent);">Terms of Service</a> and <a href="#" style="color: var(--accent);">Privacy Policy</a></span>
                        </label>
                    </div>

                    <div class="form-actions">
                        <button type="submit" class="btn btn-primary" style="flex: 1;">
                            <i class="fa-solid fa-user-plus"></i>
                            Create Account
                        </button>
                        <button type="reset" class="btn btn-secondary">
                            <i class="fa-solid fa-rotate-left"></i>
                            Reset
                        </button>
                    </div>
                </div>
            </form>

            <div class="form-divider"></div>

            <div class="form-group">
                <h4 style="font-size: 14px; margin-bottom: 12px;">Features:</h4>
                <ul class="feature-list">
                    <li>
                        <i class="fa-solid fa-circle-check"></i>
                        <span>Vertical layout</span>
                    </li>
                    <li>
                        <i class="fa-solid fa-circle-check"></i>
                        <span>Required field markers</span>
                    </li>
                    <li>
                        <i class="fa-solid fa-circle-check"></i>
                        <span>Helper text</span>
                    </li>
                    <li>
                        <i class="fa-solid fa-circle-check"></i>
                        <span>Form validation</span>
                    </li>
                </ul>
            </div>
        </div>
    </div>

    <div class="content-card">
        <div class="card-header">
            <div class="card-header-left">
                <div class="card-icon bg-success">
                    <i class="fa-solid fa-envelope"></i>
                </div>
                <div>
                    <h3>Contact Form</h3>
                    <p class="card-subtitle">Customer support form</p>
                </div>
            </div>
        </div>
        <div class="card-body">
            <form style="max-width: 600px;">
                <div class="form-row form-row-2">
                    <div class="form-group">
                        <label class="df-label">Your Name</label>
                        <input type="text" class="form-control" placeholder="John Doe">
                    </div>
                    <div class="form-group">
                        <label class="df-label">Phone Number</label>
                        <input type="tel" class="form-control" placeholder="+1 (555) 000-0000">
                    </div>
                </div>

                <div class="form-group">
                    <label class="df-label">Email Address</label>
                    <input type="email" class="form-control" placeholder="john@example.com">
                </div>

                <div class="form-group">
                    <label class="df-label">Subject</label>
                    <select class="form-control">
                        <option value="">Select subject...</option>
                        <option value="general">General Inquiry</option>
                        <option value="support">Technical Support</option>
                        <option value="billing">Billing Question</option>
                        <option value="feedback">Feedback</option>
                    </select>
                </div>

                <div class="form-group">
                    <label class="df-label">Priority</label>
                    <div style="display: flex; gap: 12px;">
                        <label style="display: flex; align-items: center; gap: 6px; cursor: pointer;">
                            <input type="radio" name="priority" value="low">
                            <span style="font-size: 13px;">Low</span>
                        </label>
                        <label style="display: flex; align-items: center; gap: 6px; cursor: pointer;">
                            <input type="radio" name="priority" value="medium" checked>
                            <span style="font-size: 13px;">Medium</span>
                        </label>
                        <label style="display: flex; align-items: center; gap: 6px; cursor: pointer;">
                            <input type="radio" name="priority" value="high">
                            <span style="font-size: 13px;">High</span>
                        </label>
                    </div>
                </div>

                <div class="form-group">
                    <label class="df-label">Message</label>
                    <textarea class="form-control" rows="5" placeholder="Describe your issue..."></textarea>
                    <div class="df-helper">
                        <i class="fa-solid fa-circle-info"></i>
                        Please provide as much detail as possible
                    </div>
                </div>

                <div class="form-actions">
                    <button type="submit" class="btn btn-success" style="flex: 1;">
                        <i class="fa-solid fa-paper-plane"></i>
                        Send Message
                    </button>
                </div>
            </form>

            <div class="form-divider"></div>

            <div class="helper-text">
                <i class="fa-solid fa-circle-info"></i>
                All fields are optional unless marked
            </div>
        </div>
    </div>
</div>

<!-- Horizontal Form -->
<div class="df-section-title">
    <i class="fa-solid fa-bars"></i>
    Horizontal Forms
</div>

<div class="df-grid">
    <div class="content-card">
        <div class="card-header">
            <div class="card-header-left">
                <div class="card-icon bg-info">
                    <i class="fa-solid fa-user-gear"></i>
                </div>
                <div>
                    <h3>Profile Settings</h3>
                    <p class="card-subtitle">Horizontal layout</p>
                </div>
            </div>
        </div>
        <div class="card-body">
            <form class="form-horizontal">
                <div class="form-group">
                    <label class="df-label">Username</label>
                    <div>
                        <input type="text" class="form-control" value="johndoe" placeholder="Username">
                    </div>
                </div>

                <div class="form-group">
                    <label class="df-label">Full Name</label>
                    <div>
                        <input type="text" class="form-control" value="John Doe" placeholder="Full name">
                    </div>
                </div>

                <div class="form-group">
                    <label class="df-label">Email</label>
                    <div>
                        <input type="email" class="form-control" value="john@example.com" placeholder="Email address">
                    </div>
                </div>

                <div class="form-group">
                    <label class="df-label">Phone</label>
                    <div>
                        <input type="tel" class="form-control" placeholder="+1 (555) 000-0000">
                    </div>
                </div>

                <div class="form-group">
                    <label class="df-label">Country</label>
                    <div>
                        <select class="form-control">
                            <option value="">Select country...</option>
                            <option value="us" selected>United States</option>
                            <option value="uk">United Kingdom</option>
                            <option value="ca">Canada</option>
                            <option value="au">Australia</option>
                        </select>
                    </div>
                </div>

                <div class="form-group">
                    <label class="df-label">Bio</label>
                    <div>
                        <textarea class="form-control" rows="3" placeholder="Tell us about yourself..."></textarea>
                    </div>
                </div>

                <div class="form-actions">
                    <button type="submit" class="btn btn-primary">
                        <i class="fa-solid fa-check"></i>
                        Save Changes
                    </button>
                    <button type="button" class="btn btn-secondary">
                        <i class="fa-solid fa-xmark"></i>
                        Cancel
                    </button>
                </div>
            </form>

            <div class="form-divider"></div>

            <div class="helper-text">
                <i class="fa-solid fa-circle-info"></i>
                Labels aligned left, inputs fill remaining space
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
                    <h3>Change Password</h3>
                    <p class="card-subtitle">Security settings</p>
                </div>
            </div>
        </div>
        <div class="card-body">
            <form class="form-horizontal">
                <div class="form-group">
                    <label class="df-label">Current Password</label>
                    <div>
                        <input type="password" class="form-control" placeholder="Enter current password">
                    </div>
                </div>

                <div class="form-group">
                    <label class="df-label">New Password</label>
                    <div>
                        <input type="password" class="form-control" placeholder="Enter new password">
                        <div class="df-helper">
                            <i class="fa-solid fa-circle-info"></i>
                            Minimum 8 characters
                        </div>
                    </div>
                </div>

                <div class="form-group">
                    <label class="df-label">Confirm Password</label>
                    <div>
                        <input type="password" class="form-control" placeholder="Confirm new password">
                    </div>
                </div>

                <div class="form-group">
                    <label class="df-label">Two-Factor Auth</label>
                    <div style="display: flex; align-items: center; gap: 12px;">
                        <label style="display: flex; align-items: center; gap: 6px; cursor: pointer;">
                            <input type="checkbox">
                            <span style="font-size: 13px;">Enable 2FA</span>
                        </label>
                    </div>
                </div>

                <div class="form-actions">
                    <button type="submit" class="btn btn-warning">
                        <i class="fa-solid fa-key"></i>
                        Update Password
                    </button>
                    <button type="button" class="btn btn-secondary">
                        <i class="fa-solid fa-xmark"></i>
                        Cancel
                    </button>
                </div>
            </form>

            <div class="form-divider"></div>

            <div class="helper-text">
                <i class="fa-solid fa-circle-info"></i>
                Password must be at least 8 characters
            </div>
        </div>
    </div>
</div>

<!-- Inline Form -->
<div class="df-section-title">
    <i class="fa-solid fa-grip"></i>
    Inline Forms
</div>

<div class="df-grid full-width">
    <div class="content-card">
        <div class="card-header">
            <div class="card-header-left">
                <div class="card-icon bg-primary">
                    <i class="fa-solid fa-magnifying-glass"></i>
                </div>
                <div>
                    <h3>Search Form</h3>
                    <p class="card-subtitle">Inline layout for filters</p>
                </div>
            </div>
        </div>
        <div class="card-body">
            <form class="form-inline">
                <div class="form-group">
                    <label class="df-label">Keyword</label>
                    <input type="text" class="form-control" placeholder="Search...">
                </div>
                <div class="form-group">
                    <label class="df-label">Category</label>
                    <select class="form-control">
                        <option value="">All Categories</option>
                        <option value="electronics">Electronics</option>
                        <option value="clothing">Clothing</option>
                        <option value="books">Books</option>
                    </select>
                </div>
                <div class="form-group">
                    <label class="df-label">Status</label>
                    <select class="form-control">
                        <option value="">All Status</option>
                        <option value="active">Active</option>
                        <option value="inactive">Inactive</option>
                        <option value="pending">Pending</option>
                    </select>
                </div>
                <div class="form-group">
                    <label class="df-label">Date Range</label>
                    <input type="date" class="form-control">
                </div>
                <button type="submit" class="btn btn-primary">
                    <i class="fa-solid fa-magnifying-glass"></i>
                    Search
                </button>
                <button type="reset" class="btn btn-secondary">
                    <i class="fa-solid fa-rotate-left"></i>
                    Reset
                </button>
            </form>

            <div class="form-divider"></div>

            <div class="helper-text">
                <i class="fa-solid fa-circle-info"></i>
                All fields align horizontally for compact filters
            </div>
        </div>
    </div>
</div>

<!-- Multi-Column Forms -->
<div class="df-section-title">
    <i class="fa-solid fa-table-columns"></i>
    Multi-Column Forms
</div>

<div class="df-grid full-width">
    <div class="content-card">
        <div class="card-header">
            <div class="card-header-left">
                <div class="card-icon bg-success">
                    <i class="fa-solid fa-credit-card"></i>
                </div>
                <div>
                    <h3>Payment Form</h3>
                    <p class="card-subtitle">Multi-column layout</p>
                </div>
            </div>
        </div>
        <div class="card-body">
            <form>
                <div class="form-section">
                    <i class="fa-solid fa-user"></i>
                    Billing Information
                </div>
                
                <div class="form-row form-row-2">
                    <div class="form-group">
                        <label class="df-label">First Name</label>
                        <input type="text" class="form-control" placeholder="John">
                    </div>
                    <div class="form-group">
                        <label class="df-label">Last Name</label>
                        <input type="text" class="form-control" placeholder="Doe">
                    </div>
                </div>

                <div class="form-group">
                    <label class="df-label">Email</label>
                    <input type="email" class="form-control" placeholder="john@example.com">
                </div>

                <div class="form-divider"></div>

                <div class="form-section">
                    <i class="fa-solid fa-credit-card"></i>
                    Payment Details
                </div>

                <div class="form-group">
                    <label class="df-label">Card Number</label>
                    <input type="text" class="form-control" placeholder="1234 5678 9012 3456" maxlength="19">
                </div>

                <div class="form-row form-row-3">
                    <div class="form-group">
                        <label class="df-label">Expiry Month</label>
                        <select class="form-control">
                            <option value="">MM</option>
                            <option value="01">01</option>
                            <option value="02">02</option>
                            <option value="03">03</option>
                        </select>
                    </div>
                    <div class="form-group">
                        <label class="df-label">Expiry Year</label>
                        <select class="form-control">
                            <option value="">YYYY</option>
                            <option value="2024">2024</option>
                            <option value="2025">2025</option>
                            <option value="2026">2026</option>
                        </select>
                    </div>
                    <div class="form-group">
                        <label class="df-label">CVV</label>
                        <input type="text" class="form-control" placeholder="123" maxlength="4">
                    </div>
                </div>

                <div class="form-group">
                    <label class="df-label">Cardholder Name</label>
                    <input type="text" class="form-control" placeholder="John Doe">
                </div>

                <div class="form-actions">
                    <button type="submit" class="btn btn-success" style="flex: 1;">
                        <i class="fa-solid fa-lock"></i>
                        Pay $99.00
                    </button>
                    <button type="button" class="btn btn-secondary">
                        <i class="fa-solid fa-xmark"></i>
                        Cancel
                    </button>
                </div>
            </form>

            <div class="form-divider"></div>

            <div class="helper-text">
                <i class="fa-solid fa-circle-info"></i>
                Payment information is encrypted and secure
            </div>
        </div>
    </div>
</div>

<!-- Complete Form Example -->
<div class="df-section-title">
    <i class="fa-solid fa-clipboard-list"></i>
    Complete Form Example
</div>

<div class="df-grid full-width">
    <div class="content-card">
        <div class="card-header">
            <div class="card-header-left">
                <div class="card-icon bg-primary">
                    <i class="fa-solid fa-briefcase"></i>
                </div>
                <div>
                    <h3>Job Application</h3>
                    <p class="card-subtitle">Comprehensive application form</p>
                </div>
            </div>
        </div>
        <div class="card-body">
            <form style="max-width: 900px;">
                <div class="form-section">
                    <i class="fa-solid fa-user"></i>
                    Personal Information
                </div>

                <div class="form-row form-row-2">
                    <div class="form-group">
                        <label class="df-label">
                            First Name
                            <span style="color: var(--danger);">*</span>
                        </label>
                        <input type="text" class="form-control" placeholder="John" required>
                    </div>
                    <div class="form-group">
                        <label class="df-label">
                            Last Name
                            <span style="color: var(--danger);">*</span>
                        </label>
                        <input type="text" class="form-control" placeholder="Doe" required>
                    </div>
                </div>

                <div class="form-row form-row-3">
                    <div class="form-group">
                        <label class="df-label">
                            Email
                            <span style="color: var(--danger);">*</span>
                        </label>
                        <input type="email" class="form-control" placeholder="john@example.com" required>
                    </div>
                    <div class="form-group">
                        <label class="df-label">Phone</label>
                        <input type="tel" class="form-control" placeholder="+1 (555) 000-0000">
                    </div>
                    <div class="form-group">
                        <label class="df-label">Date of Birth</label>
                        <input type="date" class="form-control">
                    </div>
                </div>

                <div class="form-divider"></div>

                <div class="form-section">
                    <i class="fa-solid fa-graduation-cap"></i>
                    Education
                </div>

                <div class="form-group">
                    <label class="df-label">Highest Education</label>
                    <select class="form-control">
                        <option value="">Select education...</option>
                        <option value="highschool">High School</option>
                        <option value="bachelor">Bachelor's Degree</option>
                        <option value="master">Master's Degree</option>
                        <option value="phd">PhD</option>
                    </select>
                </div>

                <div class="form-row form-row-2">
                    <div class="form-group">
                        <label class="df-label">University/College</label>
                        <input type="text" class="form-control" placeholder="University name">
                    </div>
                    <div class="form-group">
                        <label class="df-label">Graduation Year</label>
                        <input type="number" class="form-control" placeholder="2020" min="1950" max="2026">
                    </div>
                </div>

                <div class="form-divider"></div>

                <div class="form-section">
                    <i class="fa-solid fa-briefcase"></i>
                    Work Experience
                </div>

                <div class="form-group">
                    <label class="df-label">Current/Last Position</label>
                    <input type="text" class="form-control" placeholder="Software Engineer">
                </div>

                <div class="form-row form-row-2">
                    <div class="form-group">
                        <label class="df-label">Company Name</label>
                        <input type="text" class="form-control" placeholder="Company Inc.">
                    </div>
                    <div class="form-group">
                        <label class="df-label">Years of Experience</label>
                        <select class="form-control">
                            <option value="">Select...</option>
                            <option value="0-1">0-1 years</option>
                            <option value="1-3">1-3 years</option>
                            <option value="3-5">3-5 years</option>
                            <option value="5+">5+ years</option>
                        </select>
                    </div>
                </div>

                <div class="form-group">
                    <label class="df-label">Cover Letter</label>
                    <textarea class="form-control" rows="5" placeholder="Tell us why you're interested in this position..."></textarea>
                </div>

                <div class="form-divider"></div>

                <div class="form-section">
                    <i class="fa-solid fa-paperclip"></i>
                    Attachments
                </div>

                <div class="form-row form-row-2">
                    <div class="form-group">
                        <label class="df-label">Resume/CV</label>
                        <input type="file" class="form-control" accept=".pdf,.doc,.docx">
                        <div class="df-helper">
                            <i class="fa-solid fa-circle-info"></i>
                            PDF, DOC, DOCX (max 5MB)
                        </div>
                    </div>
                    <div class="form-group">
                        <label class="df-label">Portfolio (Optional)</label>
                        <input type="file" class="form-control" accept=".pdf" multiple>
                        <div class="df-helper">
                            <i class="fa-solid fa-circle-info"></i>
                            PDF files (max 10MB total)
                        </div>
                    </div>
                </div>

                <div class="form-divider"></div>

                <div class="form-group">
                    <label style="display: flex; align-items: start; gap: 8px; cursor: pointer;">
                        <input type="checkbox" style="margin-top: 2px;">
                        <span style="font-size: 13px;">I confirm that all information provided is accurate and complete. I understand that false information may result in disqualification.</span>
                    </label>
                </div>

                <div class="form-group">
                    <label style="display: flex; align-items: start; gap: 8px; cursor: pointer;">
                        <input type="checkbox" style="margin-top: 2px;">
                        <span style="font-size: 13px;">I agree to the <a href="#" style="color: var(--accent);">Privacy Policy</a> and consent to the processing of my personal data.</span>
                    </label>
                </div>

                <div class="form-actions">
                    <button type="submit" class="btn btn-primary" style="flex: 1;">
                        <i class="fa-solid fa-paper-plane"></i>
                        Submit Application
                    </button>
                    <button type="reset" class="btn btn-secondary">
                        <i class="fa-solid fa-rotate-left"></i>
                        Clear Form
                    </button>
                </div>
            </form>
        </div>
    </div>
</div>

<!-- Implementation Guide -->
<div class="df-grid full-width">
    <div class="content-card">
        <div class="card-header">
            <div class="card-header-left">
                <div class="card-icon bg-success">
                    <i class="fa-solid fa-code"></i>
                </div>
                <div>
                    <h3>Implementation Guide</h3>
                    <p class="card-subtitle">Best practices for form design</p>
                </div>
            </div>
        </div>
        <div class="card-body">
            <div class="code-block">
                <div style="color: var(--text-tertiary); margin-bottom: 8px;">1. Vertical Form (Default):</div>
                <code style="color: var(--accent);">
                    &lt;form&gt;<br>
                    &nbsp;&nbsp;&lt;div class="form-group"&gt;<br>
                    &nbsp;&nbsp;&nbsp;&nbsp;&lt;label&gt;Field Name&lt;/label&gt;<br>
                    &nbsp;&nbsp;&nbsp;&nbsp;&lt;input type="text" class="form-control"&gt;<br>
                    &nbsp;&nbsp;&lt;/div&gt;<br>
                    &lt;/form&gt;
                </code>
            </div>

            <div class="code-block">
                <div style="color: var(--text-tertiary); margin-bottom: 8px; margin-top: 16px;">2. Horizontal Form:</div>
                <code style="color: var(--success);">
                    &lt;form class="form-horizontal"&gt;<br>
                    &nbsp;&nbsp;&lt;div class="form-group"&gt;<br>
                    &nbsp;&nbsp;&nbsp;&nbsp;&lt;label class="df-label"&gt;Field&lt;/label&gt;<br>
                    &nbsp;&nbsp;&nbsp;&nbsp;&lt;div&gt;<br>
                    &nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&lt;input type="text" class="form-control"&gt;<br>
                    &nbsp;&nbsp;&nbsp;&nbsp;&lt;/div&gt;<br>
                    &nbsp;&nbsp;&lt;/div&gt;<br>
                    &lt;/form&gt;
                </code>
            </div>

            <div class="code-block">
                <div style="color: var(--text-tertiary); margin-bottom: 8px; margin-top: 16px;">3. Multi-Column Form:</div>
                <code style="color: var(--warning);">
                    &lt;div class="form-row form-row-2"&gt;<br>
                    &nbsp;&nbsp;&lt;div class="form-group"&gt;<br>
                    &nbsp;&nbsp;&nbsp;&nbsp;&lt;label&gt;First Name&lt;/label&gt;<br>
                    &nbsp;&nbsp;&nbsp;&nbsp;&lt;input type="text" class="form-control"&gt;<br>
                    &nbsp;&nbsp;&lt;/div&gt;<br>
                    &nbsp;&nbsp;&lt;div class="form-group"&gt;<br>
                    &nbsp;&nbsp;&nbsp;&nbsp;&lt;label&gt;Last Name&lt;/label&gt;<br>
                    &nbsp;&nbsp;&nbsp;&nbsp;&lt;input type="text" class="form-control"&gt;<br>
                    &nbsp;&nbsp;&lt;/div&gt;<br>
                    &lt;/div&gt;
                </code>
            </div>

            <div class="code-block">
                <div style="color: var(--text-tertiary); margin-bottom: 8px; margin-top: 16px;">4. Inline Form:</div>
                <code style="color: var(--info);">
                    &lt;form class="form-inline"&gt;<br>
                    &nbsp;&nbsp;&lt;div class="form-group"&gt;<br>
                    &nbsp;&nbsp;&nbsp;&nbsp;&lt;label&gt;Keyword&lt;/label&gt;<br>
                    &nbsp;&nbsp;&nbsp;&nbsp;&lt;input type="text" class="form-control"&gt;<br>
                    &nbsp;&nbsp;&lt;/div&gt;<br>
                    &nbsp;&nbsp;&lt;button type="submit" class="btn btn-primary"&gt;Search&lt;/button&gt;<br>
                    &lt;/form&gt;
                </code>
            </div>

            <div class="form-divider"></div>

            <div class="form-group">
                <h4 style="font-size: 14px; margin-bottom: 12px;">Form Best Practices:</h4>
                <ul class="feature-list">
                    <li>
                        <i class="fa-solid fa-circle-check"></i>
                        <span><strong>Use labels</strong> - Always include descriptive labels</span>
                    </li>
                    <li>
                        <i class="fa-solid fa-circle-check"></i>
                        <span><strong>Mark required</strong> - Use asterisks (*) for required fields</span>
                    </li>
                    <li>
                        <i class="fa-solid fa-circle-check"></i>
                        <span><strong>Helper text</strong> - Provide guidance for complex fields</span>
                    </li>
                    <li>
                        <i class="fa-solid fa-circle-check"></i>
                        <span><strong>Group logically</strong> - Organize related fields together</span>
                    </li>
                    <li>
                        <i class="fa-solid fa-circle-check"></i>
                        <span><strong>Clear actions</strong> - Use descriptive button labels</span>
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
    // Form validation feedback
    $('form').on('submit', function(e) {
        e.preventDefault();
        
        const form = this;
        let isValid = true;
        
        // Check required fields
        $(form).find('[required]').each(function() {
            if (!$(this).val()) {
                isValid = false;
                $(this).addClass('is-invalid');
                
                // Add invalid feedback if not exists
                if (!$(this).next('.invalid-feedback').length) {
                    $(this).after('<div class="invalid-feedback">This field is required</div>');
                }
            } else {
                $(this).removeClass('is-invalid').addClass('is-valid');
            }
        });
        
        if (isValid) {
            showToast('Form submitted successfully!', 'success');
            
            // Optional: Reset form after 1 second
            setTimeout(() => {
                // form.reset();
                // $(form).find('.is-valid').removeClass('is-valid');
            }, 1000);
        } else {
            showError('Please fill in all required fields');
        }
    });
    
    // Remove validation state on input
    $('form').on('input change', '.form-control', function() {
        $(this).removeClass('is-invalid');
        
        if ($(this).val()) {
            $(this).addClass('is-valid');
        } else {
            $(this).removeClass('is-valid');
        }
    });
    
    // Card number formatting
    $('input[placeholder="1234 5678 9012 3456"]').on('input', function() {
        let value = $(this).val().replace(/\s/g, '').replace(/[^0-9]/gi, '');
        let formattedValue = value.match(/.{1,4}/g)?.join(' ') || '';
        $(this).val(formattedValue);
    });
    
    // CVV formatting (numbers only)
    $('input[placeholder="123"]').on('input', function() {
        $(this).val($(this).val().replace(/[^0-9]/g, ''));
    });
    
    // Phone number formatting
    $('input[type="tel"]').on('input', function() {
        let value = $(this).val().replace(/[^0-9+\-\s()]/g, '');
        $(this).val(value);
    });
    
    // Form reset confirmation
    $('button[type="reset"]').on('click', function(e) {
        e.preventDefault();
        
        showConfirm('Reset Form', 'Are you sure you want to clear all form data?', 'Yes, Reset').then((result) => {
            if (result.isConfirmed) {
                $(this).closest('form')[0].reset();
                $(this).closest('form').find('.is-valid, .is-invalid').removeClass('is-valid is-invalid');
                $(this).closest('form').find('.invalid-feedback').remove();
                showToast('Form has been reset', 'info');
            }
        });
    });
    
    // Password strength indicator (for registration form)
    $('input[type="password"]').on('input', function() {
        const password = $(this).val();
        
        if (password.length > 0) {
            let strength = 0;
            
            // Length check
            if (password.length >= 8) strength++;
            if (password.length >= 12) strength++;
            
            // Complexity checks
            if (/[a-z]/.test(password)) strength++;
            if (/[A-Z]/.test(password)) strength++;
            if (/[0-9]/.test(password)) strength++;
            if (/[^a-zA-Z0-9]/.test(password)) strength++;
            
            // Remove existing strength indicator
            $(this).next('.password-strength').remove();
            
            // Add strength indicator
            let strengthClass = 'strength-weak';
            let strengthText = 'Weak';
            let strengthColor = 'var(--danger)';
            
            if (strength >= 5) {
                strengthClass = 'strength-strong';
                strengthText = 'Strong';
                strengthColor = 'var(--success)';
            } else if (strength >= 3) {
                strengthClass = 'strength-medium';
                strengthText = 'Medium';
                strengthColor = 'var(--warning)';
            }
            
            const strengthHtml = `
                <div class="password-strength" style="margin-top: 8px;">
                    <div style="height: 4px; background: var(--border); border-radius: 2px; overflow: hidden;">
                        <div style="height: 100%; width: ${strength * 16.66}%; background: ${strengthColor}; transition: all 0.3s;"></div>
                    </div>
                    <div style="font-size: 11px; color: ${strengthColor}; margin-top: 4px;">${strengthText}</div>
                </div>
            `;
            
            $(this).after(strengthHtml);
        } else {
            $(this).next('.password-strength').remove();
        }
    });
});
</script>
@endpush
