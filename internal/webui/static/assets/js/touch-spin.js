/* ============================================
   Touch Spin Component Logic
   Number input with increment/decrement buttons
   ============================================ */

'use strict';

const TouchSpin = {
    // Hold timers
    holdInterval: null,
    holdTimeout: null,

    // Initialize
    init() {
        this.bindEvents();
    },

    // Bind all events
    bindEvents() {
        // Increment/decrement button - use mousedown only to avoid double-click
        $(document).on('mousedown', '.touch-spin .spin-btn', (e) => {
            e.preventDefault();
            const btn = $(e.currentTarget);
            const touchSpin = btn.closest('.touch-spin');
            
            // Check if button is first (decrement) or last (increment)
            if (btn.is(':first-child')) {
                this.decrement(touchSpin);
            } else {
                this.increment(touchSpin);
            }
            
            // Start hold for continuous increment/decrement
            this.startHold(btn);
        });

        $(document).on('mouseup mouseleave', '.touch-spin .spin-btn', () => {
            this.stopHold();
        });

        // Keyboard navigation (arrow keys)
        $(document).on('keydown', '.touch-spin .spin-input', (e) => {
            const input = $(e.currentTarget);
            const touchSpin = input.closest('.touch-spin');
            
            if (e.key === 'ArrowUp') {
                e.preventDefault();
                this.increment(touchSpin);
            } else if (e.key === 'ArrowDown') {
                e.preventDefault();
                this.decrement(touchSpin);
            }
        });

        // Manual input validation
        $(document).on('change', '.touch-spin .spin-input', (e) => {
            this.validateInput($(e.currentTarget));
        });
    },

    // Increment value
    increment(touchSpin) {
        const input = touchSpin.find('.spin-input');
        
        if (input.prop('disabled')) return;
        
        const currentValue = parseFloat(input.val()) || 0;
        const step = parseFloat(input.attr('step')) || 1;
        const max = parseFloat(input.attr('max'));
        
        const newValue = currentValue + step;
        
        // Check max limit
        if (max !== undefined && newValue > max) {
            input.val(max);
            this.showWarning(input, 'Maximum value reached');
        } else {
            const precision = this.getDecimalPrecision(step);
            input.val(newValue.toFixed(precision));
            this.updatePreview(touchSpin);
        }
    },

    // Decrement value
    decrement(touchSpin) {
        const input = touchSpin.find('.spin-input');
        
        if (input.prop('disabled')) return;
        
        const currentValue = parseFloat(input.val()) || 0;
        const step = parseFloat(input.attr('step')) || 1;
        const min = parseFloat(input.attr('min'));
        
        const newValue = currentValue - step;
        
        // Check min limit
        if (min !== undefined && newValue < min) {
            input.val(min);
            this.showWarning(input, 'Minimum value reached');
        } else {
            const precision = this.getDecimalPrecision(step);
            input.val(newValue.toFixed(precision));
            this.updatePreview(touchSpin);
        }
    },

    // Start holding button
    startHold(btn) {
        const touchSpin = btn.closest('.touch-spin');
        const isFirstBtn = btn.is(':first-child');
        
        // Start hold timer - after 500ms, start repeating
        this.holdTimeout = setTimeout(() => {
            this.holdInterval = setInterval(() => {
                if (isFirstBtn) {
                    this.decrement(touchSpin);
                } else {
                    this.increment(touchSpin);
                }
            }, 100);
        }, 500);
    },

    // Stop holding button
    stopHold() {
        clearTimeout(this.holdTimeout);
        clearInterval(this.holdInterval);
    },

    // Validate manual input
    validateInput(input) {
        const touchSpin = input.closest('.touch-spin');
        
        let value = parseFloat(input.val());
        const min = parseFloat(input.attr('min'));
        const max = parseFloat(input.attr('max'));
        const step = parseFloat(input.attr('step')) || 1;
        
        // Validate against min/max
        if (!isNaN(min) && value < min) {
            value = min;
            this.showWarning(input, 'Value adjusted to minimum');
        }
        
        if (!isNaN(max) && value > max) {
            value = max;
            this.showWarning(input, 'Value adjusted to maximum');
        }
        
        // Round to step precision
        const precision = this.getDecimalPrecision(step);
        input.val(value.toFixed(precision));
        
        this.updatePreview(touchSpin);
    },

    // Get decimal precision from step
    getDecimalPrecision(step) {
        return step % 1 !== 0 ? (step.toString().split('.')[1] || '').length : 0;
    },

    // Update live preview (if exists)
    updatePreview(touchSpin) {
        // Check if we're in the first card with live preview
        const card = touchSpin.closest('.card-body');
        const previewValue = card.find('.spin-preview-value');
        
        if (previewValue.length && touchSpin.index() === 0) {
            const input = touchSpin.find('.spin-input');
            const value = input.val();
            previewValue.text(value + ' items');
        }
    },

    // Show warning animation
    showWarning(input, message) {
        const touchSpin = input.closest('.touch-spin');
        
        // Shake animation
        touchSpin.css('animation', 'shake 0.3s');
        setTimeout(() => {
            touchSpin.css('animation', '');
        }, 300);
        
        // Flash input background
        input.css('background', 'var(--warning-bg)');
        setTimeout(() => {
            input.css('background', '');
        }, 500);
        
        // Log warning (optional)
        if (typeof MetroAdmin !== 'undefined' && MetroAdmin.showToast) {
            MetroAdmin.showToast(message, 'warning');
        }
    }
};

// Initialize on DOM ready
$(document).ready(() => {
    TouchSpin.init();
});
