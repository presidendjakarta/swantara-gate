@extends('layouts.app')

@section('title', 'Datetimepicker')

@push('styles')
@endpush

@section('content')
<div class="page-header">
    <div>
        <h1>Datetimepicker</h1>
        <p>Date and time picker components menggunakan Metro UI built-in</p>
    </div>
    <div style="display: flex; gap: 10px;">
        <a href="https://v5.metroui.org.ua/components/calendar-picker/" target="_blank" class="btn btn-secondary">
            <i class="fa-solid fa-book"></i>
            <span>Documentation</span>
        </a>
    </div>
</div>

<!-- Info Alert -->
<div class="content-card" style="margin-bottom: 24px; border-left: 4px solid var(--accent);">
    <div class="card-body" style="padding: 16px 20px;">
        <div style="display: flex; gap: 12px; align-items: start;">
            <i class="fa-solid fa-circle-info" style="color: var(--accent); font-size: 20px; margin-top: 2px;"></i>
            <div style="flex: 1;">
                <h4 style="margin-bottom: 4px; font-size: 14px;">Metro UI Date & Time Components</h4>
                <p style="font-size: 13px; color: var(--text-secondary); margin: 0;">
                    Metro UI menyediakan 4 komponen: 
                    <strong>Calendar Picker</strong> (input + calendar popup), 
                    <strong>Date Picker</strong> (wheel-based), 
                    <strong>Time Picker</strong> (wheel-based), dan 
                    <strong>Calendar</strong> (inline). Semua sudah ter-include di <code>metro.js</code>.
                </p>
            </div>
        </div>
    </div>
</div>

<!-- Calendar Picker Section -->
<div class="dtp-section-title">
    <i class="fa-solid fa-calendar-days"></i>
    Calendar Picker <span class="badge badge-primary">Recommended</span>
</div>

<div class="datetimepicker-grid">
    <div class="content-card">
        <div class="card-header">
            <div class="card-header-left">
                <div class="card-icon bg-primary">
                    <i class="fa-solid fa-calendar"></i>
                </div>
                <div>
                    <h3>Basic Calendar Picker</h3>
                    <p class="card-subtitle">Input field with calendar popup</p>
                </div>
            </div>
        </div>
        <div class="card-body">
            <div class="dtp-example">
                <label class="dtp-label">
                    Default
                    <span class="dtp-pattern">data-role="calendar-picker"</span>
                </label>
                <input data-role="calendar-picker">
                <div class="dtp-helper">
                    <i class="fa-solid fa-circle-info"></i>
                    Basic calendar picker dengan format default
                </div>
            </div>

            <div class="dtp-example">
                <label class="dtp-label">
                    With Value
                    <span class="dtp-pattern">value="2026/06/15"</span>
                </label>
                <input data-role="calendar-picker" value="2026/06/15">
                <div class="dtp-helper">
                    <i class="fa-solid fa-circle-info"></i>
                    Dengan initial value
                </div>
            </div>

            <div class="dtp-example">
                <label class="dtp-label">
                    Custom Format
                    <span class="dtp-pattern">DD/MM/YYYY</span>
                </label>
                <input data-role="calendar-picker" 
                       data-format="DD/MM/YYYY" 
                       value="2026/06/15">
                <div class="dtp-helper">
                    <i class="fa-solid fa-circle-info"></i>
                    Format: <code>DD/MM/YYYY</code>
                </div>
            </div>

            <div class="dtp-example">
                <label class="dtp-label">
                    With Clear Button
                    <span class="dtp-pattern">clear-button</span>
                </label>
                <input data-role="calendar-picker" 
                       data-clear-button="true"
                       data-format="DD, MMM YYYY"
                       value="2026/07/01">
                <div class="dtp-helper">
                    <i class="fa-solid fa-circle-info"></i>
                    Dengan tombol clear & format: <code>DD, MMM YYYY</code>
                </div>
            </div>
        </div>
    </div>

    <div class="content-card">
        <div class="card-header">
            <div class="card-header-left">
                <div class="card-icon bg-success">
                    <i class="fa-solid fa-clock"></i>
                </div>
                <div>
                    <h3>Calendar Picker + Time</h3>
                    <p class="card-subtitle">Date and time selection</p>
                </div>
            </div>
        </div>
        <div class="card-body">
            <div class="dtp-example">
                <label class="dtp-label">
                    DateTime
                    <span class="dtp-pattern">show-time="true"</span>
                </label>
                <input data-role="calendar-picker" 
                       data-show-time="true"
                       data-format="DD/MM/YYYY HH:mm">
                <div class="dtp-helper">
                    <i class="fa-solid fa-circle-info"></i>
                    Calendar picker dengan time selection
                </div>
            </div>

            <div class="dtp-example">
                <label class="dtp-label">
                    With Initial Time
                    <span class="dtp-pattern">initial-time="14:30"</span>
                </label>
                <input data-role="calendar-picker" 
                       data-show-time="true"
                       data-initial-time="14:30"
                       data-format="DD/MM/YYYY HH:mm"
                       value="2026/06/15">
                <div class="dtp-helper">
                    <i class="fa-solid fa-circle-info"></i>
                    Dengan waktu awal 14:30
                </div>
            </div>

            <div class="dtp-example">
                <label class="dtp-label">
                    Min/Max Date
                    <span class="dtp-pattern">Restricted range</span>
                </label>
                <input data-role="calendar-picker" 
                       data-min-date="2026/01/01"
                       data-max-date="2026/12/31"
                       data-format="DD/MM/YYYY"
                       data-clear-button="true">
                <div class="dtp-helper">
                    <i class="fa-solid fa-circle-info"></i>
                    Hanya bisa pilih tanggal tahun 2026
                </div>
            </div>

            <div class="dtp-example">
                <label class="dtp-label">
                    Week Start Monday
                    <span class="dtp-pattern">week-start="1"</span>
                </label>
                <input data-role="calendar-picker" 
                       data-week-start="1"
                       data-format="DD/MM/YYYY"
                       data-show-week-number="true">
                <div class="dtp-helper">
                    <i class="fa-solid fa-circle-info"></i>
                    Minggu dimulai Senin + nomor minggu
                </div>
            </div>
        </div>
    </div>
