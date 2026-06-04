/**
 * Form Validation JavaScript
 * Real-time validation with feedback for all form types
 */

document.addEventListener('DOMContentLoaded', function() {
    // Initialize all validations
    initRequiredValidation();
    initEmailValidation();
    initPasswordValidation();
    initPatternValidation();
    initLengthValidation();
    initNumberValidation();
    initDateValidation();
    initFileValidation();
    initFormSubmission();
});

/**
 * Required Field Validation
 */
function initRequiredValidation() {
    const requiredInputs = document.querySelectorAll('[required]');
    
    requiredInputs.forEach(input => {
        // Validate on blur
        input.addEventListener('blur', function() {
            validateRequiredField(this);
        });
        
        // Validate on input (after first blur)
        input.addEventListener('input', function() {
            if (this.classList.contains('is-invalid') || this.classList.contains('is-valid')) {
                validateRequiredField(this);
            }
        });
    });
}

function validateRequiredField(input) {
    const value = input.value.trim();
    const feedback = input.parentElement.querySelector('.validation-feedback');
    
    if (!value) {
        setInvalidState(input, 'Field ini wajib diisi', feedback);
        return false;
    } else {
        setValidState(input, 'Looks good!', feedback);
        return true;
    }
}

/**
 * Email Validation
 */
function initEmailValidation() {
    const emailInputs = document.querySelectorAll('input[type="email"]');
    
    emailInputs.forEach(input => {
        input.addEventListener('blur', function() {
            validateEmail(this);
        });
        
        input.addEventListener('input', function() {
            if (this.classList.contains('is-invalid') || this.classList.contains('is-valid')) {
                validateEmail(this);
            }
        });
    });
}

function validateEmail(input) {
    const value = input.value.trim();
    const feedback = input.parentElement.querySelector('.validation-feedback');
    
    if (!value) {
        if (input.hasAttribute('required')) {
            setInvalidState(input, 'Email wajib diisi', feedback);
        }
        return false;
    }
    
    const emailPattern = /^[a-z0-9._%+-]+@[a-z0-9.-]+\.[a-z]{2,}$/i;
    
    if (!emailPattern.test(value)) {
        setInvalidState(input, 'Format email tidak valid', feedback);
        return false;
    }
    
    setValidState(input, 'Email valid', feedback);
    return true;
}

/**
 * Password Validation with Strength Meter
 */
function initPasswordValidation() {
    const passwordInputs = document.querySelectorAll('input[type="password"]');
    
    passwordInputs.forEach(input => {
        input.addEventListener('input', function() {
            validatePasswordStrength(this);
        });
    });
    
    // Confirm password validation
    validateConfirmPassword();
}

