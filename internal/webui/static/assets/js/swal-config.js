/**
 * Custom Toast Notification System
 * Menggunakan custom toast dengan animasi slide-in
 * Mengganti SweetAlert2 dan Metro UI Toast
 */

// Toast container
let toastContainer = null;

// Toast configuration
const toastConfig = {
    timer: 4000,
    success: {
        icon: '✓',
        text: 'Success'
    },
    error: {
        icon: '✕',
        text: 'Error'
    },
    warning: {
        icon: '⚠',
        text: 'Warning'
    },
    info: {
        icon: 'ℹ',
        text: 'Info'
    }
};

// Initialize toast container
function initToastContainer() {
    if (!toastContainer) {
        toastContainer = document.createElement('ul');
        toastContainer.className = 'custom-notifications';
        document.body.appendChild(toastContainer);
    }
    return toastContainer;
}

// Remove toast
function removeToast(toast) {
    toast.classList.add('hide');
    if (toast.timeoutId) clearTimeout(toast.timeoutId);
    setTimeout(() => toast.remove(), 300);
}

// Create toast
function createToast(type, message) {
    const container = initToastContainer();
    const config = toastConfig[type] || toastConfig.info;
    
    const toast = document.createElement('li');
    toast.className = `custom-toast ${type}`;
    toast.innerHTML = `
        <div class="toast-column">
            <span class="toast-icon">${config.icon}</span>
            <span class="toast-message">${message}</span>
        </div>
        <span class="toast-close" onclick="removeToast(this.parentElement)">×</span>
    `;
    
    container.appendChild(toast);
    toast.timeoutId = setTimeout(() => removeToast(toast), toastConfig.timer);
    
    return toast;
}

// ========================
// Toast helper
// ========================
function showToast(message, type = 'success', position = 'top-right') {
    // Position handling (will be controlled by CSS)
    const container = initToastContainer();
    container.className = `custom-notifications position-${position}`;
    
    return createToast(type, message);
}


// ========================
// Dialog helpers (SweetAlert2)
// ========================
// Global mixin: ensure SweetAlert always appears above modals (z-index: 9999)
const SwalDefaults = {
    customClass: { popup: 'swal-above-modal' },
    didOpen: function() {
        // Force z-index above modal overlay
        const popup = this.getPopup();
        if (popup) popup.style.zIndex = '10100';
        const backdrop = document.querySelector('.swal2-backdrop-show');
        if (backdrop) backdrop.style.zIndex = '10099';
    }
};

function showConfirm(title, text = '', confirmText = 'Ya', cancelText = 'Batal') {
    return Swal.fire({
        ...SwalDefaults,
        icon: 'warning',
        title: title,
        text: text,
        showCancelButton: true,
        confirmButtonColor: '#E81123',
        cancelButtonColor: '#8a8a8a',
        confirmButtonText: confirmText,
        cancelButtonText: cancelText
    });
}

function showSuccess(title, text = '', timer = 2000) {
    return Swal.fire({
        ...SwalDefaults,
        icon: 'success',
        title: title,
        text: text,
        timer: timer,
        showConfirmButton: false
    });
}

function showError(title, text = '') {
    return Swal.fire({
        ...SwalDefaults,
        icon: 'error',
        title: title,
        text: text
    });
}

function showWarning(title, text = '') {
    return Swal.fire({
        ...SwalDefaults,
        icon: 'warning',
        title: title,
        text: text
    });
}

function showInfo(title, text = '') {
    return Swal.fire({
        ...SwalDefaults,
        icon: 'info',
        title: title,
        text: text
    });
}

function showLoading(text = 'Loading...') {
    return Swal.fire({
        ...SwalDefaults,
        title: text,
        allowOutsideClick: false,
        allowEscapeKey: false,
        showConfirmButton: false,
        willOpen: () => {
            Swal.showLoading();
        }
    });
}

function closeLoading() {
    Swal.close();
}

function showLoadingSuccess(text = 'Loading...', successTitle, successText, timer = 1500) {
    return Swal.fire({
        title: text,
        allowOutsideClick: false,
        allowEscapeKey: false,
        showConfirmButton: false,
        willOpen: () => {
            Swal.showLoading();
        },
        didOpen: () => {
            setTimeout(() => {
                Swal.fire({
                    icon: 'success',
                    title: successTitle,
                    text: successText,
                    timer: timer,
                    showConfirmButton: false
                });
            }, 1500);
        }
    });
}


