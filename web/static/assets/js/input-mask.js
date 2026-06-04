/* ============================================
   Input Masks - IMask.js Implementation
   ============================================ */

'use strict';

document.addEventListener('DOMContentLoaded', function() {
    initInputMasks();
});

function initInputMasks() {
    // ==========================================
    // PHONE NUMBER MASKS
    // ==========================================
    
    // Phone mask: +62 812-3456-7890
    document.querySelectorAll('[data-mask="phone"]').forEach(input => {
        IMask(input, {
            mask: '+62 {000}-{0000}-{0000}'
        });
    });

    // US Phone: (555) 123-4567
    document.querySelectorAll('[data-mask="phone-us"]').forEach(input => {
        IMask(input, {
            mask: '(000) 000-0000'
        });
    });

    // UK Phone: +44 20 7946 0958
    document.querySelectorAll('[data-mask="phone-uk"]').forEach(input => {
        IMask(input, {
            mask: '+44 00 0000 0000'
        });
    });

    // Japan Phone: +81-90-1234-5678
    document.querySelectorAll('[data-mask="phone-jp"]').forEach(input => {
        IMask(input, {
            mask: '+81-00-0000-0000'
        });
    });

    // Landline: (021) 123-4567
    document.querySelectorAll('[data-mask="landline"]').forEach(input => {
        IMask(input, {
            mask: '(000) 000-0000'
        });
    });

    // ==========================================
    // DATE MASKS
    // ==========================================
    
    // Date DD/MM/YYYY
    document.querySelectorAll('[data-mask="date-dmy"]').forEach(input => {
        IMask(input, {
            mask: 'DD/MM/YYYY',
            blocks: {
                DD: { mask: IMask.MaskedRange, from: 1, to: 31, maxLength: 2 },
                MM: { mask: IMask.MaskedRange, from: 1, to: 12, maxLength: 2 },
                YYYY: { mask: IMask.MaskedRange, from: 1900, to: 2099, maxLength: 4 }
            }
        });
    });

    // Date MM-DD-YYYY
    document.querySelectorAll('[data-mask="date-mdy"]').forEach(input => {
        IMask(input, {
            mask: 'MM-DD-YYYY',
            blocks: {
                MM: { mask: IMask.MaskedRange, from: 1, to: 12, maxLength: 2 },
                DD: { mask: IMask.MaskedRange, from: 1, to: 31, maxLength: 2 },
                YYYY: { mask: IMask.MaskedRange, from: 1900, to: 2099, maxLength: 4 }
            }
        });
    });

    // Date YYYY.MM.DD
    document.querySelectorAll('[data-mask="date-ymd"]').forEach(input => {
        IMask(input, {
            mask: 'YYYY.MM.DD',
            blocks: {
                YYYY: { mask: IMask.MaskedRange, from: 1900, to: 2099, maxLength: 4 },
                MM: { mask: IMask.MaskedRange, from: 1, to: 12, maxLength: 2 },
                DD: { mask: IMask.MaskedRange, from: 1, to: 31, maxLength: 2 }
            }
        });
    });

    // Expiry Date MM/YY
    document.querySelectorAll('[data-mask="expiry-date"]').forEach(input => {
        IMask(input, {
            mask: 'MM/YY',
            blocks: {
                MM: { mask: IMask.MaskedRange, from: 1, to: 12, maxLength: 2 },
                YY: { mask: IMask.MaskedRange, from: 0, to: 99, maxLength: 2 }
            }
        });
    });

    // Time 24h HH:MM
    document.querySelectorAll('[data-mask="time-24h"]').forEach(input => {
        IMask(input, {
            mask: 'HH:MM',
            blocks: {
                HH: { mask: IMask.MaskedRange, from: 0, to: 23, maxLength: 2 },
                MM: { mask: IMask.MaskedRange, from: 0, to: 59, maxLength: 2 }
            }
        });
    });

    // Time 12h HH:MM AM/PM
    document.querySelectorAll('[data-mask="time-12h"]').forEach(input => {
        IMask(input, {
            mask: 'HH:MM',
            blocks: {
                HH: { mask: IMask.MaskedRange, from: 1, to: 12, maxLength: 2 },
                MM: { mask: IMask.MaskedRange, from: 0, to: 59, maxLength: 2 }
            }
        });
    });

    // DateTime DD/MM/YYYY HH:MM
    document.querySelectorAll('[data-mask="datetime"]').forEach(input => {
        IMask(input, {
            mask: 'DD/MM/YYYY HH:MM',
            blocks: {
                DD: { mask: IMask.MaskedRange, from: 1, to: 31, maxLength: 2 },
                MM: { mask: IMask.MaskedRange, from: 1, to: 12, maxLength: 2 },
                YYYY: { mask: IMask.MaskedRange, from: 1900, to: 2099, maxLength: 4 },
                HH: { mask: IMask.MaskedRange, from: 0, to: 23, maxLength: 2 },
                MM: { mask: IMask.MaskedRange, from: 0, to: 59, maxLength: 2 }
            }
        });
    });

    // ==========================================
    // CREDIT CARD & PAYMENT MASKS
    // ==========================================
    
    // Credit Card
    document.querySelectorAll('[data-mask="credit-card"]').forEach(input => {
        IMask(input, {
            mask: '0000 0000 0000 0000'
        });
    });

    // CVV
    document.querySelectorAll('[data-mask="cvv"]').forEach(input => {
        IMask(input, {
            mask: '0000',
            maxLength: 4
        });
    });

    // IDR Currency
    document.querySelectorAll('[data-mask="currency-idr"]').forEach(input => {
        IMask(input, {
            mask: Number,
            radix: ',',
            thousandsSeparator: '.',
            scale: 0
        });
    });

    // USD Currency
    document.querySelectorAll('[data-mask="currency-usd"]').forEach(input => {
        IMask(input, {
            mask: Number,
            radix: '.',
            thousandsSeparator: ',',
            scale: 2
        });
    });

    // EUR Currency
    document.querySelectorAll('[data-mask="currency-eur"]').forEach(input => {
        IMask(input, {
            mask: Number,
            radix: ',',
            thousandsSeparator: '.',
            scale: 2
        });
    });

    // Number with thousands separator
    document.querySelectorAll('[data-mask="number-thousands"]').forEach(input => {
        IMask(input, {
            mask: Number,
            radix: '.',
            thousandsSeparator: ',',
            scale: 0
        });
    });

    // ==========================================
    // NUMBER MASKS
    // ==========================================
    
    // Integer only
    document.querySelectorAll('[data-mask="integer"]').forEach(input => {
        IMask(input, {
            mask: '000000000000',
            maxLength: 12
        });
    });

    // Decimal
    document.querySelectorAll('[data-mask="decimal"]').forEach(input => {
        IMask(input, {
            mask: Number,
            radix: '.',
            scale: 2
        });
    });

    // Percentage
    document.querySelectorAll('[data-mask="percentage"]').forEach(input => {
        IMask(input, {
            mask: Number,
            min: 0,
            max: 100,
            radix: '.',
            scale: 0
        });
    });

    // Percentage Decimal
    document.querySelectorAll('[data-mask="percentage-decimal"]').forEach(input => {
        IMask(input, {
            mask: Number,
            min: 0,
            max: 100,
            radix: '.',
            scale: 2
        });
    });

    // ==========================================
    // INDONESIAN ID MASKS
    // ==========================================
    
    // NIK (KTP) - 16 digits
    document.querySelectorAll('[data-mask="nik"]').forEach(input => {
        IMask(input, {
            mask: '0000000000000000',
            maxLength: 16
        });
    });

    // NPWP
    document.querySelectorAll('[data-mask="npwp"]').forEach(input => {
        IMask(input, {
            mask: '00.000.000.0-000.000'
        });
    });

    // Passport
    document.querySelectorAll('[data-mask="passport"]').forEach(input => {
        IMask(input, {
            mask: 'AA0000000',
            definitions: {
                'A': { mask: /[a-zA-Z]/ }
            }
        });
    });

    // ==========================================
    // NETWORK & TECH MASKS
    // ==========================================
    
    // IP Address
    document.querySelectorAll('[data-mask="ip-address"]').forEach(input => {
        IMask(input, {
            mask: 'IP.IP.IP.IP',
            blocks: {
                IP: {
                    mask: Number,
                    scale: 0,
                    min: 0,
                    max: 255,
                    maxLength: 3
                }
            }
        });
    });

    // MAC Address
    document.querySelectorAll('[data-mask="mac-address"]').forEach(input => {
        IMask(input, {
            mask: 'AA:AA:AA:AA:AA:AA',
            definitions: {
                'A': {
                    mask: /[0-9a-fA-F]/
                }
            }
        });
    });

    // HEX Color
    document.querySelectorAll('[data-mask="hex-color"]').forEach(input => {
        IMask(input, {
            mask: /#[0-9a-fA-F]{0,6}/
        });
    });

    // ==========================================
    // CUSTOM PATTERNS
    // ==========================================
    
    // ISBN
    document.querySelectorAll('[data-mask="isbn"]').forEach(input => {
        IMask(input, {
            mask: '000-0-00000-000-0'
        });
    });

    // Serial Number
    document.querySelectorAll('[data-mask="serial"]').forEach(input => {
        IMask(input, {
            mask: 'AAA-000-00000',
            definitions: {
                'A': { mask: /[a-zA-Z]/ }
            }
        });
    });

    // License Plate
    document.querySelectorAll('[data-mask="license-plate"]').forEach(input => {
        IMask(input, {
            mask: 'AA 0000 AAA',
            definitions: {
                'A': { mask: /[a-zA-Z]/ }
            }
        });
    });

    // Bank Account
    document.querySelectorAll('[data-mask="bank-account"]').forEach(input => {
        IMask(input, {
            mask: '000000000000',
            maxLength: 12
        });
    });

    // Toll Free
    document.querySelectorAll('[data-mask="toll-free"]').forEach(input => {
        IMask(input, {
            mask: '000-000-0000'
        });
    });

    console.log('✅ Input masks initialized successfully');
    console.log('📊 Total masks applied:', document.querySelectorAll('[data-mask]').length);
}
