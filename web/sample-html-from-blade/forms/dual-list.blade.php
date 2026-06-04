@extends('layouts.app')

@section('title', 'Dual List Boxes')

@push('styles')
@endpush

@section('content')
<div class="page-header">
    <div>
        <h1>Dual List Boxes</h1>
        <p>Transfer items between two lists with search, select, and bulk operations</p>
    </div>
    <div style="display: flex; gap: 10px;">
        <button class="btn btn-secondary">
            <i class="fa-solid fa-book"></i>
            <span>Documentation</span>
        </button>
        <button class="btn btn-primary">
            <i class="fa-solid fa-arrows-left-right"></i>
            <span>Try Dual List</span>
        </button>
    </div>
</div>

<!-- Info Alert -->
<div class="content-card" style="margin-bottom: 24px; border-left: 4px solid var(--accent);">
    <div class="card-body" style="padding: 16px 20px;">
        <div style="display: flex; gap: 12px; align-items: start;">
            <i class="fa-solid fa-circle-info" style="color: var(--accent); font-size: 20px; margin-top: 2px;"></i>
            <div style="flex: 1;">
                <h4 style="margin-bottom: 4px; font-size: 14px;">What is Dual List Box?</h4>
                <p style="font-size: 13px; color: var(--text-secondary); margin: 0;">Dual List Box (Transfer List) allows users to move items between two lists. Commonly used for assigning users to roles, selecting permissions, or managing collections. Supports search, select all, and bulk transfer operations.</p>
            </div>
        </div>
    </div>
</div>

<!-- Basic Dual List -->
<div class="dl-section-title">
    <i class="fa-solid fa-arrows-left-right"></i>
    Basic Dual List <span class="badge badge-primary">Essential</span>
</div>

