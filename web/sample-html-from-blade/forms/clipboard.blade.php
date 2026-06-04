@extends('layouts.app')

@section('title', 'Clipboard')

@push('styles')
@endpush

@section('content')
<div class="page-header">
    <div>
        <h1>Clipboard</h1>
        <p>Copy-to-clipboard functionality dengan visual feedback dan toast notifications</p>
    </div>
    <div style="display: flex; gap: 10px;">
        <button class="btn btn-secondary">
            <i class="fa-solid fa-book"></i>
            <span>Documentation</span>
        </button>
        <button class="btn btn-primary">
            <i class="fa-solid fa-copy"></i>
            <span>Try Clipboard</span>
        </button>
    </div>
</div>

<!-- Info Alert -->
<div class="content-card" style="margin-bottom: 24px; border-left: 4px solid var(--accent);">
    <div class="card-body" style="padding: 16px 20px;">
        <div style="display: flex; gap: 12px; align-items: start;">
            <i class="fa-solid fa-circle-info" style="color: var(--accent); font-size: 20px; margin-top: 2px;"></i>
            <div style="flex: 1;">
                <h4 style="margin-bottom: 4px; font-size: 14px;">What is Clipboard?</h4>
                <p style="font-size: 13px; color: var(--text-secondary); margin: 0;">Clipboard functionality allows users to copy text, URLs, code snippets, and other content with a single click. Uses the modern <strong>Clipboard API</strong> with fallbacks for older browsers. Provides instant visual feedback to confirm successful copy.</p>
            </div>
        </div>
    </div>
</div>

<!-- Basic Clipboard -->
<div class="cb-section-title">
    <i class="fa-solid fa-copy"></i>
    Basic Clipboard <span class="badge badge-primary">Essential</span>
</div>

