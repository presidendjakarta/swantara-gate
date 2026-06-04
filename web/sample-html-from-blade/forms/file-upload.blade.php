@extends('layouts.app')

@section('title', 'File Upload')

@push('styles')
@endpush

@section('content')
<div class="page-header">
    <div>
        <h1>File Upload</h1>
        <p>Drag & drop file upload dengan preview, progress bar, dan validasi file</p>
    </div>
    <div style="display: flex; gap: 10px;">
        <button class="btn btn-secondary">
            <i class="fa-solid fa-book"></i>
            <span>Documentation</span>
        </button>
        <button class="btn btn-primary">
            <i class="fa-solid fa-upload"></i>
            <span>Try Upload</span>
        </button>
    </div>
</div>

<!-- Info Alert -->
<div class="content-card" style="margin-bottom: 24px; border-left: 4px solid var(--accent);">
    <div class="card-body" style="padding: 16px 20px;">
        <div style="display: flex; gap: 12px; align-items: start;">
            <i class="fa-solid fa-circle-info" style="color: var(--accent); font-size: 20px; margin-top: 2px;"></i>
            <div style="flex: 1;">
                <h4 style="margin-bottom: 4px; font-size: 14px;">File Upload Features</h4>
                <p style="font-size: 13px; color: var(--text-secondary); margin: 0;">Modern file upload with drag & drop support, file preview, progress tracking, validation, and multiple file selection. Uses HTML5 File API with visual feedback for better UX.</p>
            </div>
        </div>
    </div>
</div>

<!-- Drag & Drop Upload -->
<div class="fu-section-title">
    <i class="fa-solid fa-cloud-arrow-up"></i>
    Drag & Drop Upload <span class="badge badge-primary">Essential</span>
</div>

<div class="file-upload-grid">
    <div class="content-card">
        <div class="card-header">
            <div class="card-header-left">
                <div class="card-icon bg-primary">
                    <i class="fa-solid fa-cloud-arrow-up"></i>
                </div>
                <div>
                    <h3>Drop Zone</h3>
                    <p class="card-subtitle">Drag & drop files here</p>
                </div>
            </div>
        </div>
        <div class="card-body">
            <div class="fu-example">
                <label class="fu-label">
                    Upload Files
                    <span class="fu-hint">Max 10MB</span>
                </label>
                <div class="upload-zone" id="uploadZone1">
                    <div class="upload-zone-icon">
                        <i class="fa-solid fa-cloud-arrow-up"></i>
                    </div>
                    <div class="upload-zone-title">Drag & Drop files here</div>
                    <div class="upload-zone-desc">or click to browse from your computer</div>
                    <div class="upload-zone-btn">
                        <i class="fa-solid fa-folder-open"></i>
                        Browse Files
                    </div>
                </div>
                <div class="fu-helper">
                    <i class="fa-solid fa-circle-info"></i>
                    Supported: JPG, PNG, GIF, PDF, DOC
                </div>
            </div>

            <div class="fu-example">
                <label class="fu-label">
                    Upload Images
                    <span class="fu-hint">Images only</span>
                </label>
                <div class="upload-zone">
                    <div class="upload-zone-icon" style="color: var(--success);">
                        <i class="fa-solid fa-image"></i>
                    </div>
                    <div class="upload-zone-title">Drop images here</div>
                    <div class="upload-zone-desc">JPG, PNG, GIF up to 5MB each</div>
                    <div class="upload-zone-btn" style="background: var(--success);">
                        <i class="fa-solid fa-upload"></i>
                        Select Images
                    </div>
                </div>
            </div>

            <div class="divider"></div>

            <div class="form-group">
                <h4 style="font-size: 14px; margin-bottom: 12px;">Features:</h4>
                <ul class="feature-list">
                    <li>
                        <i class="fa-solid fa-circle-check"></i>
                        <span>Drag & drop support</span>
                    </li>
                    <li>
                        <i class="fa-solid fa-circle-check"></i>
                        <span>Visual hover feedback</span>
                    </li>
                    <li>
                        <i class="fa-solid fa-circle-check"></i>
                        <span>File type validation</span>
                    </li>
                    <li>
                        <i class="fa-solid fa-circle-check"></i>
                        <span>Size limit enforcement</span>
                    </li>
                </ul>
            </div>
        </div>
    </div>

    <div class="content-card">
        <div class="card-header">
            <div class="card-header-left">
                <div class="card-icon bg-success">
                    <i class="fa-solid fa-file-arrow-up"></i>
                </div>
                <div>
                    <h3>File Input</h3>
                    <p class="card-subtitle">Traditional file picker</p>
                </div>
            </div>
        </div>
        <div class="card-body">
            <div class="fu-example">
                <label class="fu-label">
                    Single File
                    <span class="fu-hint">One file</span>
                </label>
                <div class="file-input-wrapper">
                    <input type="file" id="singleFile">
                    <div class="file-input-label">
                        <i class="fa-solid fa-paperclip file-input-icon"></i>
                        <span class="file-input-text">Choose a file...</span>
                    </div>
                </div>
            </div>

            <div class="fu-example">
                <label class="fu-label">
                    Multiple Files
                    <span class="fu-hint">Multiple selection</span>
                </label>
                <div class="file-input-wrapper">
                    <input type="file" id="multipleFiles" multiple>
                    <div class="file-input-label">
                        <i class="fa-solid fa-folder-open file-input-icon"></i>
                        <span class="file-input-text">Choose files...</span>
                    </div>
                </div>
            </div>

            <div class="fu-example">
                <label class="fu-label">
                    With Selected File
                    <span class="fu-hint">File selected</span>
                </label>
                <div class="file-input-wrapper">
                    <input type="file" id="fileSelected">
                    <div class="file-input-label">
                        <i class="fa-solid fa-file file-input-icon"></i>
                        <div>
                            <div class="file-input-name">document.pdf</div>
                            <div class="file-input-text">2.5 MB • PDF</div>
                        </div>
                    </div>
                </div>
            </div>

            <div class="divider"></div>

            <div class="helper-text">
                <i class="fa-solid fa-circle-info"></i>
                Click on the input to browse files
            </div>
        </div>
    </div>
