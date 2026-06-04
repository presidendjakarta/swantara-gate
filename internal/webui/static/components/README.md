# Components

Reusable UI components untuk Swantara Gate Admin Panel.

## 📁 Structure

```
components/
├── sidebar.html          # Sidebar navigation
├── navbar.html           # Top navbar  
└── README.md             # This file
```

## 🚀 Usage

### 1. Add Container Elements

Di halaman HTML, tambahkan container untuk components:

```html
<div id="sidebar-container"></div>
<div id="navbar-container"></div>
```

### 2. Load Components dengan jQuery

```javascript
$(document).ready(function() {
    $('#sidebar-container').load('/components/sidebar.html');
    $('#navbar-container').load('/components/navbar.html');
});
```

### 3. Event Handling (Automatic)

✅ **Tidak perlu re-init!** Semua event handlers menggunakan **event delegation**:

```javascript
// Event delegation - bind sekali di document
$(document).on('click', '#themeToggle', function() { ... });
$(document).on('click', '#sidebarToggle', function() { ... });
$(document).on('click', '.nav-link', function() { ... });
```

**Keuntungan:**
- ⚡ Events automatically work untuk elements yang di-load dynamically
- 🔄 Tidak perlu re-init setelah component loaded
- 🧹 No duplicate event listeners
- 💪 Lebih scalable dan maintainable

## 🎨 Customization

Untuk customize components, edit langsung file HTML di folder `components/`.
