/* ============================================
   Metro Admin - Core Application Logic
   ============================================ */

'use strict';

const MetroAdmin = {
    // App State
    state: {
        theme: localStorage.getItem('metro-theme') || 'light',
        sidebarCollapsed: localStorage.getItem('sidebar-collapsed') === 'true',
        currentPage: 'dashboard',
    },

    // Initialize on DOM ready
    init() {
        this.initTheme();
        this.initSidebar();
        this.initNavbar();
        this.initCommandPalette();
        this.initNotificationPanel();
        this.initNProgress();
        this.initKeyboardShortcuts();
        this.initCounters();
        this.initCharts();
        this.initColorPickers();
        this.initPasswordToggle();
        this.initRangeSliders();
        this.initNumberInputs();
        this.initSearchInputs();
        this.initInputGroupDropdowns();
        // Touch Spin initialized in touch-spin.js
    },

    // ==========================================
    // THEME SYSTEM
    // ==========================================
    initTheme() {
        document.documentElement.setAttribute('data-theme', this.state.theme);
        this.updateThemeIcon();

        // Event delegation untuk theme toggle (support dynamic loading)
        $(document).on('click', '#themeToggle', () => {
            this.toggleTheme();
        });
    },

    toggleTheme() {
        this.state.theme = this.state.theme === 'light' ? 'dark' : 'light';
        document.documentElement.setAttribute('data-theme', this.state.theme);
        localStorage.setItem('metro-theme', this.state.theme);
        this.updateThemeIcon();
    },

    updateThemeIcon() {
        const icon = $('#themeToggle i');
        if (this.state.theme === 'dark') {
            icon.removeClass('fa-moon').addClass('fa-sun');
        } else {
            icon.removeClass('fa-sun').addClass('fa-moon');
        }
    },

    // ==========================================
    // SIDEBAR
    // ==========================================
    initSidebar() {
        const sidebar = $('#appSidebar');
        
        if (this.state.sidebarCollapsed) {
            sidebar.addClass('collapsed');
        }

        // Event delegation untuk sidebar elements (support dynamic loading)
        
        // Use MutationObserver to detect when sidebar content is added
        const observer = new MutationObserver((mutations) => {
            mutations.forEach((mutation) => {
                if (mutation.addedNodes.length && $('#appSidebar .nav-list').length) {
                    this.setActiveMenu();
                    observer.disconnect(); // Stop observing once we've set the active menu
                }
            });
        });
        
        // Start observing the sidebar container for changes
        const sidebarContainer = document.getElementById('sidebar-container');
        if (sidebarContainer) {
            observer.observe(sidebarContainer, { childList: true, subtree: true });
        }
        
        // Fallback: check on interval for dynamically loaded sidebar
        const checkSidebar = setInterval(() => {
            if ($('#appSidebar .nav-list').length) {
                this.setActiveMenu();
                clearInterval(checkSidebar);
            }
        }, 100);
        
        // Stop checking after 5 seconds to avoid infinite loop
        setTimeout(() => clearInterval(checkSidebar), 5000);
        
        // Toggle sidebar
        $(document).on('click', '#sidebarToggle', () => {
            sidebar.toggleClass('collapsed');
            this.state.sidebarCollapsed = sidebar.hasClass('collapsed');
            localStorage.setItem('sidebar-collapsed', this.state.sidebarCollapsed);
        });

        // Mobile toggle
        $(document).on('click', '#mobileToggle', () => {
            sidebar.toggleClass('mobile-open');
            $('#sidebarOverlay').toggleClass('active');
        });

        // Close sidebar on overlay click (mobile)
        $(document).on('click', '#sidebarOverlay', () => {
            sidebar.removeClass('mobile-open');
            $('#sidebarOverlay').removeClass('active');
        });

        // Close sidebar on outside click (mobile)
        $(document).on('click', (e) => {
            if ($(window).width() < 992) {
                if (!$(e.target).closest('.app-sidebar, #mobileToggle').length) {
                    sidebar.removeClass('mobile-open');
                    $('#sidebarOverlay').removeClass('active');
                }
            }
        });

        // Close sidebar when clicking nav link (mobile)
        $(document).on('click', '.sidebar-nav .nav-link', function() {
            if ($(window).width() < 992 && !$(this).parent().hasClass('has-submenu')) {
                sidebar.removeClass('mobile-open');
                $('#sidebarOverlay').removeClass('active');
            }
        });

        // Submenu toggle
        $(document).on('click', '.sidebar-nav .nav-item.has-submenu > .nav-link', function(e) {
            e.preventDefault();
            $(this).parent().toggleClass('open');
        });

        // Menu search
        $(document).on('input', '#menuSearch', function() {
            const query = $(this).val().toLowerCase();
            $('.nav-item').each(function() {
                const text = $(this).find('.nav-link span').first().text().toLowerCase();
                if (query === '' || text.includes(query)) {
                    $(this).show();
                } else {
                    $(this).hide();
                }
            });
            $('.nav-section').toggle(query === '');
        });
    },

    // Set active menu based on current URL
    setActiveMenu() {
        const currentPath = window.location.pathname;
        const $sidebar = $('#appSidebar');
        
        if (!$sidebar.length) return;
        
        console.log('[Sidebar] Setting active menu for:', currentPath);
        
        // Remove all active classes first
        $sidebar.find('.nav-item').removeClass('active');
        $sidebar.find('.has-submenu').removeClass('open');
        $sidebar.find('.submenu li').removeClass('active');
        
        // Find matching link
        const $activeLink = $sidebar.find(`a[href="${currentPath}"]`);
        
        console.log('[Sidebar] Found active link:', $activeLink.length);
        
        if ($activeLink.length) {
            // Check if it's a submenu item
            const $submenuLi = $activeLink.parent('li');
            if ($submenuLi.parent().hasClass('submenu')) {
                // It's a submenu item - activate the parent .has-submenu
                console.log('[Sidebar] Activating submenu item');
                $submenuLi.addClass('active');
                const $hasSubmenu = $activeLink.closest('.has-submenu');
                if ($hasSubmenu.length) {
                    $hasSubmenu.addClass('open');
                    $hasSubmenu.addClass('active');
                }
            } else {
                // It's a top-level menu item
                console.log('[Sidebar] Activating top-level menu item');
                const $navItem = $activeLink.closest('.nav-item');
                $navItem.addClass('active');
            }
        } else {
            // Default to dashboard if no match
            console.log('[Sidebar] No match, defaulting to dashboard');
            $sidebar.find('a[href="/dashboard.html"]').closest('.nav-item').addClass('active');
        }
    },

    // ==========================================
    // NAVBAR
    // ==========================================
    initNavbar() {
        // Event delegation untuk navbar elements (support dynamic loading)
        
        // Fullscreen toggle
        $(document).on('click', '#fullscreenToggle', () => {
            if (!document.fullscreenElement) {
                document.documentElement.requestFullscreen();
                $('#fullscreenToggle i').removeClass('fa-expand').addClass('fa-compress');
            } else {
                document.exitFullscreen();
                $('#fullscreenToggle i').removeClass('fa-compress').addClass('fa-expand');
            }
        });

        // Notification dropdown toggle
        $(document).on('click', '#notifBtn', (e) => {
            e.stopPropagation();
            const dropdown = $('#notifDropdown .navbar-dropdown');
            const isShowing = dropdown.hasClass('show');
            
            // Close all dropdowns first
            $('.navbar-dropdown').removeClass('show');
            
            if (!isShowing) {
                dropdown.addClass('show');
            }
        });

        // Message dropdown toggle
        $(document).on('click', '#msgBtn', (e) => {
            e.stopPropagation();
            const dropdown = $('#msgDropdown .navbar-dropdown');
            const isShowing = dropdown.hasClass('show');
            
            // Close all dropdowns first
            $('.navbar-dropdown').removeClass('show');
            
            if (!isShowing) {
                dropdown.addClass('show');
            }
        });

        // User dropdown toggle
        $(document).on('click', '#userDropdown', (e) => {
            e.stopPropagation();
            const dropdown = $('#userDropdownWrapper .navbar-dropdown');
            const isShowing = dropdown.hasClass('show');
            
            // Close all dropdowns first
            $('.navbar-dropdown').removeClass('show');
            
            if (!isShowing) {
                dropdown.addClass('show');
            }
        });

        // Close dropdowns when clicking outside
        $(document).on('click', (e) => {
            if (!$(e.target).closest('.navbar-dropdown-wrapper').length) {
                $('.navbar-dropdown').removeClass('show');
            }
        });

        // Close dropdowns on escape
        $(document).on('keydown', (e) => {
            if (e.key === 'Escape') {
                $('.navbar-dropdown').removeClass('show');
            }
        });
    },

    // ==========================================
    // COMMAND PALETTE (CTRL+K)
    // ==========================================
    initCommandPalette() {
        const overlay = $('#commandPaletteOverlay');
        const input = $('#commandInput');
        const results = $('#commandResults');

        // Event delegation untuk command palette
        
        // Open command palette
        $(document).on('click', '#searchTrigger', () => this.openCommandPalette());

        // Close on overlay click
        overlay.on('click', (e) => {
            if ($(e.target).is(overlay)) {
                this.closeCommandPalette();
            }
        });

        // Filter commands
        input.on('input', function() {
            const query = $(this).val().toLowerCase();
            $('.command-item').each(function() {
                const text = $(this).find('span').text().toLowerCase();
                $(this).toggle(text.includes(query));
            });
            // Hide empty groups
            $('.command-group').each(function() {
                const visible = $(this).find('.command-item:visible').length;
                $(this).toggle(visible > 0);
            });
        });

        // Command item click (event delegation)
        $(document).on('click', '.command-item', function() {
            const action = $(this).data('action');
            const target = $(this).data('target');

            if (action === 'navigate' && target) {
                window.location.href = target;
            } else if (action === 'theme') {
                MetroAdmin.toggleTheme();
            } else if (action === 'fullscreen') {
                $('#fullscreenToggle').click();
            }
            MetroAdmin.closeCommandPalette();
        });
    },

    openCommandPalette() {
        $('#commandPaletteOverlay').addClass('active');
        setTimeout(() => $('#commandInput').focus(), 100);
    },

    closeCommandPalette() {
        $('#commandPaletteOverlay').removeClass('active');
        $('#commandInput').val('');
        $('.command-item').show();
        $('.command-group').show();
    },

    // ==========================================
    // NOTIFICATION PANEL
    // ==========================================
    initNotificationPanel() {
        const panel = $('#notifPanel');

        // Event delegation untuk notification panel
        $(document).on('click', '#notifBtn', (e) => {
            e.stopPropagation();
            panel.toggleClass('open');
        });

        $(document).on('click', (e) => {
            if (!$(e.target).closest('#notifPanel, #notifBtn').length) {
                panel.removeClass('open');
            }
        });

        // Mark all read (event delegation)
        $(document).on('click', '.mark-all-read', () => {
            $('.notif-item').removeClass('unread');
            $('.notif-badge').text('0').hide();
        });
    },

    // ==========================================
    // NPROGRESS
    // ==========================================
    initNProgress() {
        NProgress.configure({
            showSpinner: false,
            minimum: 0.15,
            speed: 400,
            trickleSpeed: 200
        });
    },

    // ==========================================
    // KEYBOARD SHORTCUTS
    // ==========================================
    initKeyboardShortcuts() {
        $(document).on('keydown', (e) => {
            // Ctrl+K - Command Palette
            if ((e.ctrlKey || e.metaKey) && e.key === 'k') {
                e.preventDefault();
                this.openCommandPalette();
            }
            // Escape
            if (e.key === 'Escape') {
                this.closeCommandPalette();
                $('#notifPanel').removeClass('open');
            }
        });
    },

    // ==========================================
    // KEYBOARD SHORTCUTS
    // ==========================================
    showToast(message, type = 'info') {
        const icons = {
            success: 'fa-check-circle',
            error: 'fa-times-circle',
            warning: 'fa-exclamation-triangle',
            info: 'fa-info-circle'
        };
        
        const toast = $(`
            <div class="toast-notification">
                <i class="fa-solid ${icons[type]}" style="color: var(--${type === 'error' ? 'danger' : type})"></i>
                <span style="font-size: 13px">${message}</span>
            </div>
        `).appendTo('body');

        setTimeout(() => {
            toast.addClass('exit');
            setTimeout(() => toast.remove(), 300);
        }, 3500);
    },

    animateCounter(element, target, duration = 1500) {
        const start = 0;
        const startTime = performance.now();
        const suffix = element.dataset.suffix || '';

        const update = (currentTime) => {
            const elapsed = currentTime - startTime;
            const progress = Math.min(elapsed / duration, 1);
            const eased = 1 - Math.pow(1 - progress, 3);
            const current = Math.floor(start + (target - start) * eased);
            
            element.textContent = current.toLocaleString('id-ID') + suffix;
            
            if (progress < 1) {
                requestAnimationFrame(update);
            }
        };

        requestAnimationFrame(update);
    },

    // Init counters
    initCounters() {
        document.querySelectorAll('.counter').forEach(el => {
            const target = parseInt(el.dataset.target);
            if (target) this.animateCounter(el, target);
        });
        
        // Stagger animation
        setTimeout(() => {
            document.querySelectorAll('.stagger-children').forEach(el => el.classList.add('visible'));
        }, 100);
    },

    // Init charts (if on dashboard)
    initCharts() {
        if (typeof Chart === 'undefined') return;
        
        // Revenue Chart
        const revenueCtx = document.getElementById('revenueChart');
        if (revenueCtx) {
            new Chart(revenueCtx, {
                type: 'line',
                data: {
                    labels: ['Jan', 'Feb', 'Mar', 'Apr', 'Mei', 'Jun', 'Jul', 'Agt', 'Sep', 'Okt', 'Nov', 'Des'],
                    datasets: [{
                        label: 'Pendapatan (Juta)',
                        data: [65, 72, 68, 85, 92, 78, 95, 88, 102, 110, 98, 125],
                        borderColor: '#0078D4',
                        backgroundColor: 'rgba(0, 120, 212, 0.06)',
                        borderWidth: 2.5,
                        fill: true,
                        tension: 0.4,
                        pointRadius: 4,
                        pointHoverRadius: 6
                    }]
                },
                options: {
                    responsive: true,
                    maintainAspectRatio: false,
                    plugins: { legend: { display: false } },
                    scales: {
                        x: { grid: { display: false }, ticks: { font: { size: 11 } } },
                        y: { grid: { color: 'rgba(0,0,0,0.04)' }, ticks: { font: { size: 11 } } }
                    },
                    interaction: { intersect: false, mode: 'index' }
                }
            });
        }

        // User Growth
        const userCtx = document.getElementById('userGrowthChart');
        if (userCtx) {
            new Chart(userCtx, {
                type: 'bar',
                data: {
                    labels: ['Jan', 'Feb', 'Mar', 'Apr', 'Mei', 'Jun'],
                    datasets: [{
                        label: 'Pengguna Baru',
                        data: [320, 450, 380, 520, 610, 740],
                        backgroundColor: 'rgba(0, 120, 212, 0.7)',
                        borderRadius: 6
                    }]
                },
                options: {
                    responsive: true,
                    maintainAspectRatio: false,
                    plugins: { legend: { display: false } },
                    scales: {
                        x: { grid: { display: false }, ticks: { font: { size: 11 } } },
                        y: { grid: { color: 'rgba(0,0,0,0.04)' }, ticks: { font: { size: 11 } } }
                    }
                }
            });
        }

        // Sales Chart
        const salesCtx = document.getElementById('salesChart');
        if (salesCtx) {
            new Chart(salesCtx, {
                type: 'doughnut',
                data: {
                    labels: ['Elektronik', 'Fashion', 'Makanan', 'Kesehatan', 'Lainnya'],
                    datasets: [{
                        data: [35, 25, 20, 12, 8],
                        backgroundColor: ['#0078D4', '#16C60C', '#FFB900', '#E81123', '#00BCF2'],
                        borderWidth: 0,
                        spacing: 2
                    }]
                },
                options: {
                    responsive: true,
                    maintainAspectRatio: false,
                    cutout: '65%',
                    plugins: { legend: { position: 'bottom', labels: { padding: 16, font: { size: 11 } } } }
                }
            });
        }

        // Top Products
        const topCtx = document.getElementById('topProductsChart');
        if (topCtx) {
            new Chart(topCtx, {
                type: 'bar',
                data: {
                    labels: ['Laptop Pro', 'Headphone X', 'Mouse Wireless', 'Keyboard Mech', 'Monitor 27"'],
                    datasets: [{
                        label: 'Terjual',
                        data: [142, 98, 87, 76, 65],
                        backgroundColor: ['#0078D4', '#00BCF2', '#16C60C', '#FFB900', '#E81123'],
                        borderRadius: 4
                    }]
                },
                options: {
                    responsive: true,
                    maintainAspectRatio: false,
                    indexAxis: 'y',
                    plugins: { legend: { display: false } },
                    scales: {
                        x: { grid: { color: 'rgba(0,0,0,0.04)' } },
                        y: { grid: { display: false } }
                    }
                }
            });
        }
    },

    // ==========================================
    // COLOR PICKER SYNC
    // ==========================================
    initColorPickers() {
        // Find all color picker groups
        $('input[type="color"]').each(function() {
            const colorInput = $(this);
            const parent = colorInput.closest('.form-group');
            const textInput = parent.find('input[type="text"][readonly]');
            
            if (textInput.length) {
                // Sync color picker to text input
                colorInput.on('input', function() {
                    const colorValue = $(this).val();
                    textInput.val(colorValue);
                });
                
                // Sync text input to color picker (if user edits manually)
                textInput.on('input', function() {
                    const textValue = $(this).val();
                    // Validate hex color format
                    if (/^#[0-9A-Fa-f]{6}$/.test(textValue)) {
                        colorInput.val(textValue);
                    }
                });
            }
        });
    },

    // ==========================================
    // PASSWORD VISIBILITY TOGGLE
    // ==========================================
    initPasswordToggle() {
        $(document).on('click', '.input-with-icon-right i.fa-eye, .input-with-icon-right i.fa-eye-slash', function() {
            const icon = $(this);
            const input = icon.siblings('input');
            
            if (input.attr('type') === 'password') {
                input.attr('type', 'text');
                icon.removeClass('fa-eye').addClass('fa-eye-slash');
                icon.attr('title', 'Hide password');
            } else {
                input.attr('type', 'password');
                icon.removeClass('fa-eye-slash').addClass('fa-eye');
                icon.attr('title', 'Toggle password visibility');
            }
        });
    },

    // ==========================================
    // RANGE SLIDER VALUE DISPLAY
    // ==========================================
    initRangeSliders() {
        $('input[type="range"]').each(function() {
            const rangeInput = $(this);
            const helperText = rangeInput.closest('.form-group').find('.helper-text');
            
            // Function to update helper text based on range label
            const updateDisplay = function() {
                const value = rangeInput.val();
                const min = rangeInput.attr('min') || 0;
                const max = rangeInput.attr('max') || 100;
                const label = rangeInput.closest('.form-group').find('label').first().text();
                
                // Update helper text with current value
                if (label.includes('Volume')) {
                    helperText.html('<i class="fa-solid fa-volume-high"></i> Current: ' + value + '%');
                } else if (label.includes('Price')) {
                    helperText.html('<i class="fa-solid fa-dollar-sign"></i> Current: $' + value);
                } else if (label.includes('Brightness')) {
                    helperText.html('<i class="fa-solid fa-sun"></i> Current: ' + value);
                } else {
                    helperText.html('<i class="fa-solid fa-sliders"></i> Current: ' + value + ' (' + min + '-' + max + ')');
                }
            };
            
            // Update on input change
            rangeInput.on('input', updateDisplay);
            
            // Initial display
            updateDisplay();
        });
    },

    // ==========================================
    // NUMBER INPUT SPINNER
    // ==========================================
    initNumberInputs() {
        $('input[type="number"]').each(function() {
            const numberInput = $(this);
            const helperText = numberInput.closest('.form-group').find('.helper-text');
            
            numberInput.on('input change', function() {
                const value = $(this).val();
                const min = $(this).attr('min');
                const max = $(this).attr('max');
                const step = $(this).attr('step');
                
                // Update helper text if exists
                if (helperText.length) {
                    const label = $(this).closest('.form-group').find('label').first().text();
                    
                    if (label.includes('Decimal')) {
                        helperText.html('<i class="fa-solid fa-coins"></i> Value: ' + value + ' (Step: ' + step + ')');
                    } else if (label.includes('Range')) {
                        helperText.html('<i class="fa-solid fa-arrows-left-right"></i> Min: ' + min + ', Max: ' + max + ', Current: ' + value);
                    } else {
                        helperText.html('<i class="fa-solid fa-circle-info"></i> Current value: ' + value);
                    }
                }
            });
        });
    },

    // ==========================================
    // SEARCH INPUT CLEAR BUTTON
    // ==========================================
    initSearchInputs() {
        $('input[type="search"]').each(function() {
            const searchInput = $(this);
            const parent = searchInput.closest('.input-with-icon');
            
            // Add clear button if not exists
            if (!parent.find('.search-clear').length) {
                searchInput.after('<i class="fa-solid fa-xmark search-clear" style="position: absolute; right: 12px; cursor: pointer; color: var(--text-tertiary); display: none;"></i>');
            }
            
            const clearBtn = parent.find('.search-clear');
            
            // Show/hide clear button
            searchInput.on('input', function() {
                if ($(this).val().length > 0) {
                    clearBtn.show();
                } else {
                    clearBtn.hide();
                }
            });
            
            // Clear on click
            clearBtn.on('click', function() {
                searchInput.val('').focus();
                clearBtn.hide();
            });
        });
    },

    // ==========================================
    // INPUT GROUP DROPDOWNS
    // ==========================================
    initInputGroupDropdowns() {
        // Dropdown toggle button click
        $(document).on('click', '.input-group .dropdown-toggle', function(e) {
            e.preventDefault();
            e.stopPropagation();
            
            const btn = $(this);
            const inputGroup = btn.closest('.input-group');
            
            // Check if dropdown menu already exists
            let dropdownMenu = inputGroup.find('.dropdown-menu');
            
            if (!dropdownMenu.length) {
                // Create dropdown menu
                dropdownMenu = $('<div class="dropdown-menu">').html(`
                    <a class="dropdown-item" href="#">Action 1</a>
                    <a class="dropdown-item" href="#">Action 2</a>
                    <a class="dropdown-item" href="#">Action 3</a>
                    <div class="dropdown-divider"></div>
                    <a class="dropdown-item" href="#">Something else</a>
                `);
                
                inputGroup.append(dropdownMenu);
            }
            
            // Toggle dropdown
            dropdownMenu.toggleClass('show');
            
            // Position dropdown below input group
            const inputGroupOffset = inputGroup.offset();
            const inputGroupHeight = inputGroup.outerHeight();
            
            dropdownMenu.css({
                position: 'absolute',
                top: inputGroupHeight + 'px',
                left: 0,
                zIndex: 1000
            });
        });
        
        // Close dropdown when clicking outside
        $(document).on('click', function(e) {
            if (!$(e.target).closest('.input-group').length) {
                $('.input-group .dropdown-menu').removeClass('show');
            }
        });
        
        // Handle dropdown item click
        $(document).on('click', '.input-group .dropdown-menu .dropdown-item', function(e) {
            e.preventDefault();
            const text = $(this).text();
            const inputGroup = $(this).closest('.input-group');
            const input = inputGroup.find('input.form-control');
            
            // Update input value
            if (input.length) {
                input.val(text);
            }
            
            // Close dropdown
            inputGroup.find('.dropdown-menu').removeClass('show');
        });
    }
};