<div class="clipboard-grid">
    <div class="content-card">
        <div class="card-header">
            <div class="card-header-left">
                <div class="card-icon bg-primary">
                    <i class="fa-solid fa-copy"></i>
                </div>
                <div>
                    <h3>Input with Copy Button</h3>
                    <p class="card-subtitle">Inline copy button</p>
                </div>
            </div>
        </div>
        <div class="card-body">
            <div class="cb-example">
                <label class="cb-label">
                    Website URL
                    <span class="cb-hint">Click to copy</span>
                </label>
                <div class="clipboard-group">
                    <input type="text" class="clipboard-input" value="https://example.com" readonly>
                    <button class="clipboard-btn" onclick="copyToClipboard(this, 'https://example.com')">
                        <i class="fa-solid fa-copy"></i>
                        <span>Copy</span>
                    </button>
                </div>
                <div class="cb-helper">
                    <i class="fa-solid fa-circle-info"></i>
                    Click button to copy URL
                </div>
            </div>

            <div class="cb-example">
                <label class="cb-label">
                    API Key
                    <span class="cb-hint">Secret key</span>
                </label>
                <div class="clipboard-group">
                    <input type="text" class="clipboard-input" value="sk-1234567890abcdef" readonly>
                    <button class="clipboard-btn" onclick="copyToClipboard(this, 'sk-1234567890abcdef')">
                        <i class="fa-solid fa-copy"></i>
                        <span>Copy</span>
                    </button>
                </div>
                <div class="cb-helper">
                    <i class="fa-solid fa-circle-info"></i>
                    Secure API key copy
                </div>
            </div>

            <div class="cb-example">
                <label class="cb-label">
                    Email Address
                    <span class="cb-hint">Contact email</span>
                </label>
                <div class="clipboard-group">
                    <input type="text" class="clipboard-input" value="support@example.com" readonly>
                    <button class="clipboard-btn" onclick="copyToClipboard(this, 'support@example.com')">
                        <i class="fa-solid fa-copy"></i>
                        <span>Copy</span>
                    </button>
                </div>
            </div>

            <div class="divider"></div>

            <div class="form-group">
                <h4 style="font-size: 14px; margin-bottom: 12px;">Features:</h4>
                <ul class="feature-list">
                    <li>
                        <i class="fa-solid fa-circle-check"></i>
                        <span>One-click copy</span>
                    </li>
                    <li>
                        <i class="fa-solid fa-circle-check"></i>
                        <span>Visual feedback</span>
                    </li>
                    <li>
                        <i class="fa-solid fa-circle-check"></i>
                        <span>Read-only input</span>
                    </li>
                    <li>
                        <i class="fa-solid fa-circle-check"></i>
                        <span>Button state change</span>
                    </li>
                </ul>
            </div>
        </div>
    </div>

    <div class="content-card">
        <div class="card-header">
            <div class="card-header-left">
                <div class="card-icon bg-success">
                    <i class="fa-solid fa-clipboard"></i>
                </div>
                <div>
                    <h3>Copy Button Variants</h3>
                    <p class="card-subtitle">Different button styles</p>
                </div>
            </div>
        </div>
        <div class="card-body">
            <div class="cb-example">
                <label class="cb-label">Standalone Copy Button</label>
                <div class="copy-area">
                    <code id="code1">npm install package-name</code>
                </div>
                <button class="copy-btn" onclick="copyToClipboard(this, 'npm install package-name')">
                    <i class="fa-solid fa-copy"></i>
                    <span>Copy Command</span>
                </button>
            </div>

            <div class="cb-example">
                <label class="cb-label">Copy with Icon</label>
                <div class="copy-area">
                    <div style="font-size: 16px; font-weight: 600;">📧 contact@company.com</div>
                </div>
                <button class="copy-btn" onclick="copyToClipboard(this, 'contact@company.com')">
                    <i class="fa-solid fa-clipboard"></i>
                    <span>Copy Email</span>
                </button>
            </div>

            <div class="cb-example">
                <label class="cb-label">Small Copy Button</label>
                <div class="copy-area">
                    <span style="font-family: monospace;">$ git clone https://github.com/repo.git</span>
                </div>
                <button class="copy-btn" style="padding: 6px 10px; font-size: 12px;" onclick="copyToClipboard(this, 'git clone https://github.com/repo.git')">
                    <i class="fa-solid fa-copy"></i>
                    <span>Copy</span>
                </button>
            </div>

            <div class="divider"></div>

            <div class="helper-text">
                <i class="fa-solid fa-circle-info"></i>
                Buttons change to green when copied
            </div>
        </div>
    </div>
</div>

<!-- Code Blocks -->
<div class="cb-section-title">
    <i class="fa-solid fa-code"></i>
    Code Blocks with Copy <span class="badge badge-success">Popular</span>
</div>

<div class="clipboard-grid">
    <div class="content-card">
        <div class="card-header">
            <div class="card-header-left">
                <div class="card-icon bg-info">
                    <i class="fa-brands fa-html5"></i>
                </div>
                <div>
                    <h3>HTML Code</h3>
                    <p class="card-subtitle">Copy HTML snippets</p>
                </div>
            </div>
        </div>
        <div class="card-body">
            <div class="cb-example">
                <label class="cb-label">HTML Template</label>
                <div class="code-block-clipboard">
                    <button class="copy-btn-absolute" onclick="copyCode(this)">
                        <i class="fa-solid fa-copy"></i>
                        <span>Copy</span>
                    </button>
                    <code>&lt;!DOCTYPE html&gt;
&lt;html lang="en"&gt;
&lt;head&gt;
    &lt;meta charset="UTF-8"&gt;
    &lt;title&gt;My Page&lt;/title&gt;
&lt;/head&gt;
&lt;body&gt;
    &lt;h1&gt;Hello World&lt;/h1&gt;