</div>

<!-- File List & Preview -->
<div class="fu-section-title">
    <i class="fa-solid fa-list"></i>
    File List & Preview
</div>

<div class="file-upload-grid">
    <div class="content-card">
        <div class="card-header">
            <div class="card-header-left">
                <div class="card-icon bg-info">
                    <i class="fa-solid fa-list-check"></i>
                </div>
                <div>
                    <h3>File List</h3>
                    <p class="card-subtitle">Uploaded files with details</p>
                </div>
            </div>
        </div>
        <div class="card-body">
            <div class="fu-example">
                <label class="fu-label">Uploaded Files (4)</label>
                <div class="file-list">
                    <div class="file-item">
                        <div class="file-item-icon">📄</div>
                        <div class="file-item-content">
                            <div class="file-item-name">resume.pdf</div>
                            <div class="file-item-size">245 KB • PDF</div>
                        </div>
                        <button class="file-item-remove">
                            <i class="fa-solid fa-xmark"></i>
                        </button>
                    </div>

                    <div class="file-item">
                        <div class="file-item-icon">🖼️</div>
                        <div class="file-item-content">
                            <div class="file-item-name">profile-photo.jpg</div>
                            <div class="file-item-size">1.8 MB • JPEG</div>
                        </div>
                        <button class="file-item-remove">
                            <i class="fa-solid fa-xmark"></i>
                        </button>
                    </div>

                    <div class="file-item">
                        <div class="file-item-icon">📊</div>
                        <div class="file-item-content">
                            <div class="file-item-name">report-2024.xlsx</div>
                            <div class="file-item-size">523 KB • Excel</div>
                        </div>
                        <button class="file-item-remove">
                            <i class="fa-solid fa-xmark"></i>
                        </button>
                    </div>

                    <div class="file-item">
                        <div class="file-item-icon">📝</div>
                        <div class="file-item-content">
                            <div class="file-item-name">notes.txt</div>
                            <div class="file-item-size">12 KB • Text</div>
                        </div>
                        <button class="file-item-remove">
                            <i class="fa-solid fa-xmark"></i>
                        </button>
                    </div>
                </div>
            </div>

            <div class="divider"></div>

            <div class="helper-text">
                <i class="fa-solid fa-circle-info"></i>
                Click × to remove file from list
            </div>
        </div>
    </div>

    <div class="content-card">
        <div class="card-header">
            <div class="card-header-left">
                <div class="card-icon bg-warning">
                    <i class="fa-solid fa-images"></i>
                </div>
                <div>
                    <h3>Image Preview</h3>
                    <p class="card-subtitle">Grid preview for images</p>
                </div>
            </div>
        </div>
        <div class="card-body">
            <div class="fu-example">
                <label class="fu-label">Selected Images (4)</label>
                <div class="file-preview">
                    <div class="file-preview-item">
                        <div style="width: 100%; height: 100%; background: linear-gradient(135deg, #667eea 0%, #764ba2 100%);"></div>
                        <div class="file-preview-overlay">
                            <button class="file-preview-remove">
                                <i class="fa-solid fa-xmark"></i>
                            </button>
                        </div>
                    </div>

                    <div class="file-preview-item">
                        <div style="width: 100%; height: 100%; background: linear-gradient(135deg, #f093fb 0%, #f5576c 100%);"></div>
                        <div class="file-preview-overlay">
                            <button class="file-preview-remove">
                                <i class="fa-solid fa-xmark"></i>
                            </button>
                        </div>
                    </div>

                    <div class="file-preview-item">
                        <div style="width: 100%; height: 100%; background: linear-gradient(135deg, #4facfe 0%, #00f2fe 100%);"></div>
                        <div class="file-preview-overlay">
                            <button class="file-preview-remove">
                                <i class="fa-solid fa-xmark"></i>
                            </button>
                        </div>
                    </div>

                    <div class="file-preview-item">
                        <div style="width: 100%; height: 100%; background: linear-gradient(135deg, #43e97b 0%, #38f9d7 100%);"></div>
                        <div class="file-preview-overlay">
                            <button class="file-preview-remove">
                                <i class="fa-solid fa-xmark"></i>
                            </button>
                        </div>
                    </div>
                </div>
            </div>

            <div class="divider"></div>

            <div class="helper-text">
                <i class="fa-solid fa-circle-info"></i>
                Hover over image to see remove button
            </div>
        </div>
    </div>