function validatePasswordStrength(input) {
    const value = input.value;
    const strengthMeter = input.parentElement.querySelector('.strength-meter-fill');
    const strengthText = input.parentElement.querySelector('.strength-text');
    const requirements = input.parentElement.querySelectorAll('.requirement-list li');
    
    if (!value) {
        if (strengthMeter) strengthMeter.className = 'strength-meter-fill';
        if (strengthText) strengthText.textContent = '';
        return;
    }
    
    let strength = 0;
    const checks = {
        minLength: value.length >= 8,
        hasUppercase: /[A-Z]/.test(value),
        hasNumber: /\d/.test(value),
        hasSpecial: /[!@#$%^&*(),.?":{}|<>]/.test(value)
    };
    
    // Update requirement list
    if (requirements.length >= 4) {
        updateRequirement(requirements[0], checks.minLength, `Minimum 8 characters (${value.length}/8)`);
        updateRequirement(requirements[1], checks.hasUppercase, 'At least one uppercase letter');
        updateRequirement(requirements[2], checks.hasNumber, 'At least one number');
        updateRequirement(requirements[3], checks.hasSpecial, 'At least one special character');
    }
    
    // Calculate strength
    if (checks.minLength) strength++;
    if (checks.hasUppercase) strength++;
    if (checks.hasNumber) strength++;
    if (checks.hasSpecial) strength++;
    
    // Update strength meter
    if (strengthMeter) {
        strengthMeter.className = 'strength-meter-fill';
        
        if (strength <= 1) {
            strengthMeter.classList.add('strength-weak');
            if (strengthText) {
                strengthText.textContent = 'Weak - Add uppercase, numbers, and special characters';
                strengthText.style.color = 'var(--danger)';
            }
        } else if (strength === 2 || strength === 3) {
            strengthMeter.classList.add('strength-medium');
            if (strengthText) {
                strengthText.textContent = 'Medium - Good, but could be stronger';
                strengthText.style.color = 'var(--warning)';
            }
        } else if (strength === 4) {
            strengthMeter.classList.add('strength-strong');
            if (strengthText) {
                strengthText.textContent = 'Strong - Excellent password!';
                strengthText.style.color = 'var(--success)';
            }
        }
    }
}

function updateRequirement(element, met, text) {
    element.className = met ? 'met' : 'unmet';
    const icon = element.querySelector('i');
    if (icon) {
        icon.className = met ? 'fa-solid fa-circle-check' : 'fa-solid fa-circle-xmark';
    }
    const span = element.querySelector('span');
    if (span) span.textContent = text;
}

function validateConfirmPassword() {
    const passwordInput = document.getElementById('passwordInput');
    const confirmPasswordInputs = document.querySelectorAll('input[placeholder*="Confirm"], input[placeholder*="confirm"]');
    
    if (passwordInput) {
        confirmPasswordInputs.forEach(confirmInput => {
            confirmInput.addEventListener('input', function() {
                const feedback = this.parentElement.querySelector('.validation-feedback');
                
                if (this.value !== passwordInput.value) {
                    setInvalidState(this, 'Passwords do not match', feedback);
                } else {
                    setValidState(this, 'Passwords match', feedback);
                }
            });
        });
    }
}

/**
 * Pattern Validation (Phone, URL, etc.)
 */
function initPatternValidation() {
    const patternInputs = document.querySelectorAll('input[pattern]');
    
    patternInputs.forEach(input => {
        input.addEventListener('blur', function() {
            validatePattern(this);
        });
        
        input.addEventListener('input', function() {
            if (this.classList.contains('is-invalid') || this.classList.contains('is-valid')) {
                validatePattern(this);
            }
        });
    });
}

function validatePattern(input) {
    const value = input.value.trim();
    const feedback = input.parentElement.querySelector('.validation-feedback');
    const pattern = new RegExp(input.getAttribute('pattern'));
    
    if (!value) {
        if (input.hasAttribute('required')) {
            setInvalidState(input, 'Field ini wajib diisi', feedback);
        }
        return false;
    }
    
    if (!pattern.test(value)) {
        let message = 'Format tidak valid';
        
        // Custom messages based on input type
        if (input.type === 'tel') {
            message = 'Format: +62xxxxxxxxxx';
        } else if (input.type === 'url') {
            message = 'URL harus dimulai dengan http:// atau https://';
        }
        
        setInvalidState(input, message, feedback);
        return false;
    }
    
    setValidState(input, 'Format valid', feedback);
    return true;
}

/**
 * Length Validation (minlength/maxlength)
 */
function initLengthValidation() {
    const lengthInputs = document.querySelectorAll('input[minlength], input[maxlength], textarea[minlength], textarea[maxlength]');
    
    lengthInputs.forEach(input => {
        input.addEventListener('input', function() {
            validateLength(this);
        });
    });
}

function validateLength(input) {
    const value = input.value;
    const length = value.length;
    const minLength = input.getAttribute('minlength');
    const maxLength = input.getAttribute('maxlength');
    const feedback = input.parentElement.querySelector('.validation-feedback');
    
    if (minLength && length < parseInt(minLength)) {
        const maxText = maxLength ? `/${maxLength}` : '';
        setInvalidState(input, `Minimum ${minLength} characters (${length}${maxText})`, feedback);
        return false;
    }
    
    if (maxLength && length > parseInt(maxLength)) {
        setInvalidState(input, `Maximum ${maxLength} characters allowed`, feedback);
        return false;
    }
    
    if (maxLength) {
        setValidState(input, `${length}/${maxLength} characters`, feedback);
    }
    
    return true;
}

/**
 * Number Range Validation
 */
function initNumberValidation() {
    const numberInputs = document.querySelectorAll('input[type="number"]');
    
    numberInputs.forEach(input => {
        input.addEventListener('blur', function() {
            validateNumberRange(this);
        });
        
        input.addEventListener('input', function() {
            if (this.classList.contains('is-invalid') || this.classList.contains('is-valid')) {
                validateNumberRange(this);
            }
        });
    });
}

function validateNumberRange(input) {
    const value = parseFloat(input.value);
    const feedback = input.parentElement.querySelector('.validation-feedback');
    const min = parseFloat(input.getAttribute('min'));
    const max = parseFloat(input.getAttribute('max'));
    
    if (isNaN(value)) {
        if (input.hasAttribute('required')) {
            setInvalidState(input, 'Angka wajib diisi', feedback);
        }
        return false;
    }
    
    if (!isNaN(min) && value < min) {
        setInvalidState(input, `Nilai minimum adalah ${min}`, feedback);
        return false;
    }
    
    if (!isNaN(max) && value > max) {
        setInvalidState(input, `Nilai maksimum adalah ${max}`, feedback);
        return false;
    }
    
    setValidState(input, 'Nilai valid', feedback);
    return true;
}

/**
 * Date Validation
 */
function initDateValidation() {
    const dateInputs = document.querySelectorAll('input[type="date"]');
    
    dateInputs.forEach(input => {
        input.addEventListener('blur', function() {
            validateDate(this);
        });
        
        input.addEventListener('change', function() {
            validateDate(this);
        });
    });
}

function validateDate(input) {
    const value = input.value;
    const feedback = input.parentElement.querySelector('.validation-feedback');
    const minDate = input.getAttribute('min');
    const maxDate = input.getAttribute('max');
    
    if (!value) {
        if (input.hasAttribute('required')) {
            setInvalidState(input, 'Tanggal wajib diisi', feedback);
        }
        return false;
    }
    
    const date = new Date(value);
    
    if (minDate && date < new Date(minDate)) {
        setInvalidState(input, `Tanggal tidak boleh sebelum ${formatDate(minDate)}`, feedback);
        return false;
    }
    
    if (maxDate && date > new Date(maxDate)) {
        setInvalidState(input, `Tanggal tidak boleh setelah ${formatDate(maxDate)}`, feedback);
        return false;
    }
    
    setValidState(input, 'Tanggal valid', feedback);
    return true;
}

function formatDate(dateString) {
    const date = new Date(dateString);
    return date.toLocaleDateString('id-ID', { day: 'numeric', month: 'long', year: 'numeric' });
}

/**
 * File Upload Validation
 */
function initFileValidation() {
    const fileInputs = document.querySelectorAll('input[type="file"]');
    
    fileInputs.forEach(input => {
        input.addEventListener('change', function() {
            validateFile(this);
        });
    });
}

function validateFile(input) {
    const file = input.files[0];
    const feedback = input.parentElement.querySelector('.validation-feedback');
    
    if (!file) {
        if (input.hasAttribute('required')) {
            setInvalidState(input, 'File wajib dipilih', feedback);
        }
        return false;
    }
    
    // Check file type
    const accept = input.getAttribute('accept');
    if (accept) {
        const acceptedTypes = accept.split(',').map(type => type.trim());
        const isValidType = acceptedTypes.some(type => {
            if (type.startsWith('.')) {
                return file.name.endsWith(type);
            }
            return file.type.match(type.replace('*', '.*'));
        });
        
        if (!isValidType) {
            setInvalidState(input, 'Tipe file tidak diizinkan', feedback);
            return false;
        }
    }
    
    // Check file size (if max attribute exists)
    const maxSize = input.getAttribute('max');
    if (maxSize && file.size > parseInt(maxSize)) {
        const sizeMB = (parseInt(maxSize) / 1024 / 1024).toFixed(1);
        setInvalidState(input, `Ukuran file melebihi batas maksimum (${sizeMB}MB)`, feedback);
        return false;
    }
    
    const sizeKB = (file.size / 1024).toFixed(1);
    setValidState(input, `File valid (${sizeKB}KB)`, feedback);
    return true;
}

/**
 * Form Submission Handler
 */
function initFormSubmission() {
    const forms = document.querySelectorAll('form[id$="Form"]');
    
    forms.forEach(form => {
        form.addEventListener('submit', function(e) {
            e.preventDefault();
            
            const isValid = validateForm(this);
            
            if (isValid) {
                // Show success message
                showNotification('success', 'Form berhasil divalidasi!');
                console.log('Form submitted successfully');
            } else {
                // Show error message
                showNotification('error', 'Mohon periksa kembali form Anda');
                
                // Scroll to first error
                const firstError = this.querySelector('.is-invalid');
                if (firstError) {
                    firstError.scrollIntoView({ behavior: 'smooth', block: 'center' });
                    firstError.focus();
                }
            }
        });
        
        // Reset handler
        form.addEventListener('reset', function() {
            setTimeout(() => {
                clearValidationStates(this);
            }, 0);
        });
    });
}

function validateForm(form) {
    let isValid = true;
    const inputs = form.querySelectorAll('input, select, textarea');
    
    inputs.forEach(input => {
        if (input.type !== 'submit' && input.type !== 'reset' && input.type !== 'button') {
            if (input.hasAttribute('required') && !input.value.trim()) {
                const feedback = input.parentElement.querySelector('.validation-feedback');
                setInvalidState(input, 'Field ini wajib diisi', feedback);
                isValid = false;
            }
            
            if (input.type === 'email' && input.value) {
                if (!validateEmail(input)) isValid = false;
            }
            
            if (input.type === 'file' && input.files.length > 0) {
                if (!validateFile(input)) isValid = false;
            }
        }
    });
    
    return isValid;
}

function clearValidationStates(form) {
    const inputs = form.querySelectorAll('.is-valid, .is-invalid');
    inputs.forEach(input => {
        input.classList.remove('is-valid', 'is-invalid');
    });
    
    const feedbacks = form.querySelectorAll('.validation-feedback');
    feedbacks.forEach(feedback => {
        feedback.style.display = 'none';
    });
}

/**
 * Helper Functions
 */
function setInvalidState(input, message, feedback) {
    input.classList.remove('is-valid');
    input.classList.add('is-invalid');
    
    if (feedback) {
        feedback.classList.remove('valid-feedback', 'warning-feedback');
        feedback.classList.add('invalid-feedback');
        feedback.style.display = 'block';
        
        const messageSpan = feedback.querySelector('span');
        if (messageSpan) {
            messageSpan.textContent = message;
        }
    }
}

function setValidState(input, message, feedback) {
    input.classList.remove('is-invalid');
    input.classList.add('is-valid');
    
    if (feedback) {
        feedback.classList.remove('invalid-feedback', 'warning-feedback');
        feedback.classList.add('valid-feedback');
        feedback.style.display = 'block';
        
        const messageSpan = feedback.querySelector('span');
        if (messageSpan) {
            messageSpan.textContent = message;
        }
    }
}

function showNotification(type, message) {
    // You can integrate with SweetAlert or your notification system
    console.log(`[${type.toUpperCase()}] ${message}`);
    
    // Example with SweetAlert (if available):
    if (typeof Swal !== 'undefined') {
        Swal.fire({
            icon: type === 'success' ? 'success' : 'error',
            title: message,
            toast: true,
            position: 'top-end',
            showConfirmButton: false,
            timer: 3000
        });
    }
}