&lt;/body&gt;
&lt;/html&gt;</code>
                </div>
                <div class="cb-helper">
                    <i class="fa-solid fa-circle-info"></i>
                    Hover to see copy button
                </div>
            </div>

            <div class="cb-example">
                <label class="cb-label">Meta Tags</label>
                <div class="code-block-clipboard">
                    <button class="copy-btn-absolute" onclick="copyCode(this)">
                        <i class="fa-solid fa-copy"></i>
                        <span>Copy</span>
                    </button>
                    <code>&lt;meta name="viewport" content="width=device-width, initial-scale=1.0"&gt;
&lt;meta name="description" content="Page description"&gt;
&lt;meta name="keywords" content="HTML, CSS, JavaScript"&gt;</code>
                </div>
            </div>

            <div class="divider"></div>

            <div class="helper-text">
                <i class="fa-solid fa-circle-info"></i>
                Code blocks have copy button in top-right corner
            </div>
        </div>
    </div>

    <div class="content-card">
        <div class="card-header">
            <div class="card-header-left">
                <div class="card-icon bg-warning">
                    <i class="fa-brands fa-js"></i>
                </div>
                <div>
                    <h3>JavaScript & CSS</h3>
                    <p class="card-subtitle">Copy code snippets</p>
                </div>
            </div>
        </div>
        <div class="card-body">
            <div class="cb-example">
                <label class="cb-label">JavaScript Function</label>
                <div class="code-block-clipboard">
                    <button class="copy-btn-absolute" onclick="copyCode(this)">
                        <i class="fa-solid fa-copy"></i>
                        <span>Copy</span>
                    </button>
                    <code>function copyToClipboard(text) {
    navigator.clipboard.writeText(text)
        .then(() => {
            console.log('Copied!');
        })
        .catch(err => {
            console.error('Failed to copy');
        });
}</code>
                </div>
            </div>

            <div class="cb-example">
                <label class="cb-label">CSS Styles</label>
                <div class="code-block-clipboard">
                    <button class="copy-btn-absolute" onclick="copyCode(this)">
                        <i class="fa-solid fa-copy"></i>
                        <span>Copy</span>
                    </button>
                    <code>.container {
    display: grid;
    grid-template-columns: repeat(2, 1fr);
    gap: 24px;
    padding: 20px;
}</code>
                </div>
            </div>

            <div class="divider"></div>

            <div class="helper-text">
                <i class="fa-solid fa-circle-info"></i>
                Monospace font preserves formatting
            </div>
        </div>
    </div>
</div>

<!-- Special Content -->
<div class="cb-section-title">
    <i class="fa-solid fa-wand-magic-sparkles"></i>
    Special Content
</div>