</div>

<!-- Date Picker (Wheel) Section -->
<div class="dtp-section-title">
    <i class="fa-solid fa-arrows-up-down"></i>
    Date Picker <span class="badge badge-info">Wheel-based</span>
</div>

<div class="datetimepicker-grid">
    <div class="content-card">
        <div class="card-header">
            <div class="card-header-left">
                <div class="card-icon bg-info">
                    <i class="fa-solid fa-arrows-up-down"></i>
                </div>
                <div>
                    <h3>Wheel Date Picker</h3>
                    <p class="card-subtitle">Scroll-based date selection</p>
                </div>
            </div>
        </div>
        <div class="card-body">
            <div class="dtp-example">
                <label class="dtp-label">
                    Default
                    <span class="dtp-pattern">data-role="date-picker"</span>
                </label>
                <input data-role="date-picker">
                <div class="dtp-helper">
                    <i class="fa-solid fa-circle-info"></i>
                    Wheel-based date picker default
                </div>
            </div>

            <div class="dtp-example">
                <label class="dtp-label">
                    Custom Format
                    <span class="dtp-pattern">DD/MM/YYYY</span>
                </label>
                <input data-role="date-picker" 
                       data-format="DD/MM/YYYY"
                       data-value="2026/06/15">
                <div class="dtp-helper">
                    <i class="fa-solid fa-circle-info"></i>
                    Format: <code>DD/MM/YYYY</code>
                </div>
            </div>

            <div class="dtp-example">
                <label class="dtp-label">
                    Month & Year Only
                    <span class="dtp-pattern">day="false"</span>
                </label>
                <input data-role="date-picker" 
                       data-day="false"
                       data-format="MMM YYYY">
                <div class="dtp-helper">
                    <i class="fa-solid fa-circle-info"></i>
                    Tanpa pemilihan hari
                </div>
            </div>

            <div class="dtp-example">
                <label class="dtp-label">
                    Year Range
                    <span class="dtp-pattern">min/max year</span>
                </label>
                <input data-role="date-picker" 
                       data-min-year="2020"
                       data-max-year="2030"
                       data-format="YYYY-MM-DD">
                <div class="dtp-helper">
                    <i class="fa-solid fa-circle-info"></i>
                    Tahun dibatasi 2020-2030
                </div>
            </div>
        </div>
    </div>

    <div class="content-card">
        <div class="card-header">
            <div class="card-header-left">
                <div class="card-icon bg-warning">
                    <i class="fa-solid fa-clock"></i>
                </div>
                <div>
                    <h3>Wheel Time Picker</h3>
                    <p class="card-subtitle">Scroll-based time selection</p>
                </div>
            </div>
        </div>
        <div class="card-body">
            <div class="dtp-example">
                <label class="dtp-label">
                    Default (H:M:S)
                    <span class="dtp-pattern">data-role="time-picker"</span>
                </label>
                <input data-role="time-picker">
                <div class="dtp-helper">
                    <i class="fa-solid fa-circle-info"></i>
                    Wheel-based time picker default (jam, menit, detik)
                </div>
            </div>

            <div class="dtp-example">
                <label class="dtp-label">
                    Hours & Minutes Only
                    <span class="dtp-pattern">seconds="false"</span>
                </label>
                <input data-role="time-picker" 
                       data-seconds="false"
                       data-value="14:30:00">
                <div class="dtp-helper">
                    <i class="fa-solid fa-circle-info"></i>
                    Tanpa detik, initial value 14:30
                </div>
            </div>

            <div class="dtp-example">
                <label class="dtp-label">
                    Custom Steps
                    <span class="dtp-pattern">15-minute steps</span>
                </label>
                <input data-role="time-picker" 
                       data-seconds="false"
                       data-minutes-step="15">
                <div class="dtp-helper">
                    <i class="fa-solid fa-circle-info"></i>
                    Menit naik per 15 (00, 15, 30, 45)
                </div>
            </div>

            <div class="dtp-example">
                <label class="dtp-label">
                    With Label
                    <span class="dtp-pattern">data-label</span>
                </label>
                <input data-role="time-picker" 
                       data-label="Select time"
                       data-seconds="false"
                       data-hours-step="1"
                       data-minutes-step="5">
                <div class="dtp-helper">
                    <i class="fa-solid fa-circle-info"></i>
                    Menit step 5, dengan label
                </div>
            </div>
        </div>
    </div>