</div>

<!-- Upload Progress -->
<div class="fu-section-title">
    <i class="fa-solid fa-spinner"></i>
    Upload Progress
</div>

<div class="file-upload-grid">
    <div class="content-card">
        <div class="card-header">
            <div class="card-header-left">
                <div class="card-icon bg-primary">
                    <i class="fa-solid fa-upload"></i>
                </div>
                <div>
                    <h3>Progress Bars</h3>
                    <p class="card-subtitle">Real-time upload tracking</p>
                </div>
            </div>
        </div>
        <div class="card-body">
            <div class="fu-example">
                <label class="fu-label">Uploading...</label>
                <div class="file-item">
                    <div class="file-item-icon">📄</div>
                    <div class="file-item-content">
                        <div class="file-item-name">presentation.pptx</div>
                        <div class="file-item-size">3.2 MB • 65% uploaded</div>
                        <div class="upload-progress">
                            <div class="progress-bar-bg">
                                <div class="progress-bar-fill" style="width: 65%;"></div>
                            </div>
                            <div class="progress-text">
                                <span>2.08 MB / 3.2 MB</span>
                                <span>65%</span>
                            </div>
                        </div>
                    </div>
                </div>
            </div>

            <div class="fu-example">
                <label class="fu-label">Completed</label>
                <div class="file-item">
                    <div class="file-item-icon">✅</div>
                    <div class="file-item-content">
                        <div class="file-item-name">backup.zip</div>
                        <div class="file-item-size">15.8 MB • Upload complete</div>
                        <div class="upload-progress">
                            <div class="progress-bar-bg">
                                <div class="progress-bar-fill success" style="width: 100%;"></div>
                            </div>
                            <div class="progress-text">
                                <span>Upload successful</span>
                                <span style="color: var(--success);">✓</span>
                            </div>
                        </div>
                    </div>
                </div>
            </div>

            <div class="fu-example">
                <label class="fu-label">Multiple Files</label>
                <div class="file-item">
                    <div class="file-item-icon">🖼️</div>
                    <div class="file-item-content">
                        <div class="file-item-name">photo-001.jpg</div>
                        <div class="file-item-size">2.1 MB • Uploading...</div>
                        <div class="upload-progress">
                            <div class="progress-bar-bg">
                                <div class="progress-bar-fill" style="width: 42%;"></div>
                            </div>
                            <div class="progress-text">
                                <span>882 KB / 2.1 MB</span>
                                <span>42%</span>
                            </div>
                        </div>
                    </div>
                </div>

                <div class="file-item">
                    <div class="file-item-icon">🖼️</div>
                    <div class="file-item-content">
                        <div class="file-item-name">photo-002.jpg</div>
                        <div class="file-item-size">1.9 MB • Uploading...</div>
                        <div class="upload-progress">
                            <div class="progress-bar-bg">
                                <div class="progress-bar-fill" style="width: 28%;"></div>
                            </div>
                            <div class="progress-text">
                                <span>532 KB / 1.9 MB</span>
                                <span>28%</span>
                            </div>
                        </div>
                    </div>
                </div>
            </div>

            <div class="divider"></div>

            <div class="helper-text">
                <i class="fa-solid fa-circle-info"></i>
                Progress updates in real-time
            </div>
        </div>
    </div>

    <div class="content-card">
        <div class="card-header">
            <div class="card-header-left">
                <div class="card-icon bg-success">
                    <i class="fa-solid fa-triangle-exclamation"></i>
                </div>
                <div>
                    <h3>Upload States</h3>
                    <p class="card-subtitle">Different upload statuses</p>
                </div>
            </div>
        </div>
        <div class="card-body">
            <div class="fu-example">
                <label class="fu-label">Waiting</label>
                <div class="file-item">
                    <div class="file-item-icon">⏳</div>
                    <div class="file-item-content">
                        <div class="file-item-name">video.mp4</div>
                        <div class="file-item-size">48.5 MB • Waiting to upload</div>
                        <div class="upload-progress">
                            <div class="progress-bar-bg">
                                <div class="progress-bar-fill" style="width: 0%;"></div>
                            </div>
                            <div class="progress-text">
                                <span>Queued</span>
                                <span>0%</span>
                            </div>
                        </div>
                    </div>
                </div>
            </div>

            <div class="fu-example">
                <label class="fu-label">Error</label>
                <div class="file-item" style="border-color: var(--danger);">
                    <div class="file-item-icon">❌</div>
                    <div class="file-item-content">
                        <div class="file-item-name">large-file.zip</div>
                        <div class="file-item-size" style="color: var(--danger);">52.3 MB • Upload failed</div>
                        <div class="upload-progress">
                            <div class="progress-bar-bg">
                                <div class="progress-bar-fill" style="width: 35%; background: var(--danger);"></div>
                            </div>
                            <div class="progress-text">
                                <span style="color: var(--danger);">File too large (max 50MB)</span>
                                <span>35%</span>
                            </div>
                        </div>
                    </div>
                    <button class="file-item-remove">
                        <i class="fa-solid fa-rotate-right"></i>
                    </button>
                </div>
            </div>

            <div class="divider"></div>

            <div class="form-group">
                <h4 style="font-size: 14px; margin-bottom: 12px;">Upload States:</h4>
                <ul class="feature-list">
                    <li>
                        <i class="fa-solid fa-circle-check"></i>
                        <span>Waiting (queued)</span>
                    </li>
                    <li>
                        <i class="fa-solid fa-circle-check"></i>
                        <span>Uploading (in progress)</span>
                    </li>
                    <li>
                        <i class="fa-solid fa-circle-check"></i>
                        <span>Completed (success)</span>
                    </li>
                    <li>
                        <i class="fa-solid fa-circle-check"></i>
                        <span>Error (failed)</span>
                    </li>
                </ul>
            </div>
        </div>
    </div>