<div class="clipboard-grid">
    <div class="content-card">
        <div class="card-header">
            <div class="card-header-left">
                <div class="card-icon bg-primary">
                    <i class="fa-solid fa-wifi"></i>
                </div>
                <div>
                    <h3>WiFi Password</h3>
                    <p class="card-subtitle">Quick password sharing</p>
                </div>
            </div>
        </div>
        <div class="card-body">
            <div class="cb-example">
                <label class="cb-label">
                    Network Name
                    <span class="cb-hint">SSID</span>
                </label>
                <div class="clipboard-group">
                    <input type="text" class="clipboard-input" value="MyWiFiNetwork_5G" readonly>
                    <button class="clipboard-btn" onclick="copyToClipboard(this, 'MyWiFiNetwork_5G')">
                        <i class="fa-solid fa-copy"></i>
                        <span>Copy</span>
                    </button>
                </div>
            </div>

            <div class="cb-example">
                <label class="cb-label">
                    Password
                    <span class="cb-hint">Secure</span>
                </label>
                <div class="clipboard-group">
                    <input type="password" class="clipboard-input" value="SecurePass123!" readonly>
                    <button class="clipboard-btn" onclick="copyToClipboard(this, 'SecurePass123!')">
                        <i class="fa-solid fa-copy"></i>
                        <span>Copy</span>
                    </button>
                </div>
                <div class="cb-helper">
                    <i class="fa-solid fa-circle-info"></i>
                    Password masked but copies actual value
                </div>
            </div>

            <div class="divider"></div>

            <div class="helper-text">
                <i class="fa-solid fa-circle-info"></i>
                Great for sharing credentials securely
            </div>
        </div>
    </div>

    <div class="content-card">
        <div class="card-header">
            <div class="card-header-left">
                <div class="card-icon bg-success">
                    <i class="fa-solid fa-terminal"></i>
                </div>
                <div>
                    <h3>Terminal Commands</h3>
                    <p class="card-subtitle">Quick command copy</p>
                </div>
            </div>
        </div>
        <div class="card-body">
            <div class="cb-example">
                <label class="cb-label">Laravel Install</label>
                <div class="code-block-clipboard">
                    <button class="copy-btn-absolute" onclick="copyCode(this)">
                        <i class="fa-solid fa-copy"></i>
                        <span>Copy</span>
                    </button>
                    <code>composer create-project laravel/laravel project-name</code>
                </div>
            </div>

            <div class="cb-example">
                <label class="cb-label">NPM Install</label>
                <div class="code-block-clipboard">
                    <button class="copy-btn-absolute" onclick="copyCode(this)">
                        <i class="fa-solid fa-copy"></i>
                        <span>Copy</span>
                    </button>
                    <code>npm install && npm run dev</code>
                </div>
            </div>

            <div class="cb-example">
                <label class="cb-label">Git Clone</label>
                <div class="code-block-clipboard">
                    <button class="copy-btn-absolute" onclick="copyCode(this)">
                        <i class="fa-solid fa-copy"></i>
                        <span>Copy</span>
                    </button>
                    <code>git clone https://github.com/user/repository.git</code>
                </div>
            </div>

            <div class="divider"></div>

            <div class="helper-text">
                <i class="fa-solid fa-circle-info"></i>
                One-click copy for developer commands
            </div>
        </div>
    </div>
</div>

<!-- Toast Notification Demo -->
<div class="cb-section-title">
    <i class="fa-solid fa-bell"></i>
    Toast Notification Demo
</div>

<div class="clipboard-grid full-width">
    <div class="content-card">
        <div class="card-header">
            <div class="card-header-left">
                <div class="card-icon bg-info">
                    <i class="fa-solid fa-bell"></i>
                </div>
                <div>
                    <h3>Visual Feedback</h3>
                    <p class="card-subtitle">Toast notification on successful copy</p>
                </div>
            </div>
        </div>
        <div class="card-body">
            <div style="display: flex; gap: 12px; flex-wrap: wrap;">
                <button class="btn btn-primary" onclick="showToast('Copied to clipboard!')">
                    <i class="fa-solid fa-copy"></i>
                    Copy & Show Toast
                </button>
                <button class="btn btn-success" onclick="showToast('URL copied successfully!', 'success')">
                    <i class="fa-solid fa-link"></i>
                    Copy URL
                </button>
                <button class="btn btn-warning" onclick="showToast('Code snippet copied!', 'warning')">
                    <i class="fa-solid fa-code"></i>
                    Copy Code
                </button>
                <button class="btn btn-info" onclick="showToast('Email address copied!', 'info')">
                    <i class="fa-solid fa-envelope"></i>
                    Copy Email
                </button>
            </div>

            <div class="divider"></div>

            <div class="form-group">
                <h4 style="font-size: 14px; margin-bottom: 12px;">Toast Features:</h4>
                <ul class="feature-list">
                    <li>
                        <i class="fa-solid fa-circle-check"></i>
                        <span>Slide-in animation</span>
                    </li>
                    <li>
                        <i class="fa-solid fa-circle-check"></i>
                        <span>Auto-dismiss after 3 seconds</span>
                    </li>
                    <li>
                        <i class="fa-solid fa-circle-check"></i>
                        <span>Success icon included</span>
                    </li>
                    <li>
                        <i class="fa-solid fa-circle-check"></i>
                        <span>Positioned bottom-right</span>
                    </li>
                </ul>
            </div>
        </div>
    </div>