</div>

<!-- Inline Calendar Section -->
<div class="dtp-section-title">
    <i class="fa-solid fa-calendar"></i>
    Inline Calendar
</div>

<div class="datetimepicker-grid">
    <div class="content-card">
        <div class="card-header">
            <div class="card-header-left">
                <div class="card-icon bg-primary">
                    <i class="fa-solid fa-calendar"></i>
                </div>
                <div>
                    <h3>Basic Calendar</h3>
                    <p class="card-subtitle">Embedded calendar display</p>
                </div>
            </div>
        </div>
        <div class="card-body">
            <div class="dtp-example">
                <label class="dtp-label">
                    Default Calendar
                    <span class="dtp-pattern">data-role="calendar"</span>
                </label>
                <div data-role="calendar"></div>
            </div>
        </div>
    </div>

    <div class="content-card">
        <div class="card-header">
            <div class="card-header-left">
                <div class="card-icon bg-success">
                    <i class="fa-solid fa-calendar-check"></i>
                </div>
                <div>
                    <h3>Calendar Options</h3>
                    <p class="card-subtitle">Customized inline calendars</p>
                </div>
            </div>
        </div>
        <div class="card-body">
            <div class="dtp-example">
                <label class="dtp-label">
                    With Time
                    <span class="dtp-pattern">show-time="true"</span>
                </label>
                <div data-role="calendar" 
                     data-show-time="true"
                     data-initial-time="09:30"
                     data-buttons="today, clear, done"></div>
            </div>
        </div>
    </div>
</div>

<!-- Multi-Select & Wide Calendar -->
<div class="dtp-section-title">
    <i class="fa-solid fa-calendar-week"></i>
    Advanced Calendar
</div>