<div class="dual-list-grid">
    <div class="content-card">
        <div class="card-header">
            <div class="card-header-left">
                <div class="card-icon bg-primary">
                    <i class="fa-solid fa-users"></i>
                </div>
                <div>
                    <h3>Assign Users to Group</h3>
                    <p class="card-subtitle">Simple transfer list</p>
                </div>
            </div>
        </div>
        <div class="card-body">
            <div class="dl-example">
                <label class="dl-label">
                    User Assignment
                    <span class="dl-hint">Select & transfer</span>
                </label>
                <div class="dual-list-container">
                    <!-- Available List -->
                    <div class="list-box">
                        <div class="list-box-header">
                            <div class="list-box-title">Available Users</div>
                            <div class="list-box-count">8</div>
                        </div>
                        <div class="list-box-search">
                            <input type="text" placeholder="Search users...">
                        </div>
                        <div class="list-box-items">
                            <div class="list-box-item selected">
                                <input type="checkbox" checked>
                                <div class="list-box-item-content">
                                    <div class="list-box-item-title">John Doe</div>
                                    <div class="list-box-item-desc">john@example.com</div>
                                </div>
                            </div>
                            <div class="list-box-item selected">
                                <input type="checkbox" checked>
                                <div class="list-box-item-content">
                                    <div class="list-box-item-title">Jane Smith</div>
                                    <div class="list-box-item-desc">jane@example.com</div>
                                </div>
                            </div>
                            <div class="list-box-item">
                                <input type="checkbox">
                                <div class="list-box-item-content">
                                    <div class="list-box-item-title">Bob Johnson</div>
                                    <div class="list-box-item-desc">bob@example.com</div>
                                </div>
                            </div>
                            <div class="list-box-item">
                                <input type="checkbox">
                                <div class="list-box-item-content">
                                    <div class="list-box-item-title">Alice Brown</div>
                                    <div class="list-box-item-desc">alice@example.com</div>
                                </div>
                            </div>
                            <div class="list-box-item">
                                <input type="checkbox">
                                <div class="list-box-item-content">
                                    <div class="list-box-item-title">Mike Wilson</div>
                                    <div class="list-box-item-desc">mike@example.com</div>
                                </div>
                            </div>
                        </div>
                    </div>

                    <!-- Transfer Buttons -->
                    <div class="transfer-buttons">
                        <button class="transfer-btn" title="Move selected to right">
                            <i class="fa-solid fa-chevron-right"></i>
                        </button>
                        <button class="transfer-btn" title="Move selected to left">
                            <i class="fa-solid fa-chevron-left"></i>
                        </button>
                        <button class="transfer-btn" title="Move all to right">
                            <i class="fa-solid fa-angles-right"></i>
                        </button>
                        <button class="transfer-btn" title="Move all to left">
                            <i class="fa-solid fa-angles-left"></i>
                        </button>
                    </div>

                    <!-- Selected List -->
                    <div class="list-box">
                        <div class="list-box-header">
                            <div class="list-box-title">Selected Users</div>
                            <div class="list-box-count">2</div>
                        </div>
                        <div class="list-box-items">
                            <div class="list-box-item">
                                <input type="checkbox">
                                <div class="list-box-item-content">
                                    <div class="list-box-item-title">Sarah Davis</div>
                                    <div class="list-box-item-desc">sarah@example.com</div>
                                </div>
                            </div>
                            <div class="list-box-item">
                                <input type="checkbox">
                                <div class="list-box-item-content">
                                    <div class="list-box-item-title">Tom Anderson</div>
                                    <div class="list-box-item-desc">tom@example.com</div>
                                </div>
                            </div>
                        </div>
                    </div>
                </div>
                <div class="dl-helper">
                    <i class="fa-solid fa-circle-info"></i>
                    Select items and use arrows to transfer
                </div>
            </div>

            <div class="divider"></div>

            <div class="form-group">
                <h4 style="font-size: 14px; margin-bottom: 12px;">Features:</h4>
                <ul class="feature-list">
                    <li>
                        <i class="fa-solid fa-circle-check"></i>
                        <span>Search functionality</span>
                    </li>
                    <li>
                        <i class="fa-solid fa-circle-check"></i>
                        <span>Multiple selection</span>
                    </li>
                    <li>
                        <i class="fa-solid fa-circle-check"></i>
                        <span>Bulk transfer</span>
                    </li>
                    <li>
                        <i class="fa-solid fa-circle-check"></i>
                        <span>Item count badges</span>
                    </li>
                </ul>
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
                    <h3>Assign Roles</h3>
                    <p class="card-subtitle">Permission management</p>
                </div>
            </div>
        </div>
        <div class="card-body">
            <div class="dl-example">
                <label class="dl-label">
                    Role Permissions
                    <span class="dl-hint">Manage access</span>
                </label>
                <div class="dual-list-container">
                    <!-- Available Permissions -->
                    <div class="list-box">
                        <div class="list-box-header">
                            <div class="list-box-title">Available Permissions</div>
                            <div class="list-box-count">10</div>
                        </div>
                        <div class="list-box-search">
                            <input type="text" placeholder="Search permissions...">
                        </div>
                        <div class="list-box-items">
                            <div class="list-box-item selected">
                                <input type="checkbox" checked>
                                <div class="list-box-item-content">
                                    <div class="list-box-item-title">Create Posts</div>
                                    <div class="list-box-item-desc">Allow creating new posts</div>
                                </div>
                            </div>
                            <div class="list-box-item selected">
                                <input type="checkbox" checked>
                                <div class="list-box-item-content">
                                    <div class="list-box-item-title">Edit Posts</div>
                                    <div class="list-box-item-desc">Allow editing existing posts</div>
                                </div>
                            </div>
                            <div class="list-box-item selected">
                                <input type="checkbox" checked>
                                <div class="list-box-item-content">
                                    <div class="list-box-item-title">Delete Posts</div>
                                    <div class="list-box-item-desc">Allow deleting posts</div>
                                </div>
                            </div>
                            <div class="list-box-item">
                                <input type="checkbox">
                                <div class="list-box-item-content">
                                    <div class="list-box-item-title">Manage Users</div>
                                    <div class="list-box-item-desc">User administration</div>
                                </div>
                            </div>
                            <div class="list-box-item">
                                <input type="checkbox">
                                <div class="list-box-item-content">
                                    <div class="list-box-item-title">View Reports</div>
                                    <div class="list-box-item-desc">Access analytics</div>
                                </div>
                            </div>
                        </div>
                    </div>

                    <!-- Transfer Buttons -->
                    <div class="transfer-buttons">
                        <button class="transfer-btn" title="Move selected to right">
                            <i class="fa-solid fa-chevron-right"></i>
                        </button>
                        <button class="transfer-btn" title="Move selected to left">
                            <i class="fa-solid fa-chevron-left"></i>
                        </button>
                        <button class="transfer-btn" title="Move all to right">
                            <i class="fa-solid fa-angles-right"></i>
                        </button>
                        <button class="transfer-btn" title="Move all to left">
                            <i class="fa-solid fa-angles-left"></i>
                        </button>
                    </div>

                    <!-- Assigned Permissions -->
                    <div class="list-box">
                        <div class="list-box-header">
                            <div class="list-box-title">Assigned Permissions</div>
                            <div class="list-box-count">3</div>
                        </div>
                        <div class="list-box-items">
                            <div class="list-box-item">
                                <input type="checkbox">
                                <div class="list-box-item-content">
                                    <div class="list-box-item-title">Publish Posts</div>
                                    <div class="list-box-item-desc">Allow publishing</div>
                                </div>
                            </div>
                            <div class="list-box-item">
                                <input type="checkbox">
                                <div class="list-box-item-content">
                                    <div class="list-box-item-title">Moderate Comments</div>
                                    <div class="list-box-item-desc">Comment approval</div>
                                </div>
                            </div>
                            <div class="list-box-item">
                                <input type="checkbox">
                                <div class="list-box-item-content">
                                    <div class="list-box-item-title">Upload Files</div>
                                    <div class="list-box-item-desc">Media library access</div>
                                </div>
                            </div>
                        </div>
                    </div>
                </div>
            </div>

            <div class="divider"></div>

            <div class="helper-text">
                <i class="fa-solid fa-circle-info"></i>
                Select permissions and transfer to assigned list
            </div>
        </div>
    </div>