</div>

<!-- Implementation Guide -->
<div class="clipboard-grid full-width">
    <div class="content-card">
        <div class="card-header">
            <div class="card-header-left">
                <div class="card-icon bg-success">
                    <i class="fa-solid fa-code"></i>
                </div>
                <div>
                    <h3>Implementation Guide</h3>
                    <p class="card-subtitle">How to implement clipboard with JavaScript</p>
                </div>
            </div>
        </div>
        <div class="card-body">
            <div class="code-block">
                <div style="color: var(--text-tertiary); margin-bottom: 8px;">1. Modern Clipboard API:</div>
                <code style="color: var(--accent);">
                    function copyToClipboard(button, text) {<br>
                    &nbsp;&nbsp;navigator.clipboard.writeText(text)<br>
                    &nbsp;&nbsp;&nbsp;&nbsp;.then(() => {<br>
                    &nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;// Change button state<br>
                    &nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;button.classList.add('copied');<br>
                    &nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;button.innerHTML = '<i class="fa-solid fa-check"></i> Copied!';<br>
                    &nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;<br>
                    &nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;// Reset after 2 seconds<br>
                    &nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;setTimeout(() => {<br>
                    &nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;button.classList.remove('copied');<br>
                    &nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;button.innerHTML = '<i class="fa-solid fa-copy"></i> Copy';<br>
                    &nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;}, 2000);<br>
                    &nbsp;&nbsp;&nbsp;&nbsp;})<br>
                    &nbsp;&nbsp;&nbsp;&nbsp;.catch(err => {<br>
                    &nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;console.error('Failed to copy:', err);<br>
                    &nbsp;&nbsp;&nbsp;&nbsp;});<br>
                    }
                </code>
            </div>

            <div class="code-block">
                <div style="color: var(--text-tertiary); margin-bottom: 8px; margin-top: 16px;">2. Copy from Code Block:</div>
                <code style="color: var(--success);">
                    function copyCode(button) {<br>
                    &nbsp;&nbsp;const codeBlock = button.parentElement.querySelector('code');<br>
                    &nbsp;&nbsp;const text = codeBlock.textContent;<br>
                    &nbsp;&nbsp;<br>
                    &nbsp;&nbsp;navigator.clipboard.writeText(text).then(() => {<br>
                    &nbsp;&nbsp;&nbsp;&nbsp;button.classList.add('copied');<br>
                    &nbsp;&nbsp;&nbsp;&nbsp;button.innerHTML = '<i class="fa-solid fa-check"></i> Copied!';<br>
                    &nbsp;&nbsp;&nbsp;&nbsp;<br>
                    &nbsp;&nbsp;&nbsp;&nbsp;setTimeout(() => {<br>
                    &nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;button.classList.remove('copied');<br>
                    &nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;button.innerHTML = '<i class="fa-solid fa-copy"></i> Copy';<br>
                    &nbsp;&nbsp;&nbsp;&nbsp;}, 2000);<br>
                    &nbsp;&nbsp;});<br>
                    }
                </code>
            </div>

            <div class="code-block">
                <div style="color: var(--text-tertiary); margin-bottom: 8px; margin-top: 16px;">3. Show Toast Notification:</div>
                <code style="color: var(--warning);">
                    function showToast(message) {<br>
                    &nbsp;&nbsp;// Remove existing toast<br>
                    &nbsp;&nbsp;const existingToast = document.querySelector('.copy-toast');<br>
                    &nbsp;&nbsp;if (existingToast) existingToast.remove();<br>
                    &nbsp;&nbsp;<br>
                    &nbsp;&nbsp;// Create new toast<br>
                    &nbsp;&nbsp;const toast = document.createElement('div');<br>
                    &nbsp;&nbsp;toast.className = 'copy-toast';<br>
                    &nbsp;&nbsp;toast.innerHTML = `<br>
                    &nbsp;&nbsp;&nbsp;&nbsp;<i class="fa-solid fa-circle-check"></i><br>
                    &nbsp;&nbsp;&nbsp;&nbsp;<span>${message}</span><br>
                    &nbsp;&nbsp;`;<br>
                    &nbsp;&nbsp;document.body.appendChild(toast);<br>
                    &nbsp;&nbsp;<br>
                    &nbsp;&nbsp;// Remove after 3 seconds<br>
                    &nbsp;&nbsp;setTimeout(() => {<br>
                    &nbsp;&nbsp;&nbsp;&nbsp;toast.classList.add('hide');<br>
                    &nbsp;&nbsp;&nbsp;&nbsp;setTimeout(() => toast.remove(), 300);<br>
                    &nbsp;&nbsp;}, 3000);<br>
                    }
                </code>
            </div>

            <div class="code-block">
                <div style="color: var(--text-tertiary); margin-bottom: 8px; margin-top: 16px;">4. Fallback for Older Browsers:</div>
                <code style="color: var(--info);">
                    function fallbackCopy(text) {<br>
                    &nbsp;&nbsp;const textarea = document.createElement('textarea');<br>
                    &nbsp;&nbsp;textarea.value = text;<br>
                    &nbsp;&nbsp;textarea.style.position = 'fixed';<br>
                    &nbsp;&nbsp;textarea.style.opacity = '0';<br>
                    &nbsp;&nbsp;document.body.appendChild(textarea);<br>
                    &nbsp;&nbsp;textarea.select();<br>
                    &nbsp;&nbsp;document.execCommand('copy');<br>
                    &nbsp;&nbsp;document.body.removeChild(textarea);<br>
                    }
                </code>
            </div>

            <div class="divider"></div>

            <div class="form-group">
                <h4 style="font-size: 14px; margin-bottom: 12px;">Browser Support:</h4>
                <ul class="feature-list">
                    <li>
                        <i class="fa-solid fa-circle-check"></i>
                        <span><strong>Chrome 66+</strong> - Full support</span>
                    </li>
                    <li>
                        <i class="fa-solid fa-circle-check"></i>
                        <span><strong>Firefox 63+</strong> - Full support</span>
                    </li>
                    <li>
                        <i class="fa-solid fa-circle-check"></i>
                        <span><strong>Safari 13.1+</strong> - Full support</span>
                    </li>
                    <li>
                        <i class="fa-solid fa-circle-check"></i>
                        <span><strong>Edge 79+</strong> - Full support</span>
                    </li>
                </ul>
            </div>
        </div>
    </div>