<div class="datetimepicker-grid">
    <div class="content-card">
        <div class="card-header">
            <div class="card-header-left">
                <div class="card-icon bg-info">
                    <i class="fa-solid fa-list-check"></i>
                </div>
                <div>
                    <h3>Multi Select</h3>
                    <p class="card-subtitle">Select multiple dates</p>
                </div>
            </div>
        </div>
        <div class="card-body">
            <div class="dtp-example">
                <label class="dtp-label">
                    Multi-Select Calendar
                    <span class="dtp-pattern">multi-select="true"</span>
                </label>
                <div data-role="calendar" 
                     data-multi-select="true"
                     data-buttons="today, clear, done"
                     data-week-start="1"></div>
                <div class="dtp-helper">
                    <i class="fa-solid fa-circle-info"></i>
                    Klik beberapa tanggal untuk multi-select
                </div>
            </div>
        </div>
    </div>

    <div class="content-card">
        <div class="card-header">
            <div class="card-header-left">
                <div class="card-icon bg-warning">
                    <i class="fa-solid fa-arrows-left-right"></i>
                </div>
                <div>
                    <h3>Wide Calendar</h3>
                    <p class="card-subtitle">Wide layout with week numbers</p>
                </div>
            </div>
        </div>
        <div class="card-body">
            <div class="dtp-example">
                <label class="dtp-label">
                    Wide Mode
                    <span class="dtp-pattern">wide="true"</span>
                </label>
                <div data-role="calendar" 
                     data-wide="true"
                     data-show-week-number="true"
                     data-week-start="1"
                     data-buttons="today, clear"></div>
                <div class="dtp-helper">
                    <i class="fa-solid fa-circle-info"></i>
                    Layout lebar dengan nomor minggu
                </div>
            </div>
        </div>
    </div>
</div>

<!-- Booking Form Example -->
<div class="dtp-section-title">
    <i class="fa-solid fa-clipboard-list"></i>
    Real-World Form Example
</div>

<div class="datetimepicker-grid full-width">
    <div class="content-card">
        <div class="card-header">
            <div class="card-header-left">
                <div class="card-icon bg-primary">
                    <i class="fa-solid fa-hotel"></i>
                </div>
                <div>
                    <h3>Hotel Booking Form</h3>
                    <p class="card-subtitle">Contoh penggunaan di form nyata</p>
                </div>
            </div>
        </div>
        <div class="card-body">
            <form style="max-width: 800px;">
                <div style="display: grid; grid-template-columns: 1fr 1fr; gap: 20px; margin-bottom: 20px;">
                    <div class="dtp-example">
                        <label class="dtp-label">Check-in Date <span style="color: var(--danger);">*</span></label>
                        <input data-role="calendar-picker" 
                               data-format="DD/MM/YYYY"
                               data-min-date="2026/01/01"
                               data-clear-button="true"
                               id="checkin-date">
                    </div>

                    <div class="dtp-example">
                        <label class="dtp-label">Check-out Date <span style="color: var(--danger);">*</span></label>
                        <input data-role="calendar-picker" 
                               data-format="DD/MM/YYYY"
                               data-min-date="2026/01/01"
                               data-clear-button="true"
                               id="checkout-date">
                    </div>
                </div>

                <div style="display: grid; grid-template-columns: 1fr 1fr; gap: 20px; margin-bottom: 20px;">
                    <div class="dtp-example">
                        <label class="dtp-label">Check-in Time</label>
                        <input data-role="time-picker" 
                               data-seconds="false"
                               data-value="14:00:00">
                    </div>

                    <div class="dtp-example">
                        <label class="dtp-label">Check-out Time</label>
                        <input data-role="time-picker" 
                               data-seconds="false"
                               data-value="12:00:00">
                    </div>
                </div>

                <div class="dtp-example" style="margin-bottom: 20px;">
                    <label class="dtp-label">Event Date & Time (Optional)</label>
                    <input data-role="calendar-picker" 
                           data-show-time="true"
                           data-format="DD/MM/YYYY HH:mm"
                           data-clear-button="true">
                    <div class="dtp-helper">
                        <i class="fa-solid fa-circle-info"></i>
                        Untuk layanan atau event terjadwal
                    </div>
                </div>

                <div class="divider"></div>

                <div style="display: flex; gap: 12px;">
                    <button type="submit" class="btn btn-primary" style="flex: 1;">
                        <i class="fa-solid fa-calendar-check"></i>
                        Book Now
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
<div class="dtp-section-title">
    <i class="fa-solid fa-code"></i>
    Implementation Guide