</div>

<!-- Advanced Dual List -->
<div class="dl-section-title">
    <i class="fa-solid fa-wand-magic-sparkles"></i>
    Advanced Dual List
</div>

<div class="dual-list-grid">
    <div class="content-card">
        <div class="card-header">
            <div class="card-header-left">
                <div class="card-icon bg-info">
                    <i class="fa-solid fa-tags"></i>
                </div>
                <div>
                    <h3>Product Categories</h3>
                    <p class="card-subtitle">Category management with icons</p>
                </div>
            </div>
        </div>
        <div class="card-body">
            <div class="dl-example">
                <label class="dl-label">
                    Product Categories
                    <span class="dl-hint">With icons</span>
                </label>
                <div class="dual-list-container">
                    <!-- Available Categories -->
                    <div class="list-box">
                        <div class="list-box-header">
                            <div class="list-box-title">📦 All Categories</div>
                            <div class="list-box-count">8</div>
                        </div>
                        <div class="list-box-search">
                            <input type="text" placeholder="Search categories...">
                        </div>
                        <div class="list-box-items">
                            <div class="list-box-item">
                                <input type="checkbox">
                                <div class="list-box-item-content">
                                    <div class="list-box-item-title">💻 Electronics</div>
                                    <div class="list-box-item-desc">156 products</div>
                                </div>
                            </div>
                            <div class="list-box-item">
                                <input type="checkbox">
                                <div class="list-box-item-content">
                                    <div class="list-box-item-title">👕 Clothing</div>
                                    <div class="list-box-item-desc">243 products</div>
                                </div>
                            </div>
                            <div class="list-box-item">
                                <input type="checkbox">
                                <div class="list-box-item-content">
                                    <div class="list-box-item-title">🏠 Home & Garden</div>
                                    <div class="list-box-item-desc">189 products</div>
                                </div>
                            </div>
                            <div class="list-box-item">
                                <input type="checkbox">
                                <div class="list-box-item-content">
                                    <div class="list-box-item-title">📚 Books</div>
                                    <div class="list-box-item-desc">421 products</div>
                                </div>
                            </div>
                            <div class="list-box-item">
                                <input type="checkbox">
                                <div class="list-box-item-content">
                                    <div class="list-box-item-title">⚽ Sports</div>
                                    <div class="list-box-item-desc">98 products</div>
                                </div>
                            </div>
                        </div>
                    </div>

                    <!-- Transfer Buttons -->
                    <div class="transfer-buttons">
                        <button class="transfer-btn" title="Move right">
                            <i class="fa-solid fa-chevron-right"></i>
                        </button>
                        <button class="transfer-btn" title="Move left">
                            <i class="fa-solid fa-chevron-left"></i>
                        </button>
                        <button class="transfer-btn" title="Move all right">
                            <i class="fa-solid fa-angles-right"></i>
                        </button>
                        <button class="transfer-btn" title="Move all left">
                            <i class="fa-solid fa-angles-left"></i>
                        </button>
                    </div>

                    <!-- Selected Categories -->
                    <div class="list-box">
                        <div class="list-box-header">
                            <div class="list-box-title">✅ Selected Categories</div>
                            <div class="list-box-count">2</div>
                        </div>
                        <div class="list-box-items">
                            <div class="list-box-item">
                                <input type="checkbox">
                                <div class="list-box-item-content">
                                    <div class="list-box-item-title">🎨 Art & Design</div>
                                    <div class="list-box-item-desc">67 products</div>
                                </div>
                            </div>
                            <div class="list-box-item">
                                <input type="checkbox">
                                <div class="list-box-item-content">
                                    <div class="list-box-item-title">🚗 Automotive</div>
                                    <div class="list-box-item-desc">134 products</div>
                                </div>
                            </div>
                        </div>
                    </div>
                </div>
            </div>

            <div class="divider"></div>

            <div class="helper-text">
                <i class="fa-solid fa-circle-info"></i>
                Categories with emoji icons and product counts
            </div>
        </div>
    </div>

    <div class="content-card">
        <div class="card-header">
            <div class="card-header-left">
                <div class="card-icon bg-warning">
                    <i class="fa-solid fa-language"></i>
                </div>
                <div>
                    <h3>Language Selection</h3>
                    <p class="card-subtitle">Multi-language support</p>
                </div>
            </div>
        </div>
        <div class="card-body">
            <div class="dl-example">
                <label class="dl-label">
                    Website Languages
                    <span class="dl-hint">Translation</span>
                </label>
                <div class="dual-list-container">
                    <!-- Available Languages -->
                    <div class="list-box">
                        <div class="list-box-header">
                            <div class="list-box-title">🌍 Available Languages</div>
                            <div class="list-box-count">8</div>
                        </div>
                        <div class="list-box-search">
                            <input type="text" placeholder="Search languages...">
                        </div>
                        <div class="list-box-items">
                            <div class="list-box-item selected">
                                <input type="checkbox" checked>
                                <div class="list-box-item-content">
                                    <div class="list-box-item-title">🇬🇧 English</div>
                                    <div class="list-box-item-desc">English (US)</div>
                                </div>
                            </div>
                            <div class="list-box-item selected">
                                <input type="checkbox" checked>
                                <div class="list-box-item-content">
                                    <div class="list-box-item-title">🇪🇸 Spanish</div>
                                    <div class="list-box-item-desc">Español</div>
                                </div>
                            </div>
                            <div class="list-box-item selected">
                                <input type="checkbox" checked>
                                <div class="list-box-item-content">
                                    <div class="list-box-item-title">🇫🇷 French</div>
                                    <div class="list-box-item-desc">Français</div>
                                </div>
                            </div>
                            <div class="list-box-item">
                                <input type="checkbox">
                                <div class="list-box-item-content">
                                    <div class="list-box-item-title">🇩🇪 German</div>
                                    <div class="list-box-item-desc">Deutsch</div>
                                </div>
                            </div>
                            <div class="list-box-item">
                                <input type="checkbox">
                                <div class="list-box-item-content">
                                    <div class="list-box-item-title">🇯🇵 Japanese</div>
                                    <div class="list-box-item-desc">日本語</div>
                                </div>
                            </div>
                        </div>
                    </div>

                    <!-- Transfer Buttons -->
                    <div class="transfer-buttons">
                        <button class="transfer-btn" title="Move right">
                            <i class="fa-solid fa-chevron-right"></i>
                        </button>
                        <button class="transfer-btn" title="Move left">
                            <i class="fa-solid fa-chevron-left"></i>
                        </button>
                        <button class="transfer-btn" title="Move all right">
                            <i class="fa-solid fa-angles-right"></i>
                        </button>
                        <button class="transfer-btn" title="Move all left">
                            <i class="fa-solid fa-angles-left"></i>
                        </button>
                    </div>

                    <!-- Selected Languages -->
                    <div class="list-box">
                        <div class="list-box-header">
                            <div class="list-box-title">✅ Active Languages</div>
                            <div class="list-box-count">2</div>
                        </div>
                        <div class="list-box-items">
                            <div class="list-box-item">
                                <input type="checkbox">
                                <div class="list-box-item-content">
                                    <div class="list-box-item-title">🇮🇩 Indonesian</div>
                                    <div class="list-box-item-desc">Bahasa Indonesia</div>
                                </div>
                            </div>
                            <div class="list-box-item">
                                <input type="checkbox">
                                <div class="list-box-item-content">
                                    <div class="list-box-item-title">🇨🇳 Chinese</div>
                                    <div class="list-box-item-desc">中文</div>
                                </div>
                            </div>
                        </div>
                    </div>
                </div>
            </div>

            <div class="divider"></div>

            <div class="helper-text">
                <i class="fa-solid fa-circle-info"></i>
                Country flags and native language names
            </div>
        </div>
    </div>