</div>

@endsection

@push('scripts')
<script>
// Copy to clipboard function
function copyToClipboard(button, text) {
    // Try modern Clipboard API first
    if (navigator.clipboard && navigator.clipboard.writeText) {
        navigator.clipboard.writeText(text).then(() => {
            handleCopySuccess(button);
        }).catch(err => {
            console.error('Failed to copy:', err);
            handleCopyError();
        });
    } else {
        // Fallback for older browsers
        try {
            const textarea = document.createElement('textarea');
            textarea.value = text;
            textarea.style.position = 'fixed';
            textarea.style.opacity = '0';
            textarea.style.left = '-9999px';
            document.body.appendChild(textarea);
            textarea.select();
            document.execCommand('copy');
            document.body.removeChild(textarea);
            
            handleCopySuccess(button);
        } catch (err) {
            console.error('Fallback copy failed:', err);
            handleCopyError();
        }
    }
}

// Handle copy success
function handleCopySuccess(button) {
    const originalHTML = button.innerHTML;
    
    // Add copied class
    button.classList.add('copied');
    
    // Change button content
    button.innerHTML = '<i class="fa-solid fa-check"></i> <span>Copied!</span>';
    
    // Show toast notification
    showToast('Copied to clipboard!');
    
    // Reset button after 2 seconds
    setTimeout(() => {
        button.classList.remove('copied');
        button.innerHTML = originalHTML;
    }, 2000);
}