</div>

<div class="datetimepicker-grid full-width">
    <div class="content-card">
        <div class="card-header">
            <div class="card-header-left">
                <div class="card-icon bg-success">
                    <i class="fa-solid fa-code"></i>
                </div>
                <div>
                    <h3>Usage Examples</h3>
                    <p class="card-subtitle">How to implement Metro UI date/time pickers</p>
                </div>
            </div>
        </div>
        <div class="card-body">
            <div class="code-block">
                <div style="color: var(--text-tertiary); margin-bottom: 8px;">1. Calendar Picker (Input + Popup):</div>
                <code style="color: var(--accent);">
                    &lt;input data-role="calendar-picker"<br>
                    &nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;data-format="DD/MM/YYYY"<br>
                    &nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;data-clear-button="true"<br>
                    &nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;data-show-time="true"&gt;
                </code>
            </div>

            <div class="code-block">
                <div style="color: var(--text-tertiary); margin-bottom: 8px; margin-top: 16px;">2. Date Picker (Wheel-based):</div>
                <code style="color: var(--success);">
                    &lt;input data-role="date-picker"<br>
                    &nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;data-format="YYYY-MM-DD"<br>
                    &nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;data-min-year="2020"<br>
                    &nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;data-max-year="2030"&gt;
                </code>
            </div>

            <div class="code-block">
                <div style="color: var(--text-tertiary); margin-bottom: 8px; margin-top: 16px;">3. Time Picker (Wheel-based):</div>
                <code style="color: var(--warning);">
                    &lt;input data-role="time-picker"<br>
                    &nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;data-seconds="false"<br>
                    &nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;data-minutes-step="15"<br>
                    &nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;data-value="14:30:00"&gt;
                </code>
            </div>

            <div class="code-block">
                <div style="color: var(--text-tertiary); margin-bottom: 8px; margin-top: 16px;">4. Inline Calendar:</div>
                <code style="color: var(--info);">
                    &lt;div data-role="calendar"<br>
                    &nbsp;&nbsp;&nbsp;&nbsp;&nbsp;data-multi-select="true"<br>
                    &nbsp;&nbsp;&nbsp;&nbsp;&nbsp;data-show-time="true"<br>
                    &nbsp;&nbsp;&nbsp;&nbsp;&nbsp;data-buttons="today, clear, done"&gt;&lt;/div&gt;
                </code>
            </div>

            <div class="code-block">
                <div style="color: var(--text-tertiary); margin-bottom: 8px; margin-top: 16px;">5. JavaScript API:</div>
                <code style="color: var(--text-primary);">
                    // Get/Set value<br>
                    const picker = Metro.getPlugin('#myPicker', 'calendar-picker');<br>
                    picker.val('2026/12/25');<br>
                    const value = picker.val();<br><br>
                    // Time picker<br>
                    const tp = Metro.getPlugin('#myTime', 'time-picker');<br>
                    tp.val('14:30:00');<br>
                    const time = tp.time(); // {h: 14, m: 30, s: 0}
                </code>
            </div>

            <div class="divider"></div>

            <div class="form-group">
                <h4 style="font-size: 14px; margin-bottom: 12px;">Available Components:</h4>
                <ul class="feature-list">
                    <li>
                        <i class="fa-solid fa-circle-check"></i>
                        <span><strong>calendar-picker</strong> - Input + calendar popup (recommended)</span>
                    </li>
                    <li>
                        <i class="fa-solid fa-circle-check"></i>
                        <span><strong>date-picker</strong> - Wheel/scroll-based date selection</span>
                    </li>
                    <li>
                        <i class="fa-solid fa-circle-check"></i>
                        <span><strong>time-picker</strong> - Wheel/scroll-based time selection</span>
                    </li>
                    <li>
                        <i class="fa-solid fa-circle-check"></i>
                        <span><strong>calendar</strong> - Inline embedded calendar</span>
                    </li>
                </ul>
            </div>
        </div>
    </div>
</div>
@endsection