</div>

<!-- Upload Form Example -->
<div class="fu-section-title">
    <i class="fa-solid fa-clipboard-list"></i>
    Upload Form Example
</div>

<div class="file-upload-grid full-width">
    <div class="content-card">
        <div class="card-header">
            <div class="card-header-left">
                <div class="card-icon bg-primary">
                    <i class="fa-solid fa-paper-plane"></i>
                </div>
                <div>
                    <h3>Job Application Form</h3>
                    <p class="card-subtitle">Complete form with file uploads</p>
                </div>
            </div>
        </div>
        <div class="card-body">
            <form style="max-width: 800px;">
                <div style="display: grid; grid-template-columns: 1fr 1fr; gap: 20px;">
                    <div class="fu-example">
                        <label class="fu-label">Full Name <span style="color: var(--danger);">*</span></label>
                        <input type="text" class="file-input-label" style="cursor: text;" placeholder="John Doe" required>
                    </div>

                    <div class="fu-example">
                        <label class="fu-label">Email <span style="color: var(--danger);">*</span></label>
                        <input type="email" class="file-input-label" style="cursor: text;" placeholder="john@example.com" required>
                    </div>
                </div>

                <div class="fu-example">
                    <label class="fu-label">
                        Resume/CV
                        <span class="fu-hint">PDF, DOC, DOCX</span>
                    </label>
                    <div class="upload-zone" style="padding: 30px;">
                        <div class="upload-zone-icon" style="font-size: 36px;">
                            <i class="fa-solid fa-file-pdf"></i>
                        </div>
                        <div class="upload-zone-title" style="font-size: 14px;">Upload your resume</div>
                        <div class="upload-zone-desc" style="font-size: 12px;">PDF or Word document, max 5MB</div>
                        <div class="upload-zone-btn" style="padding: 8px 16px; font-size: 13px;">
                            <i class="fa-solid fa-upload"></i>
                            Choose File
                        </div>
                    </div>
                </div>

                <div class="fu-example">
                    <label class="fu-label">
                        Portfolio (Optional)
                        <span class="fu-hint">Images</span>
                    </label>
                    <div class="upload-zone" style="padding: 30px;">
                        <div class="upload-zone-icon" style="font-size: 36px; color: var(--success);">
                            <i class="fa-solid fa-images"></i>
                        </div>
                        <div class="upload-zone-title" style="font-size: 14px;">Upload portfolio samples</div>
                        <div class="upload-zone-desc" style="font-size: 12px;">JPG, PNG up to 3MB each</div>
                        <div class="upload-zone-btn" style="padding: 8px 16px; font-size: 13px; background: var(--success);">
                            <i class="fa-solid fa-upload"></i>
                            Select Images
                        </div>
                    </div>
                </div>

                <div class="fu-example">
                    <label class="fu-label">Cover Letter</label>
                    <div class="file-input-wrapper">
                        <input type="file" id="coverLetter" accept=".pdf,.doc,.docx" style="display: none;">
                        <div class="file-input-label" style="cursor: pointer;">
                            <i class="fa-solid fa-paperclip file-input-icon"></i>
                            <span class="file-input-text">Choose cover letter...</span>
                        </div>
                    </div>
                </div>

                <div class="divider"></div>

                <div style="display: flex; gap: 12px;">
                    <button type="submit" class="btn btn-primary" style="flex: 1;">
                        <i class="fa-solid fa-paper-plane"></i>
                        Submit Application
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
<div class="file-upload-grid full-width">
    <div class="content-card">
        <div class="card-header">
            <div class="card-header-left">
                <div class="card-icon bg-success">
                    <i class="fa-solid fa-code"></i>
                </div>
                <div>
                    <h3>Implementation Guide</h3>
                    <p class="card-subtitle">How to implement file upload with JavaScript</p>
                </div>
            </div>
        </div>
        <div class="card-body">
            <div class="code-block">
                <div style="color: var(--text-tertiary); margin-bottom: 8px;">1. Drag & Drop Handler:</div>
                <code style="color: var(--accent);">
                    const dropZone = document.querySelector('.upload-zone');<br>
                    <br>
                    dropZone.addEventListener('dragover', (e) => {<br>
                    &nbsp;&nbsp;e.preventDefault();<br>
                    &nbsp;&nbsp;dropZone.classList.add('drag-over');<br>
                    });<br>
                    <br>
                    dropZone.addEventListener('dragleave', () => {<br>
                    &nbsp;&nbsp;dropZone.classList.remove('drag-over');<br>
                    });<br>
                    <br>
                    dropZone.addEventListener('drop', (e) => {<br>
                    &nbsp;&nbsp;e.preventDefault();<br>
                    &nbsp;&nbsp;dropZone.classList.remove('drag-over');<br>
                    &nbsp;&nbsp;const files = e.dataTransfer.files;<br>
                    &nbsp;&nbsp;handleFiles(files);<br>
                    });
                </code>
            </div>

            <div class="code-block">
                <div style="color: var(--text-tertiary); margin-bottom: 8px; margin-top: 16px;">2. File Upload with Progress:</div>
                <code style="color: var(--success);">
                    function uploadFile(file) {<br>
                    &nbsp;&nbsp;const formData = new FormData();<br>
                    &nbsp;&nbsp;formData.append('file', file);<br>
                    &nbsp;&nbsp;<br>
                    &nbsp;&nbsp;const xhr = new XMLHttpRequest();<br>
                    &nbsp;&nbsp;xhr.open('POST', '/upload');<br>
                    &nbsp;&nbsp;<br>
                    &nbsp;&nbsp;xhr.upload.addEventListener('progress', (e) => {<br>
                    &nbsp;&nbsp;&nbsp;&nbsp;if (e.lengthComputable) {<br>
                    &nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;const percent = (e.loaded / e.total) * 100;<br>
                    &nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;updateProgressBar(percent);<br>
                    &nbsp;&nbsp;&nbsp;&nbsp;}<br>
                    &nbsp;&nbsp;});<br>
                    &nbsp;&nbsp;<br>
                    &nbsp;&nbsp;xhr.send(formData);<br>
                    }
                </code>
            </div>

            <div class="code-block">
                <div style="color: var(--text-tertiary); margin-bottom: 8px; margin-top: 16px;">3. File Validation:</div>
                <code style="color: var(--warning);">
                    function validateFile(file) {<br>
                    &nbsp;&nbsp;const maxSize = 10 * 1024 * 1024; // 10MB<br>
                    &nbsp;&nbsp;const allowedTypes = ['image/jpeg', 'image/png', 'application/pdf'];<br>
                    &nbsp;&nbsp;<br>
                    &nbsp;&nbsp;if (file.size > maxSize) {<br>
                    &nbsp;&nbsp;&nbsp;&nbsp;alert('File too large');<br>
                    &nbsp;&nbsp;&nbsp;&nbsp;return false;<br>
                    &nbsp;&nbsp;}<br>
                    &nbsp;&nbsp;<br>
                    &nbsp;&nbsp;if (!allowedTypes.includes(file.type)) {<br>
                    &nbsp;&nbsp;&nbsp;&nbsp;alert('Invalid file type');<br>
                    &nbsp;&nbsp;&nbsp;&nbsp;return false;<br>
                    &nbsp;&nbsp;}<br>
                    &nbsp;&nbsp;<br>
                    &nbsp;&nbsp;return true;<br>
                    }
                </code>
            </div>

            <div class="code-block">
                <div style="color: var(--text-tertiary); margin-bottom: 8px; margin-top: 16px;">4. Image Preview:</div>
                <code style="color: var(--info);">
                    function previewImage(file) {<br>
                    &nbsp;&nbsp;const reader = new FileReader();<br>
                    &nbsp;&nbsp;reader.onload = (e) => {<br>
                    &nbsp;&nbsp;&nbsp;&nbsp;const img = document.createElement('img');<br>
                    &nbsp;&nbsp;&nbsp;&nbsp;img.src = e.target.result;<br>
                    &nbsp;&nbsp;&nbsp;&nbsp;img.classList.add('file-preview-image');<br>
                    &nbsp;&nbsp;&nbsp;&nbsp;previewContainer.appendChild(img);<br>
                    &nbsp;&nbsp;};<br>
                    &nbsp;&nbsp;reader.readAsDataURL(file);<br>
                    }
                </code>
            </div>

            <div class="divider"></div>

            <div class="form-group">
                <h4 style="font-size: 14px; margin-bottom: 12px;">Popular Libraries:</h4>
                <ul class="feature-list">
                    <li>
                        <i class="fa-solid fa-circle-check"></i>
                        <span><strong>Dropzone.js</strong> - Drag & drop library</span>
                    </li>
                    <li>
                        <i class="fa-solid fa-circle-check"></i>
                        <span><strong>FilePond</strong> - Advanced file input</span>
                    </li>
                    <li>
                        <i class="fa-solid fa-circle-check"></i>
                        <span><strong>Uppy</strong> - Modern uploader</span>
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
    // Initialize all upload zones
    let isTriggeringClick = false;
    
    $('.upload-zone').each(function() {
        const $zone = $(this);
        
        // Create hidden file input if not exists
        let $input = $zone.find('input[type="file"]');
        if ($input.length === 0) {
            $input = $('<input type="file" style="display: none;">');
            $zone.append($input);
        }
        
        // Click to browse
        $zone.on('click', function(e) {
            // Don't trigger if clicking remove buttons
            if ($(e.target).closest('.file-item-remove, .file-preview-remove').length) {
                return;
            }
            
            // Prevent infinite loop
            if (isTriggeringClick) return;
            isTriggeringClick = true;
            
            $input.trigger('click');
            
            // Reset flag after a short delay
            setTimeout(() => {
                isTriggeringClick = false;
            }, 100);
        });
    });
    
    // Global drag and drop handlers
    $(document).on('dragover dragenter', '.upload-zone', function(e) {
        e.preventDefault();
        e.stopPropagation();
        $(this).addClass('drag-over');
    });
    
    $(document).on('dragleave', '.upload-zone', function(e) {
        e.preventDefault();
        e.stopPropagation();
        // Only remove if actually leaving the zone
        if (!$(this).is(e.relatedTarget) && !$(this).has(e.relatedTarget).length) {
            $(this).removeClass('drag-over');
        }
    });
    
    $(document).on('drop', '.upload-zone', function(e) {
        e.preventDefault();
        e.stopPropagation();
        $(this).removeClass('drag-over');
        
        const files = e.originalEvent.dataTransfer.files;
        if (files && files.length > 0) {
            handleFiles(files, $(this));
        }
    });
    
    // File input change (delegated)
    $(document).on('change', '.upload-zone input[type="file"]', function() {
        const $zone = $(this).closest('.upload-zone');
        const files = this.files;
        if (files && files.length > 0) {
            handleFiles(files, $zone);
            // Reset input so same file can be selected again
            $(this).val('');
        }
    });
    
    // Handle file input change (delegated)
    $(document).on('change', '.file-input-wrapper input[type="file"]', function() {
        console.log('File input changed:', this.id || 'unnamed');
        
        const $input = $(this);
        const $label = $input.siblings('.file-input-label');
        const files = this.files;
        
        console.log('Files selected:', files.length);
        
        if (files.length > 0) {
            // Handle multiple files
            if (files.length > 1) {
                const fileNames = Array.from(files).map(f => f.name).join(', ');
                const totalSize = Array.from(files).reduce((sum, f) => sum + f.size, 0);
                
                $label.html(`
                    <i class="fa-solid fa-folder-open file-input-icon"></i>
                    <div>
                        <div class="file-input-name">${files.length} files selected</div>
                        <div class="file-input-text">${formatFileSize(totalSize)} • ${escapeHtml(fileNames.substring(0, 50))}${fileNames.length > 50 ? '...' : ''}</div>
                    </div>
                `);
                
                showToast(`${files.length} files selected`, 'success');
            } else {
                // Handle single file
                const file = files[0];
                const fileName = file.name;
                const fileSize = formatFileSize(file.size);
                const fileType = getFileExtension(fileName);
                
                $label.html(`
                    <i class="fa-solid fa-file file-input-icon"></i>
                    <div>
                        <div class="file-input-name">${escapeHtml(fileName)}</div>
                        <div class="file-input-text">${fileSize} • ${fileType.toUpperCase()}</div>
                    </div>
                `);
                
                // Show success toast
                showToast('File selected: ' + fileName, 'success');
            }
        } else {
            // Reset to default state
            const isMultiple = $input.attr('multiple');
            
            if (isMultiple) {
                $label.html(`
                    <i class="fa-solid fa-folder-open file-input-icon"></i>
                    <span class="file-input-text">Choose files...</span>
                `);
            } else {
                $label.html(`
                    <i class="fa-solid fa-paperclip file-input-icon"></i>
                    <span class="file-input-text">Choose a file...</span>
                `);
            }
        }
    });
    
    console.log('File input handler attached');
    
    // Make file input label clickable
    $(document).on('click', '.file-input-label', function(e) {
        e.preventDefault();
        e.stopPropagation();
        
        console.log('File label clicked');
        const $label = $(this);
        const $input = $label.siblings('input[type="file"]');
        
        if ($input.length > 0) {
            console.log('Triggering file input click');
            $input.trigger('click');
        }
    });
    
    // Handle dropped/selected files
    function handleFiles(files, $zone) {
        if (!files || files.length === 0) {
            console.log('No files to handle');
            return;
        }
        
        console.log('Handling', files.length, 'file(s)');
        
        Array.from(files).forEach(file => {
            console.log('Processing file:', file.name, file.size, file.type);
            
            // Validate file
            if (!validateFile(file)) {
                console.log('File validation failed:', file.name);
                return;
            }
            
            console.log('File validated:', file.name);
            
            // Show success message
            showToast('File added: ' + file.name, 'success');
            
            // Simulate upload progress
            simulateUpload(file);
        });
    }
    
    // Validate file
    function validateFile(file) {
        const maxSize = 10 * 1024 * 1024; // 10MB
        const allowedTypes = [
            'image/jpeg', 'image/png', 'image/gif',
            'application/pdf', 'application/msword',
            'application/vnd.openxmlformats-officedocument.wordprocessingml.document',
            'application/vnd.ms-excel',
            'application/vnd.openxmlformats-officedocument.spreadsheetml.sheet',
            'text/plain'
        ];
        
        // Check file size
        if (file.size > maxSize) {
            Swal.fire({
                toast: true,
                position: 'top-end',
                icon: 'error',
                title: 'File too large! Max 10MB',
                showConfirmButton: false,
                timer: 2500
            });
            return false;
        }
        
        // Check file type
        if (!allowedTypes.includes(file.type)) {
            Swal.fire({
                toast: true,
                position: 'top-end',
                icon: 'warning',
                title: 'Invalid file type!',
                showConfirmButton: false,
                timer: 2500
            });
            return false;
        }
        
        return true;
    }
    
    // Simulate upload progress
    function simulateUpload(file) {
        const fileName = file.name;
        const fileSize = formatFileSize(file.size);
        const fileType = getFileExtension(fileName);
        
        // Create file item with progress
        const $fileItem = $(`
            <div class="file-item upload-item-${Date.now()}">
                <div class="file-item-icon">${getFileIcon(fileType)}</div>
                <div class="file-item-content">
                    <div class="file-item-name">${escapeHtml(fileName)}</div>
                    <div class="file-item-size">${fileSize} • Uploading...</div>
                    <div class="upload-progress">
                        <div class="progress-bar-bg">
                            <div class="progress-bar-fill" style="width: 0%;"></div>
                        </div>
                        <div class="progress-text">
                            <span>0 MB / ${fileSize}</span>
                            <span>0%</span>
                        </div>
                    </div>
                </div>
                <button class="file-item-remove">
                    <i class="fa-solid fa-xmark"></i>
                </button>
            </div>
        `);
        
        // Add to first upload zone example
        const $firstExample = $('.fu-example').first();
        if ($firstExample.find('.file-list').length === 0) {
            $firstExample.append('<div class="file-list"></div>');
        }
        $firstExample.find('.file-list').append($fileItem);
        
        // Simulate progress
        let progress = 0;
        const interval = setInterval(() => {
            progress += Math.random() * 15;
            
            if (progress >= 100) {
                progress = 100;
                clearInterval(interval);
                
                // Update to complete
                $fileItem.find('.file-item-size').text(`${fileSize} • Upload complete`);
                $fileItem.find('.progress-bar-fill').addClass('success');
                $fileItem.find('.progress-text').html(`
                    <span>Upload successful</span>
                    <span style="color: var(--success);">✓</span>
                `);
                $fileItem.find('.file-item-icon').text('✅');
                
                showToast('Upload complete: ' + fileName, 'success');
            } else {
                const uploadedSize = (progress / 100 * file.size);
                $fileItem.find('.file-item-size').text(`${fileSize} • ${Math.round(progress)}% uploaded`);
                $fileItem.find('.progress-bar-fill').css('width', progress + '%');
                $fileItem.find('.progress-text').html(`
                    <span>${formatFileSize(uploadedSize)} / ${fileSize}</span>
                    <span>${Math.round(progress)}%</span>
                `);
            }
        }, 300);
    }
    
    // Remove file from list
    $(document).on('click', '.file-item-remove', function(e) {
        e.preventDefault();
        e.stopPropagation();
        
        const $item = $(this).closest('.file-item');
        const fileName = $item.find('.file-item-name').text();
        
        Swal.fire({
            title: 'Remove file?',
            text: `Remove "${fileName}" from list?`,
            icon: 'question',
            showCancelButton: true,
            confirmButtonColor: 'var(--danger)',
            cancelButtonColor: 'var(--text-tertiary)',
            confirmButtonText: 'Yes, remove',
            cancelButtonText: 'Cancel'
        }).then((result) => {
            if (result.isConfirmed) {
                $item.fadeOut(300, function() {
                    $(this).remove();
                    showToast('File removed: ' + fileName, 'info');
                });
            }
        });
    });
    
    // Remove file from preview
    $(document).on('click', '.file-preview-remove', function(e) {
        e.preventDefault();
        e.stopPropagation();
        
        const $item = $(this).closest('.file-preview-item');
        
        $item.fadeOut(300, function() {
            $(this).remove();
            showToast('Image removed', 'info');
        });
    });
    
    // Helper functions
    function formatFileSize(bytes) {
        if (bytes === 0) return '0 Bytes';
        const k = 1024;
        const sizes = ['Bytes', 'KB', 'MB', 'GB'];
        const i = Math.floor(Math.log(bytes) / Math.log(k));
        return Math.round(bytes / Math.pow(k, i) * 100) / 100 + ' ' + sizes[i];
    }
    
    function getFileExtension(filename) {
        return filename.split('.').pop().toLowerCase();
    }
    
    function getFileIcon(type) {
        const icons = {
            'pdf': '📄',
            'doc': '📝',
            'docx': '📝',
            'xls': '📊',
            'xlsx': '📊',
            'jpg': '🖼️',
            'jpeg': '🖼️',
            'png': '🖼️',
            'gif': '🖼️',
            'txt': '📃',
            'zip': '📦',
            'mp4': '🎥'
        };
        return icons[type] || '📄';
    }
    
    function escapeHtml(text) {
        const div = document.createElement('div');
        div.textContent = text;
        return div.innerHTML;
    }
    

    
    // Page header buttons
    $('.page-header .btn-primary').on('click', function() {
        $('.upload-zone').first().trigger('click');
    });
    
    $('.page-header .btn-secondary').on('click', function() {
        Swal.fire({
            icon: 'info',
            title: 'File Upload Documentation',
            html: `
                <div style="text-align: left; font-size: 13px;">
                    <p><strong>File Upload</strong> features include:</p>
                    <ul>
                        <li>Drag & drop support</li>
                        <li>File validation</li>
                        <li>Progress tracking</li>
                        <li>Image preview</li>
                        <li>Multiple file upload</li>
                    </ul>
                </div>
            `,
            confirmButtonText: 'OK',
            confirmButtonColor: 'var(--accent)'
        });
    });
    
    // Form submission
    $('form').on('submit', function(e) {
        e.preventDefault();
        
        const $form = $(this);
        const formData = new FormData();
        
        // Collect form data
        $form.find('input[type="text"], input[type="email"]').each(function() {
            const value = $(this).val().trim();
            if (value) {
                formData.append($(this).attr('placeholder') || 'field', value);
            }
        });
        
        // Validate required fields
        const requiredFields = $form.find('[required]');
        let allFilled = true;
        
        requiredFields.each(function() {
            if (!$(this).val().trim()) {
                allFilled = false;
                $(this).focus();
                return false;
            }
        });
        
        if (!allFilled) {
            Swal.fire({
                toast: true,
                position: 'top-end',
                icon: 'warning',
                title: 'Please fill in all required fields',
                showConfirmButton: false,
                timer: 2500
            });
            return;
        }
        
        // Show success
        Swal.fire({
            icon: 'success',
            title: 'Application Submitted!',
            text: 'Your job application has been submitted successfully.',
            confirmButtonText: 'OK',
            confirmButtonColor: 'var(--accent)'
        });
        
        console.log('Form Data:', formData);
    });
    
    // Form reset
    $('form').on('reset', function(e) {
        e.preventDefault();
        
        // Reset file inputs
        $(this).find('input[type="file"]').val('');
        
        // Reset file input labels
        $(this).find('.file-input-label').html(`
            <i class="fa-solid fa-paperclip file-input-icon"></i>
            <span class="file-input-text">Choose file...</span>
        `);
        
        Swal.fire({
            toast: true,
            position: 'top-end',
            icon: 'info',
            title: 'Form reset',
            showConfirmButton: false,
            timer: 1500
        });
    });
});
</script>
@endpush