// Handle copy error
function handleCopyError() {
    Swal.fire({
        toast: true,
        position: 'top-end',
        icon: 'error',
        title: 'Failed to copy to clipboard',
        showConfirmButton: false,
        timer: 2000
    });
}

// Copy from code block
function copyCode(button) {
    const codeBlock = button.parentElement.querySelector('code');
    
    if (!codeBlock) {
        console.error('Code block not found');
        handleCopyError();
        return;
    }
    
    const text = codeBlock.textContent;
    copyToClipboard(button, text);
}

// Show toast notification
function showToast(message, type = 'success') {
    // Remove existing toast
    const existingToast = document.querySelector('.copy-toast');
    if (existingToast) {
        existingToast.remove();
    }
    
    // Create toast element
    const toast = document.createElement('div');
    toast.className = 'copy-toast';
    
    // Set icon based on type
    let icon = 'fa-solid fa-circle-check';
    let bgColor = 'var(--success)';
    let boxShadow = '0 4px 16px rgba(16, 185, 129, 0.3)';
    
    if (type === 'warning') {
        icon = 'fa-solid fa-triangle-exclamation';
        bgColor = 'var(--warning)';
        boxShadow = '0 4px 16px rgba(245, 158, 11, 0.3)';
    } else if (type === 'info') {
        icon = 'fa-solid fa-circle-info';
        bgColor = 'var(--info)';
        boxShadow = '0 4px 16px rgba(59, 130, 246, 0.3)';
    }
    
    toast.style.background = bgColor;
    toast.style.boxShadow = boxShadow;
    
    toast.innerHTML = `
        <i class="${icon}"></i>
        <span>${message}</span>
    `;
    
    document.body.appendChild(toast);
    
    // Remove after 3 seconds
    setTimeout(() => {
        toast.classList.add('hide');
        setTimeout(() => {
            if (toast.parentElement) {
                toast.remove();
            }
        }, 300);
    }, 3000);
}

// Page header buttons
$(document).ready(function() {
    // Try Clipboard button
    $('.page-header .btn-primary').on('click', function() {
        const sampleText = 'https://example.com';
        copyToClipboard(this, sampleText);
    });
    
    // Documentation button
    $('.page-header .btn-secondary').on('click', function() {
        Swal.fire({
            icon: 'info',
            title: 'Clipboard Documentation',
            html: `
                <div style="text-align: left; font-size: 13px;">
                    <p><strong>Clipboard API</strong> allows copying text to clipboard.</p>
                    <p><strong>Features:</strong></p>
                    <ul>
                        <li>One-click copy</li>
                        <li>Visual feedback</li>
                        <li>Toast notifications</li>
                        <li>Fallback support</li>
                    </ul>
                </div>
            `,
            confirmButtonText: 'OK',
            confirmButtonColor: 'var(--accent)'
        });
    });
    
    // Toast demo buttons
    $('.btn-primary[onclick*="showToast"]').on('click', function(e) {
        e.preventDefault();
        const text = 'Sample text copied';
        copyToClipboard(this, text);
    });
    
    $('.btn-success[onclick*="showToast"]').on('click', function(e) {
        e.preventDefault();
        const url = 'https://example.com/page';
        copyToClipboard(this, url);
    });
    
    $('.btn-warning[onclick*="showToast"]').on('click', function(e) {
        e.preventDefault();
        const code = 'console.log("Hello World");';
        copyToClipboard(this, code);
    });
    
    $('.btn-info[onclick*="showToast"]').on('click', function(e) {
        e.preventDefault();
        const email = 'contact@example.com';
        copyToClipboard(this, email);
    });
});
</script>
@endpush