// Initialize on DOM ready
$(document).ready(() => {
    // Auto-refresh token on every page load
    refreshTokenOnLoad();
    
    MetroAdmin.init();
});

/**
 * Refresh access token using refresh_token on page load
 * This prevents users from being logged out due to expired access_token
 */
function refreshTokenOnLoad() {
    const refreshToken = localStorage.getItem('refresh_token');
    const accessToken = localStorage.getItem('access_token');
    
    // Skip if on login page or no tokens
    if (window.location.pathname.includes('login.html')) {
        return;
    }
    
    // If no refresh token, redirect to login
    if (!refreshToken) {
        if (!window.location.pathname.includes('login.html')) {
            window.location.href = '/login.html';
        }
        return;
    }
    
    // Synchronously refresh token before page continues
    $.ajax({
        url: '/api/admin/auth/refresh',
        method: 'POST',
        async: false, // Synchronous to block page until token is refreshed
        contentType: 'application/json',
        data: JSON.stringify({ refresh_token: refreshToken }),
        success: function(response) {
            if (response.success && response.data) {
                localStorage.setItem('access_token', response.data.access_token);
                if (response.data.refresh_token) {
                    localStorage.setItem('refresh_token', response.data.refresh_token);
                }
                console.log('[Auth] Token refreshed successfully');
            }
        },
        error: function(xhr) {
            console.warn('[Auth] Token refresh failed, redirecting to login');
            localStorage.removeItem('access_token');
            localStorage.removeItem('refresh_token');
            window.location.href = '/login.html';
        }
    });
}