</div>

<!-- Implementation Guide -->
<div class="dual-list-grid full-width">
    <div class="content-card">
        <div class="card-header">
            <div class="card-header-left">
                <div class="card-icon bg-success">
                    <i class="fa-solid fa-code"></i>
                </div>
                <div>
                    <h3>Implementation Guide</h3>
                    <p class="card-subtitle">How to implement dual list with JavaScript</p>
                </div>
            </div>
        </div>
        <div class="card-body">
            <div class="code-block">
                <div style="color: var(--text-tertiary); margin-bottom: 8px;">1. Transfer Selected Items:</div>
                <code style="color: var(--accent);">
                    function transferSelected(fromList, toList) {<br>
                    &nbsp;&nbsp;const selectedItems = fromList.querySelectorAll('.list-box-item.selected');<br>
                    &nbsp;&nbsp;<br>
                    &nbsp;&nbsp;selectedItems.forEach(item => {<br>
                    &nbsp;&nbsp;&nbsp;&nbsp;// Clone and move to target list<br>
                    &nbsp;&nbsp;&nbsp;&nbsp;const clone = item.cloneNode(true);<br>
                    &nbsp;&nbsp;&nbsp;&nbsp;clone.classList.remove('selected');<br>
                    &nbsp;&nbsp;&nbsp;&nbsp;toList.appendChild(clone);<br>
                    &nbsp;&nbsp;&nbsp;&nbsp;item.remove();<br>
                    &nbsp;&nbsp;});<br>
                    &nbsp;&nbsp;<br>
                    &nbsp;&nbsp;updateCounts();<br>
                    }
                </code>
            </div>

            <div class="code-block">
                <div style="color: var(--text-tertiary); margin-bottom: 8px; margin-top: 16px;">2. Select/Deselect Items:</div>
                <code style="color: var(--success);">
                    document.querySelectorAll('.list-box-item').forEach(item => {<br>
                    &nbsp;&nbsp;item.addEventListener('click', (e) => {<br>
                    &nbsp;&nbsp;&nbsp;&nbsp;if (e.target.tagName !== 'INPUT') {<br>
                    &nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;const checkbox = item.querySelector('input[type="checkbox"]');<br>
                    &nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;checkbox.checked = !checkbox.checked;<br>
                    &nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;item.classList.toggle('selected');<br>
                    &nbsp;&nbsp;&nbsp;&nbsp;}<br>
                    &nbsp;&nbsp;});<br>
                    });
                </code>
            </div>

            <div class="code-block">
                <div style="color: var(--text-tertiary); margin-bottom: 8px; margin-top: 16px;">3. Search Filter:</div>
                <code style="color: var(--warning);">
                    function filterList(searchInput, listContainer) {<br>
                    &nbsp;&nbsp;searchInput.addEventListener('input', (e) => {<br>
                    &nbsp;&nbsp;&nbsp;&nbsp;const searchTerm = e.target.value.toLowerCase();<br>
                    &nbsp;&nbsp;&nbsp;&nbsp;const items = listContainer.querySelectorAll('.list-box-item');<br>
                    &nbsp;&nbsp;&nbsp;&nbsp;<br>
                    &nbsp;&nbsp;&nbsp;&nbsp;items.forEach(item => {<br>
                    &nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;const title = item.querySelector('.list-box-item-title').textContent.toLowerCase();<br>
                    &nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;item.style.display = title.includes(searchTerm) ? '' : 'none';<br>
                    &nbsp;&nbsp;&nbsp;&nbsp;});<br>
                    &nbsp;&nbsp;});<br>
                    }
                </code>
            </div>

            <div class="code-block">
                <div style="color: var(--text-tertiary); margin-bottom: 8px; margin-top: 16px;">4. Update Counts:</div>
                <code style="color: var(--info);">
                    function updateCounts() {<br>
                    &nbsp;&nbsp;const leftCount = document.querySelector('.list-box:first-child .list-box-items .list-box-item').length;<br>
                    &nbsp;&nbsp;const rightCount = document.querySelector('.list-box:last-child .list-box-items .list-box-item').length;<br>
                    &nbsp;&nbsp;<br>
                    &nbsp;&nbsp;document.querySelector('.list-box:first-child .list-box-count').textContent = leftCount;<br>
                    &nbsp;&nbsp;document.querySelector('.list-box:last-child .list-box-count').textContent = rightCount;<br>
                    }
                </code>
            </div>

            <div class="divider"></div>

            <div class="form-group">
                <h4 style="font-size: 14px; margin-bottom: 12px;">Popular Libraries:</h4>
                <ul class="feature-list">
                    <li>
                        <i class="fa-solid fa-circle-check"></i>
                        <span><strong>jQuery UI Sortable</strong> - Drag & drop lists</span>
                    </li>
                    <li>
                        <i class="fa-solid fa-circle-check"></i>
                        <span><strong>Multi.js</strong> - Visual multi-select</span>
                    </li>
                    <li>
                        <i class="fa-solid fa-circle-check"></i>
                        <span><strong>Bootstrap Dual Listbox</strong> - Bootstrap integration</span>
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
    // Handle item click to select/deselect
    $(document).on('click', '.list-box-item', function(e) {
        if (e.target.tagName !== 'INPUT') {
            const checkbox = $(this).find('input[type="checkbox"]');
            checkbox.prop('checked', !checkbox.prop('checked'));
            $(this).toggleClass('selected');
        }
    });

    // Handle checkbox change
    $(document).on('change', '.list-box-item input[type="checkbox"]', function() {
        $(this).closest('.list-box-item').toggleClass('selected', this.checked);
    });

    // Transfer selected items to right
    $(document).on('click', '.transfer-buttons .transfer-btn:first-child', function() {
        const container = $(this).closest('.dual-list-container');
        const leftList = container.find('.list-box:first-child .list-box-items');
        const rightList = container.find('.list-box:last-child .list-box-items');
        
        leftList.find('.list-box-item.selected').each(function() {
            const clone = $(this).clone();
            clone.removeClass('selected');
            clone.find('input[type="checkbox"]').prop('checked', false);
            rightList.append(clone);
            $(this).remove();
        });
        
        updateCounts(container);
        showToast('Items transferred', 'success');
    });

    // Transfer selected items to left
    $(document).on('click', '.transfer-buttons .transfer-btn:nth-child(2)', function() {
        const container = $(this).closest('.dual-list-container');
        const leftList = container.find('.list-box:first-child .list-box-items');
        const rightList = container.find('.list-box:last-child .list-box-items');
        
        rightList.find('.list-box-item.selected').each(function() {
            const clone = $(this).clone();
            clone.removeClass('selected');
            clone.find('input[type="checkbox"]').prop('checked', false);
            leftList.append(clone);
            $(this).remove();
        });
        
        updateCounts(container);
        showToast('Items transferred back', 'info');
    });

    // Transfer all items to right
    $(document).on('click', '.transfer-buttons .transfer-btn:nth-child(3)', function() {
        const container = $(this).closest('.dual-list-container');
        const leftList = container.find('.list-box:first-child .list-box-items');
        const rightList = container.find('.list-box:last-child .list-box-items');
        
        leftList.find('.list-box-item').each(function() {
            const clone = $(this).clone();
            clone.removeClass('selected');
            clone.find('input[type="checkbox"]').prop('checked', false);
            rightList.append(clone);
            $(this).remove();
        });
        
        updateCounts(container);
        showToast('All items transferred', 'success');
    });

    // Transfer all items to left
    $(document).on('click', '.transfer-buttons .transfer-btn:nth-child(4)', function() {
        const container = $(this).closest('.dual-list-container');
        const leftList = container.find('.list-box:first-child .list-box-items');
        const rightList = container.find('.list-box:last-child .list-box-items');
        
        rightList.find('.list-box-item').each(function() {
            const clone = $(this).clone();
            clone.removeClass('selected');
            clone.find('input[type="checkbox"]').prop('checked', false);
            leftList.append(clone);
            $(this).remove();
        });
        
        updateCounts(container);
        showToast('All items transferred back', 'info');
    });

    // Search functionality
    $(document).on('input', '.list-box-search input', function() {
        const searchTerm = $(this).val().toLowerCase();
        const listItems = $(this).closest('.list-box').find('.list-box-item');
        
        listItems.each(function() {
            const title = $(this).find('.list-box-item-title').text().toLowerCase();
            const desc = $(this).find('.list-box-item-desc').text().toLowerCase();
            
            if (title.includes(searchTerm) || desc.includes(searchTerm)) {
                $(this).show();
            } else {
                $(this).hide();
            }
        });
    });

    // Update counts
    function updateCounts(container) {
        const leftCount = container.find('.list-box:first-child .list-box-items .list-box-item').length;
        const rightCount = container.find('.list-box:last-child .list-box-items .list-box-item').length;
        
        container.find('.list-box:first-child .list-box-count').text(leftCount);
        container.find('.list-box:last-child .list-box-count').text(rightCount);
    }
});
</script>
@endpush
