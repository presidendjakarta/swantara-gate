@extends('layouts.app')

@section('title', 'Textarea')

@push('styles')
<style>
/* ============================================
   START: Textarea Component Styles
   ============================================ */

/* Textarea Grid */
.textarea-grid {
    display: grid;
    grid-template-columns: repeat(2, 1fr);
    gap: 24px;
    margin-bottom: 24px;
}

.textarea-grid.three-cols {
    grid-template-columns: repeat(3, 1fr);
}

.textarea-grid.full-width {
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

/* Textarea Control */
.textarea-control {
    width: 100%;
    padding: 12px 16px;
    border: 2px solid var(--border-color);
    border-radius: 8px;
    font-size: 14px;
    font-family: inherit;
    line-height: 1.6;
    transition: all 0.2s;
    background: var(--bg-primary);
    color: var(--text-primary);
    resize: vertical;
    overflow-y: auto;
}

.textarea-control:focus {
    outline: none;
    border-color: var(--accent);
    box-shadow: 0 0 0 3px rgba(0, 120, 212, 0.1);
}

.textarea-control::placeholder {
    color: var(--text-tertiary);
}

/* Size Variants */
.textarea-sm {
    min-height: 60px;
}

.textarea-md {
    min-height: 100px;
}

.textarea-lg {
    min-height: 160px;
}

.textarea-xl {
    min-height: 240px;
}

/* Resize Options */
.textarea-resize-none {
    resize: none;
}

.textarea-resize-vertical {
    resize: vertical;
}

.textarea-resize-both {
    resize: both;
}

/* Monospace (Code) */
.textarea-monospace {
    font-family: 'Consolas', 'Monaco', 'Courier New', monospace;
    font-size: 13px;
    line-height: 1.5;
    background: var(--bg-secondary);
    tab-size: 4;
}

/* Validation States */
.textarea-control.is-valid {
    border-color: var(--success);
}

.textarea-control.is-valid:focus {
    box-shadow: 0 0 0 3px rgba(16, 185, 129, 0.1);
}

.textarea-control.is-invalid {
    border-color: var(--danger);
}

.textarea-control.is-invalid:focus {
    box-shadow: 0 0 0 3px rgba(239, 68, 68, 0.1);
}

/* Disabled & Read-only */
.textarea-control:disabled {
    background: var(--bg-secondary);
    color: var(--text-tertiary);
    cursor: not-allowed;
    opacity: 0.6;
}

.textarea-control[readonly] {
    background: var(--bg-secondary);
    color: var(--text-secondary);
    cursor: text;
}

/* With Toolbar */
.textarea-with-toolbar {
    border-top-left-radius: 0;
    border-top-right-radius: 0;
}

/* Character Counter */
.char-counter {
    display: flex;
    justify-content: space-between;
    align-items: center;
    padding: 8px 0;
    font-size: 12px;
    color: var(--text-secondary);
    margin-top: 4px;
}

.char-count {
    font-weight: 600;
    color: var(--text-primary);
}

.char-counter.warning .char-count {
    color: var(--warning);
}

.char-counter.danger .char-count {
    color: var(--danger);
}

.char-counter.warning {
    color: var(--warning);
}

.char-counter.danger {
    color: var(--danger);
}

/* Rich Text Toolbar */
.rich-text-toolbar {
    display: flex;
    align-items: center;
    gap: 4px;
    padding: 8px;
    background: var(--bg-secondary);
    border: 2px solid var(--border-color);
    border-bottom: none;
    border-top-left-radius: 8px;
    border-top-right-radius: 8px;
}

.rich-text-btn {
    width: 32px;
    height: 32px;
    display: flex;
    align-items: center;
    justify-content: center;
    border: 1px solid transparent;
    background: transparent;
    border-radius: 4px;
    cursor: pointer;
    color: var(--text-secondary);
    transition: all 0.2s;
    font-size: 14px;
}

.rich-text-btn:hover {
    background: var(--bg-tertiary);
    border-color: var(--border-color);
    color: var(--text-primary);
}

.rich-text-btn.active {
    background: var(--accent-bg);
    border-color: var(--accent);
    color: var(--accent);
}

.rich-text-divider {
    width: 1px;
    height: 24px;
    background: var(--border-color);
    margin: 0 4px;
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
@media (max-width: 1024px) {
    .textarea-grid.three-cols {
        grid-template-columns: repeat(2, 1fr);
    }
}

@media (max-width: 768px) {
    .textarea-grid,
    .textarea-grid.three-cols {
        grid-template-columns: 1fr;
    }
}

/* ============================================
   END: Textarea Component Styles
   ============================================ */
</style>
@endpush

@section('content')
<div class="page-header">
    <div>
        <h1>Textarea</h1>
        <p>Multi-line text input components dengan berbagai ukuran, resize options, dan character counter</p>
    </div>
    <div style="display: flex; gap: 10px;">
        <button class="btn btn-secondary">
            <i class="fa-solid fa-book"></i>
            <span>Documentation</span>
        </button>
        <button class="btn btn-primary">
            <i class="fa-solid fa-text-width"></i>
            <span>Try Textarea</span>
        </button>
    </div>
</div>

<!-- Info Alert -->
<div class="content-card" style="margin-bottom: 24px; border-left: 4px solid var(--accent);">
    <div class="card-body" style="padding: 16px 20px;">
        <div style="display: flex; gap: 12px; align-items: start;">
            <i class="fa-solid fa-circle-info" style="color: var(--accent); font-size: 20px; margin-top: 2px;"></i>
            <div style="flex: 1;">
                <h4 style="margin-bottom: 4px; font-size: 14px;">What is Textarea?</h4>
                <p style="font-size: 13px; color: var(--text-secondary); margin: 0;">Textarea is a multi-line text input control that allows users to enter longer text content. Commonly used for comments, descriptions, messages, and code. Supports character limits, auto-resize, and rich text formatting.</p>
            </div>
        </div>
    </div>
</div>

<!-- Basic Textareas -->
<div class="ta-section-title">
    <i class="fa-solid fa-text-width"></i>
    Basic Textareas <span class="badge badge-primary">Essential</span>
</div>

<div class="textarea-grid">
    <div class="content-card">
        <div class="card-header">
            <div class="card-header-left">
                <div class="card-icon bg-primary">
                    <i class="fa-solid fa-align-left"></i>
                </div>
                <div>
                    <h3>Default Textarea</h3>
                    <p class="card-subtitle">Basic multi-line input</p>
                </div>
            </div>
        </div>
        <div class="card-body">
            <div class="ta-example">
                <label class="ta-label">
                    Comments
                    <span class="ta-hint">Multi-line</span>
                </label>
                <textarea class="textarea-control textarea-md" placeholder="Enter your comments here..."></textarea>
                <div class="ta-helper">
                    <i class="fa-solid fa-circle-info"></i>
                    Rows: Auto (min-height: 100px)
                </div>
            </div>

            <div class="ta-example">
                <label class="ta-label">
                    Description
                    <span class="ta-hint">With placeholder</span>
                </label>
                <textarea class="textarea-control textarea-md" placeholder="Provide a detailed description of the product or service...">This is a sample description text that demonstrates how textarea can hold multiple lines of content. Users can type freely and the text will wrap automatically.</textarea>
                <div class="ta-helper">
                    <i class="fa-solid fa-circle-info"></i>
                    Pre-filled with content
                </div>
            </div>

            <div class="ta-example">
                <label class="ta-label">
                    Message
                    <span class="ta-hint">Empty</span>
                </label>
                <textarea class="textarea-control textarea-md" placeholder="Type your message here..."></textarea>
                <div class="ta-helper">
                    <i class="fa-solid fa-circle-info"></i>
                    Ready for input
                </div>
            </div>

            <div class="divider"></div>

            <div class="form-group">
                <h4 style="font-size: 14px; margin-bottom: 12px;">Features:</h4>
                <ul class="feature-list">
                    <li>
                        <i class="fa-solid fa-circle-check"></i>
                        <span>Multi-line text input</span>
                    </li>
                    <li>
                        <i class="fa-solid fa-circle-check"></i>
                        <span>Vertical resize by default</span>
                    </li>
                    <li>
                        <i class="fa-solid fa-circle-check"></i>
                        <span>Text wrapping enabled</span>
                    </li>
                    <li>
                        <i class="fa-solid fa-circle-check"></i>
                        <span>Placeholder support</span>
                    </li>
                </ul>
            </div>
        </div>
    </div>

    <div class="content-card">
        <div class="card-header">
            <div class="card-header-left">
                <div class="card-icon bg-success">
                    <i class="fa-solid fa-up-down"></i>
                </div>
                <div>
                    <h3>Size Variants</h3>
                    <p class="card-subtitle">Small, medium, large, XL</p>
                </div>
            </div>
        </div>
        <div class="card-body">
            <div class="ta-example">
                <label class="ta-label">
                    Small Textarea
                    <span class="ta-hint">60px</span>
                </label>
                <textarea class="textarea-control textarea-sm" placeholder="Short note..."></textarea>
                <div class="ta-helper">
                    <i class="fa-solid fa-circle-info"></i>
                    min-height: 60px
                </div>
            </div>

            <div class="ta-example">
                <label class="ta-label">
                    Medium Textarea (Default)
                    <span class="ta-hint">100px</span>
                </label>
                <textarea class="textarea-control textarea-md" placeholder="Standard input..."></textarea>
                <div class="ta-helper">
                    <i class="fa-solid fa-circle-info"></i>
                    min-height: 100px
                </div>
            </div>

            <div class="ta-example">
                <label class="ta-label">
                    Large Textarea
                    <span class="ta-hint">160px</span>
                </label>
                <textarea class="textarea-control textarea-lg" placeholder="Longer content..."></textarea>
                <div class="ta-helper">
                    <i class="fa-solid fa-circle-info"></i>
                    min-height: 160px
                </div>
            </div>

            <div class="ta-example">
                <label class="ta-label">
                    Extra Large Textarea
                    <span class="ta-hint">240px</span>
                </label>
                <textarea class="textarea-control textarea-xl" placeholder="Extended content area..."></textarea>
                <div class="ta-helper">
                    <i class="fa-solid fa-circle-info"></i>
                    min-height: 240px
                </div>
            </div>

            <div class="divider"></div>

            <div class="helper-text">
                <i class="fa-solid fa-circle-info"></i>
                Choose size based on expected content length
            </div>
        </div>
    </div>
</div>

<!-- Resize Options -->
<div class="ta-section-title">
    <i class="fa-solid fa-arrows-up-down-left-right"></i>
    Resize Options
</div>

<div class="textarea-grid three-cols">
    <div class="content-card">
        <div class="card-header">
            <div class="card-header-left">
                <div class="card-icon bg-info">
                    <i class="fa-solid fa-ban"></i>
                </div>
                <div>
                    <h3>No Resize</h3>
                    <p class="card-subtitle">Fixed size textarea</p>
                </div>
            </div>
        </div>
        <div class="card-body">
            <div class="ta-example">
                <label class="ta-label">Fixed Size</label>
                <textarea class="textarea-control textarea-md textarea-resize-none" placeholder="Cannot be resized..."></textarea>
                <div class="ta-helper">
                    <i class="fa-solid fa-circle-info"></i>
                    resize: none
                </div>
            </div>

            <div class="divider"></div>

            <div class="helper-text">
                <i class="fa-solid fa-circle-info"></i>
                No resize handle in bottom-right corner
            </div>
        </div>
    </div>

    <div class="content-card">
        <div class="card-header">
            <div class="card-header-left">
                <div class="card-icon bg-primary">
                    <i class="fa-solid fa-arrows-up-down"></i>
                </div>
                <div>
                    <h3>Vertical Resize</h3>
                    <p class="card-subtitle">Height only (Default)</p>
                </div>
            </div>
        </div>
        <div class="card-body">
            <div class="ta-example">
                <label class="ta-label">Vertical Only</label>
                <textarea class="textarea-control textarea-md textarea-resize-vertical" placeholder="Drag bottom-right to resize vertically..."></textarea>
                <div class="ta-helper">
                    <i class="fa-solid fa-circle-info"></i>
                    resize: vertical
                </div>
            </div>

            <div class="divider"></div>

            <div class="helper-text">
                <i class="fa-solid fa-circle-info"></i>
                Most common resize mode
            </div>
        </div>
    </div>

    <div class="content-card">
        <div class="card-header">
            <div class="card-header-left">
                <div class="card-icon bg-success">
                    <i class="fa-solid fa-arrows-left-right"></i>
                </div>
                <div>
                    <h3>Both Directions</h3>
                    <p class="card-subtitle">Width and height</p>
                </div>
            </div>
        </div>
        <div class="card-body">
            <div class="ta-example">
                <label class="ta-label">Both Directions</label>
                <textarea class="textarea-control textarea-md textarea-resize-both" placeholder="Drag to resize in any direction..."></textarea>
                <div class="ta-helper">
                    <i class="fa-solid fa-circle-info"></i>
                    resize: both
                </div>
            </div>

            <div class="divider"></div>

            <div class="helper-text">
                <i class="fa-solid fa-circle-info"></i>
                Full resize control
            </div>
        </div>
    </div>
</div>

<!-- Character Counter -->
<div class="ta-section-title">
    <i class="fa-solid fa-hashtag"></i>
    Character Counter <span class="badge badge-success">Popular</span>
</div>

<div class="textarea-grid">
    <div class="content-card">
        <div class="card-header">
            <div class="card-header-left">
                <div class="card-icon bg-primary">
                    <i class="fa-solid fa-text-width"></i>
                </div>
                <div>
                    <h3>With Character Limit</h3>
                    <p class="card-subtitle">Count characters as user types</p>
                </div>
            </div>
        </div>
        <div class="card-body">
            <div class="ta-example">
                <label class="ta-label">
                    Bio (Max 100 characters)
                    <span class="ta-hint">100 chars</span>
                </label>
                <textarea class="textarea-control textarea-md" placeholder="Tell us about yourself..." maxlength="100">Hi, I'm a web developer passionate about creating beautiful and functional websites.</textarea>
                <div class="char-counter">
                    <span>Characters used</span>
                    <span class="char-count">85 / 100</span>
                </div>
            </div>

            <div class="ta-example">
                <label class="ta-label">
                    Tweet (Max 280 characters)
                    <span class="ta-hint">280 chars</span>
                </label>
                <textarea class="textarea-control textarea-md" placeholder="What's happening?" maxlength="280">Just launched my new portfolio website! Check it out 🚀 #webdev #portfolio</textarea>
                <div class="char-counter">
                    <span>Characters used</span>
                    <span class="char-count">77 / 280</span>
                </div>
            </div>

            <div class="ta-example">
                <label class="ta-label">
                    Comment (Max 50 characters)
                    <span class="ta-hint">50 chars - Warning</span>
                </label>
                <textarea class="textarea-control textarea-sm" placeholder="Leave a comment..." maxlength="50">This is getting close to the limit</textarea>
                <div class="char-counter warning">
                    <span>⚠️ Approaching limit</span>
                    <span class="char-count">42 / 50</span>
                </div>
            </div>

            <div class="ta-example">
                <label class="ta-label">
                    Title (Max 20 characters)
                    <span class="ta-hint">20 chars - Danger</span>
                </label>
                <textarea class="textarea-control textarea-sm" placeholder="Enter title..." maxlength="20">Almost there!!!</textarea>
                <div class="char-counter danger">
                    <span>🚨 Limit reached!</span>
                    <span class="char-count">18 / 20</span>
                </div>
            </div>

            <div class="divider"></div>

            <div class="helper-text">
                <i class="fa-solid fa-circle-info"></i>
                Counter changes color as limit approaches
            </div>
        </div>
    </div>

    <div class="content-card">
        <div class="card-header">
            <div class="card-header-left">
                <div class="card-icon bg-success">
                    <i class="fa-solid fa-calculator"></i>
                </div>
                <div>
                    <h3>Word & Line Count</h3>
                    <p class="card-subtitle">Advanced text statistics</p>
                </div>
            </div>
        </div>
        <div class="card-body">
            <div class="ta-example">
                <label class="ta-label">
                    Article Content
                    <span class="ta-hint">Word count</span>
                </label>
                <textarea class="textarea-control textarea-lg" placeholder="Write your article...">The quick brown fox jumps over the lazy dog. This sentence contains every letter of the alphabet at least once.</textarea>
                <div class="char-counter">
                    <span>Words: 20 | Lines: 2</span>
                    <span class="char-count">126 chars</span>
                </div>
            </div>

            <div class="ta-example">
                <label class="ta-label">
                    Poem
                    <span class="ta-hint">Line count</span>
                </label>
                <textarea class="textarea-control textarea-lg" placeholder="Write a poem...">Roses are red,
Violets are blue,
Sugar is sweet,
And so are you.</textarea>
                <div class="char-counter">
                    <span>Words: 13 | Lines: 4</span>
                    <span class="char-count">70 chars</span>
                </div>
            </div>

            <div class="divider"></div>

            <div class="form-group">
                <h4 style="font-size: 14px; margin-bottom: 12px;">Counter Features:</h4>
                <ul class="feature-list">
                    <li>
                        <i class="fa-solid fa-circle-check"></i>
                        <span>Character counting</span>
                    </li>
                    <li>
                        <i class="fa-solid fa-circle-check"></i>
                        <span>Word counting</span>
                    </li>
                    <li>
                        <i class="fa-solid fa-circle-check"></i>
                        <span>Line counting</span>
                    </li>
                    <li>
                        <i class="fa-solid fa-circle-check"></i>
                        <span>Color warnings</span>
                    </li>
                </ul>
            </div>
        </div>
    </div>
</div>

<!-- States & Special -->
<div class="ta-section-title">
    <i class="fa-solid fa-circle-check"></i>
    States & Special Textareas
</div>

<div class="textarea-grid">
    <div class="content-card">
        <div class="card-header">
            <div class="card-header-left">
                <div class="card-icon bg-success">
                    <i class="fa-solid fa-check"></i>
                </div>
                <div>
                    <h3>Valid & Invalid</h3>
                    <p class="card-subtitle">Validation states</p>
                </div>
            </div>
        </div>
        <div class="card-body">
            <div class="ta-example">
                <label class="ta-label">Valid State</label>
                <textarea class="textarea-control textarea-md is-valid" placeholder="Enter description...">This is a valid entry with proper content.</textarea>
                <div class="ta-helper" style="color: var(--success);">
                    <i class="fa-solid fa-circle-check"></i>
                    <span>Looks good!</span>
                </div>
            </div>

            <div class="ta-example">
                <label class="ta-label">Invalid State</label>
                <textarea class="textarea-control textarea-md is-invalid" placeholder="Enter description..."></textarea>
                <div class="ta-helper" style="color: var(--danger);">
                    <i class="fa-solid fa-circle-exclamation"></i>
                    <span>Description is required</span>
                </div>
            </div>

            <div class="divider"></div>

            <div class="helper-text">
                <i class="fa-solid fa-circle-info"></i>
                Green/Red borders indicate validation state
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
                    <h3>Disabled & Read-only</h3>
                    <p class="card-subtitle">Non-editable textareas</p>
                </div>
            </div>
        </div>
        <div class="card-body">
            <div class="ta-example">
                <label class="ta-label">Disabled</label>
                <textarea class="textarea-control textarea-md" disabled>This textarea is disabled and cannot be edited. The content is grayed out.</textarea>
                <div class="ta-helper">
                    <i class="fa-solid fa-circle-info"></i>
                    Cannot be focused or edited
                </div>
            </div>

            <div class="ta-example">
                <label class="ta-label">Read-only</label>
                <textarea class="textarea-control textarea-md" readonly>This textarea is read-only. Users can select and copy the text but cannot modify it.</textarea>
                <div class="ta-helper">
                    <i class="fa-solid fa-circle-info"></i>
                    Can be focused and selected
                </div>
            </div>

            <div class="divider"></div>

            <div class="helper-text">
                <i class="fa-solid fa-circle-info"></i>
                Both states have grayed background
            </div>
        </div>
    </div>
</div>

<!-- Code & Rich Text -->
<div class="ta-section-title">
    <i class="fa-solid fa-code"></i>
    Code & Rich Text
</div>

<div class="textarea-grid">
    <div class="content-card">
        <div class="card-header">
            <div class="card-header-left">
                <div class="card-icon bg-info">
                    <i class="fa-solid fa-code"></i>
                </div>
                <div>
                    <h3>Monospace (Code)</h3>
                    <p class="card-subtitle">Code editor style</p>
                </div>
            </div>
        </div>
        <div class="card-body">
            <div class="ta-example">
                <label class="ta-label">
                    HTML Code
                    <span class="ta-hint">Monospace font</span>
                </label>
                <textarea class="textarea-control textarea-lg textarea-monospace textarea-resize-none" spellcheck="false"><!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <title>My Page</title>
</head>
<body>
    <h1>Hello World</h1>
</body>
</html></textarea>
                <div class="ta-helper">
                    <i class="fa-solid fa-circle-info"></i>
                    Fixed-width font, spellcheck disabled
                </div>
            </div>

            <div class="ta-example">
                <label class="ta-label">
                    JavaScript Code
                    <span class="ta-hint">Code editor</span>
                </label>
                <textarea class="textarea-control textarea-md textarea-monospace textarea-resize-none" spellcheck="false">function greet(name) {
    return `Hello, ${name}!`;
}

console.log(greet('World'));</textarea>
            </div>

            <div class="divider"></div>

            <div class="helper-text">
                <i class="fa-solid fa-circle-info"></i>
                Monospace font aligns characters perfectly
            </div>
        </div>
    </div>

    <div class="content-card">
        <div class="card-header">
            <div class="card-header-left">
                <div class="card-icon bg-success">
                    <i class="fa-solid fa-wand-magic-sparkles"></i>
                </div>
                <div>
                    <h3>Rich Text Editor (Visual)</h3>
                    <p class="card-subtitle">With formatting toolbar</p>
                </div>
            </div>
        </div>
        <div class="card-body">
            <div class="ta-example">
                <label class="ta-label">
                    Content Editor
                    <span class="ta-hint">With toolbar</span>
                </label>
                <div class="rich-text-toolbar">
                    <button class="rich-text-btn active" title="Bold"><i class="fa-solid fa-bold"></i></button>
                    <button class="rich-text-btn" title="Italic"><i class="fa-solid fa-italic"></i></button>
                    <button class="rich-text-btn" title="Underline"><i class="fa-solid fa-underline"></i></button>
                    <div class="rich-text-divider"></div>
                    <button class="rich-text-btn" title="Align Left"><i class="fa-solid fa-align-left"></i></button>
                    <button class="rich-text-btn active" title="Align Center"><i class="fa-solid fa-align-center"></i></button>
                    <button class="rich-text-btn" title="Align Right"><i class="fa-solid fa-align-right"></i></button>
                    <div class="rich-text-divider"></div>
                    <button class="rich-text-btn" title="List"><i class="fa-solid fa-list"></i></button>
                    <button class="rich-text-btn" title="Link"><i class="fa-solid fa-link"></i></button>
                </div>
                <textarea class="textarea-control textarea-lg textarea-with-toolbar" placeholder="Start writing your content..."><h2>Welcome to Our Website</h2>

<p>This is a <strong>rich text editor</strong> example with formatting options. You can make text <em>italic</em> or <u>underlined</u>.</p>

<ul>
    <li>Feature one</li>
    <li>Feature two</li>
    <li>Feature three</li>
</ul></textarea>
            </div>

            <div class="divider"></div>

            <div class="helper-text">
                <i class="fa-solid fa-circle-info"></i>
                Toolbar shown above textarea (visual only)
            </div>
        </div>
    </div>
</div>

<!-- Contact Form Example -->
<div class="ta-section-title">
    <i class="fa-solid fa-clipboard-list"></i>
    Contact Form Example
</div>

<div class="textarea-grid full-width">
    <div class="content-card">
        <div class="card-header">
            <div class="card-header-left">
                <div class="card-icon bg-primary">
                    <i class="fa-solid fa-envelope"></i>
                </div>
                <div>
                    <h3>Contact Us Form</h3>
                    <p class="card-subtitle">Complete form with textarea</p>
                </div>
            </div>
        </div>
        <div class="card-body">
            <form style="max-width: 800px;">
                <div style="display: grid; grid-template-columns: 1fr 1fr; gap: 20px;">
                    <div class="ta-example">
                        <label class="ta-label">Full Name <span style="color: var(--danger);">*</span></label>
                        <input type="text" class="textarea-control" style="min-height: auto;" placeholder="John Doe" required>
                    </div>

                    <div class="ta-example">
                        <label class="ta-label">Email <span style="color: var(--danger);">*</span></label>
                        <input type="email" class="textarea-control" style="min-height: auto;" placeholder="john@example.com" required>
                    </div>
                </div>

                <div class="ta-example">
                    <label class="ta-label">Subject <span style="color: var(--danger);">*</span></label>
                    <input type="text" class="textarea-control" style="min-height: auto;" placeholder="What is this about?" required>
                </div>

                <div class="ta-example">
                    <label class="ta-label">Message <span style="color: var(--danger);">*</span></label>
                    <textarea class="textarea-control textarea-lg" placeholder="Write your message in detail..." maxlength="500" required></textarea>
                    <div class="char-counter">
                        <span>Minimum 50 characters</span>
                        <span class="char-count">0 / 500</span>
                    </div>
                </div>

                <div style="display: flex; gap: 12px; margin-top: 24px;">
                    <button type="submit" class="btn btn-primary" style="flex: 1;">
                        <i class="fa-solid fa-paper-plane"></i>
                        Send Message
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
<div class="textarea-grid full-width">
    <div class="content-card">
        <div class="card-header">
            <div class="card-header-left">
                <div class="card-icon bg-success">
                    <i class="fa-solid fa-code"></i>
                </div>
                <div>
                    <h3>Implementation Guide</h3>
                    <p class="card-subtitle">How to implement textarea features with JavaScript</p>
                </div>
            </div>
        </div>
        <div class="card-body">
            <div class="code-block">
                <div style="color: var(--text-tertiary); margin-bottom: 8px;">1. Basic HTML Textarea:</div>
                <code style="color: var(--accent);">
                    &lt;textarea class="textarea-control textarea-md" <br>
                    &nbsp;&nbsp;placeholder="Enter text..."<br>
                    &nbsp;&nbsp;maxlength="500"<br>
                    &nbsp;&nbsp;rows="5"&gt;<br>
                    &lt;/textarea&gt;
                </code>
            </div>

            <div class="code-block">
                <div style="color: var(--text-tertiary); margin-bottom: 8px; margin-top: 16px;">2. Character Counter:</div>
                <code style="color: var(--success);">
                    const textarea = document.querySelector('#myTextarea');<br>
                    const counter = document.querySelector('#charCounter');<br>
                    <br>
                    textarea.addEventListener('input', function() {<br>
                    &nbsp;&nbsp;const length = this.value.length;<br>
                    &nbsp;&nbsp;const max = this.maxLength;<br>
                    &nbsp;&nbsp;counter.textContent = `${length} / ${max}`;<br>
                    });
                </code>
            </div>

            <div class="code-block">
                <div style="color: var(--text-tertiary); margin-bottom: 8px; margin-top: 16px;">3. Auto-Resize:</div>
                <code style="color: var(--warning);">
                    textarea.addEventListener('input', function() {<br>
                    &nbsp;&nbsp;this.style.height = 'auto';<br>
                    &nbsp;&nbsp;this.style.height = this.scrollHeight + 'px';<br>
                    });
                </code>
            </div>

            <div class="code-block">
                <div style="color: var(--text-tertiary); margin-bottom: 8px; margin-top: 16px;">4. Word Count:</div>
                <code style="color: var(--info);">
                    textarea.addEventListener('input', function() {<br>
                    &nbsp;&nbsp;const text = this.value.trim();<br>
                    &nbsp;&nbsp;const words = text ? text.split(/\s+/).length : 0;<br>
                    &nbsp;&nbsp;const lines = text.split('\n').length;<br>
                    &nbsp;&nbsp;wordCounter.textContent = `Words: ${words}`;<br>
                    });
                </code>
            </div>

            <div class="divider"></div>

            <div class="form-group">
                <h4 style="font-size: 14px; margin-bottom: 12px;">Popular Libraries:</h4>
                <ul class="feature-list">
                    <li>
                        <i class="fa-solid fa-circle-check"></i>
                        <span><strong>CKEditor</strong> - Full-featured WYSIWYG</span>
                    </li>
                    <li>
                        <i class="fa-solid fa-circle-check"></i>
                        <span><strong>TinyMCE</strong> - Rich text editor</span>
                    </li>
                    <li>
                        <i class="fa-solid fa-circle-check"></i>
                        <span><strong>Quill</strong> - Modern, lightweight</span>
                    </li>
                    <li>
                        <i class="fa-solid fa-circle-check"></i>
                        <span><strong>CodeMirror</strong> - Code editor</span>
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
    // Character counter functionality
    $('textarea[maxlength]').each(function() {
        const $textarea = $(this);
        const $example = $textarea.closest('.ta-example');
        const $counter = $example.find('.char-counter');
        const maxlength = parseInt($textarea.attr('maxlength'));
        
        function updateCounter() {
            const length = $textarea.val().length;
            const remaining = maxlength - length;
            
            // Update counter text
            $counter.find('.char-count').text(length + ' / ' + maxlength);
            
            // Update counter color based on usage
            $counter.removeClass('warning danger');
            
            if (remaining <= 5) {
                $counter.addClass('danger');
                $counter.find('span:first').html('🚨 Limit reached!');
            } else if (remaining <= maxlength * 0.2) {
                $counter.addClass('warning');
                $counter.find('span:first').html('⚠️ Approaching limit');
            } else {
                $counter.find('span:first').text('Characters used');
            }
        }
        
        // Initial update
        updateCounter();
        
        // Update on input
        $textarea.on('input', updateCounter);
    });
    
    // Word and line counter
    $('.ta-example').each(function() {
        const $example = $(this);
        const $textarea = $example.find('textarea');
        const $counter = $example.find('.char-counter');
        
        // Check if this is a word/line counter (has "Words:" in the counter)
        if ($counter.length && $counter.find('span:first').text().includes('Words:')) {
            function updateWordLineCount() {
                const text = $textarea.val().trim();
                const chars = $textarea.val().length;
                const words = text ? text.split(/\s+/).length : 0;
                const lines = text ? text.split('\n').length : 0;
                
                $counter.find('span:first').text(`Words: ${words} | Lines: ${lines}`);
                $counter.find('.char-count').text(chars + ' chars');
            }
            
            // Initial update
            updateWordLineCount();
            
            // Update on input
            $textarea.on('input', updateWordLineCount);
        }
    });
    
    // Auto-resize textarea (optional enhancement)
    $('textarea').on('input', function() {
        // Only auto-resize if it has a specific class or data attribute
        if ($(this).hasClass('textarea-auto-resize')) {
            $(this).css('height', 'auto');
            $(this).css('height', $(this)[0].scrollHeight + 'px');
        }
    });
    
    // Tab key support for code textareas
    $('textarea.textarea-monospace').on('keydown', function(e) {
        if (e.key === 'Tab') {
            e.preventDefault();
            
            const start = this.selectionStart;
            const end = this.selectionEnd;
            const value = $(this).val();
            
            // Insert 4 spaces
            $(this).val(value.substring(0, start) + '    ' + value.substring(end));
            
            // Move cursor
            this.selectionStart = this.selectionEnd = start + 4;
            
            // Trigger input event
            $(this).trigger('input');
        }
    });
    
    // Rich text toolbar buttons (visual only)
    $('.rich-text-btn').on('click', function(e) {
        e.preventDefault();
        
        // Toggle active state
        $(this).toggleClass('active');
        
        // Show feedback
        const title = $(this).attr('title');
        if (title) {
            Swal.fire({
                toast: true,
                position: 'top-end',
                icon: 'info',
                title: title + ' toggled',
                showConfirmButton: false,
                timer: 1500
            });
        }
    });
    
    // Contact form submission
    $('form').on('submit', function(e) {
        e.preventDefault();
        
        const $form = $(this);
        const formData = {};
        
        // Collect form data
        $form.find('input[type="text"], input[type="email"], textarea').each(function() {
            const $input = $(this);
            const label = $input.closest('.ta-example').find('.ta-label').text().replace(' *', '').trim();
            const value = $input.val().trim();
            
            if (value) {
                formData[label] = value;
            }
        });
        
        // Validate required fields
        const requiredFields = ['Full Name', 'Email', 'Subject', 'Message'];
        const missingFields = requiredFields.filter(field => !formData[field]);
        
        if (missingFields.length > 0) {
            Swal.fire({
                toast: true,
                position: 'top-end',
                icon: 'warning',
                title: 'Please fill in: ' + missingFields.join(', '),
                showConfirmButton: false,
                timer: 3000
            });
            
            // Focus first missing field
            const firstMissing = missingFields[0];
            $form.find('.ta-label').each(function() {
                if ($(this).text().includes(firstMissing)) {
                    $(this).closest('.ta-example').find('input, textarea').first().focus();
                    return false;
                }
            });
            return;
        }
        
        // Validate message length (min 50 chars)
        if (formData['Message'] && formData['Message'].length < 50) {
            Swal.fire({
                toast: true,
                position: 'top-end',
                icon: 'warning',
                title: 'Message must be at least 50 characters',
                showConfirmButton: false,
                timer: 2500
            });
            $form.find('textarea').focus();
            return;
        }
        
        // Show success message
        Swal.fire({
            icon: 'success',
            title: 'Message Sent!',
            html: `
                <div style="text-align: left; font-size: 13px;">
                    <p><strong>From:</strong> ${formData['Full Name']}</p>
                    <p><strong>Email:</strong> ${formData['Email']}</p>
                    <p><strong>Subject:</strong> ${formData['Subject']}</p>
                    <p><strong>Message:</strong> ${formData['Message'].substring(0, 100)}${formData['Message'].length > 100 ? '...' : ''}</p>
                </div>
            `,
            confirmButtonText: 'OK',
            confirmButtonColor: 'var(--accent)'
        });
        
        console.log('Form Data:', formData);
    });
    
    // Form reset
    $('form').on('reset', function(e) {
        e.preventDefault();
        
        // Clear all inputs
        $(this).find('input, textarea').val('');
        
        // Reset character counters
        $(this).find('.char-counter').each(function() {
            const $counter = $(this);
            $counter.removeClass('warning danger');
            $counter.find('.char-count').text('0 / ' + ($counter.closest('.ta-example').find('textarea').attr('maxlength') || '500'));
            
            if ($counter.find('span:first').text().includes('Minimum')) {
                $counter.find('span:first').text('Minimum 50 characters');
            } else {
                $counter.find('span:first').text('Characters used');
            }
        });
        
        // Reset word/line counters
        $(this).find('textarea').trigger('input');
        
        Swal.fire({
            toast: true,
            position: 'top-end',
            icon: 'info',
            title: 'Form reset',
            showConfirmButton: false,
            timer: 1500
        });
    });
    
    // Textarea validation on blur
    $('textarea[required]').on('blur', function() {
        const $textarea = $(this);
        const $example = $textarea.closest('.ta-example');
        
        if ($textarea.val().trim() === '') {
            $textarea.removeClass('is-valid').addClass('is-invalid');
            
            // Add error message if not exists
            if ($example.find('.ta-helper.text-error').length === 0) {
                $example.append('<div class="ta-helper text-error" style="color: var(--danger);"><i class="fa-solid fa-circle-exclamation"></i><span>This field is required</span></div>');
            }
        } else {
            $textarea.removeClass('is-invalid').addClass('is-valid');
            
            // Remove error message
            $example.find('.ta-helper.text-error').remove();
        }
    });
    
    // Remove validation state on input
    $('textarea').on('input', function() {
        $(this).removeClass('is-invalid is-valid');
        $(this).closest('.ta-example').find('.ta-helper.text-error').remove();
    });
});
</script>
@endpush
