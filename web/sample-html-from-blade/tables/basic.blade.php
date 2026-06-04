@extends('layouts.app')

@section('title', 'Basic Table')

@section('content')
<div class="page-header">
    <div>
        <h1>Basic Table</h1>
        <p>Simple table component untuk menampilkan data dalam format tabel</p>
    </div>
    <div style="display: flex; gap: 10px;">
        <button class="btn btn-secondary">
            <i class="fa-solid fa-book"></i>
            <span>Documentation</span>
        </button>
        <button class="btn btn-primary">
            <i class="fa-solid fa-plus"></i>
            <span>Add Data</span>
        </button>
    </div>
</div>

<!-- Info Alert -->
<div class="content-card" style="margin-bottom: 24px; border-left: 4px solid var(--accent);">
    <div class="card-body" style="padding: 16px 20px;">
        <div style="display: flex; gap: 12px; align-items: start;">
            <i class="fa-solid fa-circle-info" style="color: var(--accent); font-size: 20px; margin-top: 2px;"></i>
            <div style="flex: 1;">
                <h4 style="margin-bottom: 4px; font-size: 14px;">What is Basic Table?</h4>
                <p style="font-size: 13px; color: var(--text-secondary); margin: 0;">Basic table menampilkan data dalam baris dan kolom sederhana tanpa fitur tambahan seperti sorting, searching, atau pagination. Cocok untuk dataset kecil.</p>
            </div>
        </div>
    </div>
</div>

<!-- Basic Table Example -->
<div class="content-card">
    <div class="card-header">
        <div class="card-header-left">
            <div class="card-icon bg-primary">
                <i class="fa-solid fa-table"></i>
            </div>
            <div>
                <h3>User Management</h3>
                <p class="card-subtitle">Daftar pengguna terdaftar</p>
            </div>
        </div>
    </div>
    <div class="card-body">
        <!-- Table Controls -->
        <div style="display: flex; justify-content: space-between; align-items: center; margin-bottom: 20px; flex-wrap: wrap; gap: 12px;">
            <div style="display: flex; align-items: center; gap: 8px;">
                <label style="font-size: 13px; color: var(--text-secondary);">Show</label>
                <select class="form-control" style="width: 70px; height: 36px; padding: 6px 12px;">
                    <option selected>10</option>
                    <option>25</option>
                    <option>50</option>
                    <option>100</option>
                </select>
                <label style="font-size: 13px; color: var(--text-secondary);">entries</label>
            </div>
            <div style="position: relative;">
                <i class="fa-solid fa-magnifying-glass" style="position: absolute; left: 12px; top: 50%; transform: translateY(-50%); color: var(--text-tertiary); font-size: 14px;"></i>
                <input type="text" class="form-control" placeholder="Search users..." style="padding-left: 38px; min-width: 250px; height: 36px;">
            </div>
        </div>

        <div class="table-responsive">
            <table class="table table-hover">
                <thead>
                    <tr>
                        <th style="width: 50px;">
                            <input type="checkbox" id="selectAll" style="cursor: pointer;">
                        </th>
                        <th style="width: 60px;">#</th>
                        <th>User</th>
                        <th>Email</th>
                        <th>Phone</th>
                        <th>Role</th>
                        <th>Status</th>
                        <th>Created</th>
                        <th style="width: 150px;">Actions</th>
                    </tr>
                </thead>
                <tbody>
                    <tr>
                        <td><input type="checkbox" class="row-checkbox" style="cursor: pointer;"></td>
                        <td>1</td>
                        <td>
                            <div style="display: flex; align-items: center; gap: 12px;">
                                <img src="https://ui-avatars.com/api/?name=Admin+User&background=0078D4&color=fff&size=40" style="border-radius: 50%; width: 40px; height: 40px;" alt="Avatar">
                                <div>
                                    <div style="font-weight: 600; margin-bottom: 2px;">Admin User</div>
                                    <div style="font-size: 12px; color: var(--text-tertiary);">@admin</div>
                                </div>
                            </div>
                        </td>
                        <td>
                            <div style="display: flex; align-items: center; gap: 6px;">
                                <i class="fa-solid fa-envelope" style="color: var(--text-tertiary); font-size: 12px;"></i>
                                <span>admin@example.com</span>
                            </div>
                        </td>
                        <td>
                            <div style="display: flex; align-items: center; gap: 6px;">
                                <i class="fa-solid fa-phone" style="color: var(--text-tertiary); font-size: 12px;"></i>
                                <span>+62 812-3456-7890</span>
                            </div>
                        </td>
                        <td><span class="badge badge-primary"><i class="fa-solid fa-shield-halved" style="margin-right: 4px;"></i>Administrator</span></td>
                        <td><span class="badge badge-success"><i class="fa-solid fa-check-circle" style="margin-right: 4px;"></i>Active</span></td>
                        <td>
                            <div style="font-size: 13px;">2024-01-15</div>
                            <div style="font-size: 11px; color: var(--text-tertiary);">08:30 AM</div>
                        </td>
                        <td>
                            <div style="display: flex; gap: 6px;">
                                <button class="btn btn-sm btn-info" title="View Details">
                                    <i class="fa-solid fa-eye"></i>
                                </button>
                                <button class="btn btn-sm btn-secondary" title="Edit">
                                    <i class="fa-solid fa-pen"></i>
                                </button>
                                <button class="btn btn-sm btn-danger" title="Delete">
                                    <i class="fa-solid fa-trash"></i>
                                </button>
                            </div>
                        </td>
                    </tr>
                    <tr>
                        <td><input type="checkbox" class="row-checkbox" style="cursor: pointer;"></td>
                        <td>2</td>
                        <td>
                            <div style="display: flex; align-items: center; gap: 12px;">
                                <img src="https://ui-avatars.com/api/?name=Budi+Santoso&background=16C60C&color=fff&size=40" style="border-radius: 50%; width: 40px; height: 40px;" alt="Avatar">
                                <div>
                                    <div style="font-weight: 600; margin-bottom: 2px;">Budi Santoso</div>
                                    <div style="font-size: 12px; color: var(--text-tertiary);">@budi</div>
                                </div>
                            </div>
                        </td>
                        <td>
                            <div style="display: flex; align-items: center; gap: 6px;">
                                <i class="fa-solid fa-envelope" style="color: var(--text-tertiary); font-size: 12px;"></i>
                                <span>budi@example.com</span>
                            </div>
                        </td>
                        <td>
                            <div style="display: flex; align-items: center; gap: 6px;">
                                <i class="fa-solid fa-phone" style="color: var(--text-tertiary); font-size: 12px;"></i>
                                <span>+62 813-4567-8901</span>
                            </div>
                        </td>
                        <td><span class="badge badge-info"><i class="fa-solid fa-pen-to-square" style="margin-right: 4px;"></i>Editor</span></td>
                        <td><span class="badge badge-success"><i class="fa-solid fa-check-circle" style="margin-right: 4px;"></i>Active</span></td>
                        <td>
                            <div style="font-size: 13px;">2024-02-20</div>
                            <div style="font-size: 11px; color: var(--text-tertiary);">10:15 AM</div>
                        </td>
                        <td>
                            <div style="display: flex; gap: 6px;">
                                <button class="btn btn-sm btn-info" title="View Details">
                                    <i class="fa-solid fa-eye"></i>
                                </button>
                                <button class="btn btn-sm btn-secondary" title="Edit">
                                    <i class="fa-solid fa-pen"></i>
                                </button>
                                <button class="btn btn-sm btn-danger" title="Delete">
                                    <i class="fa-solid fa-trash"></i>
                                </button>
                            </div>
                        </td>
                    </tr>
                    <tr>
                        <td><input type="checkbox" class="row-checkbox" style="cursor: pointer;"></td>
                        <td>3</td>
                        <td>
                            <div style="display: flex; align-items: center; gap: 12px;">
                                <img src="https://ui-avatars.com/api/?name=Siti+Rahayu&background=FFB900&color=fff&size=40" style="border-radius: 50%; width: 40px; height: 40px;" alt="Avatar">
                                <div>
                                    <div style="font-weight: 600; margin-bottom: 2px;">Siti Rahayu</div>
                                    <div style="font-size: 12px; color: var(--text-tertiary);">@siti</div>
                                </div>
                            </div>
                        </td>
                        <td>
                            <div style="display: flex; align-items: center; gap: 6px;">
                                <i class="fa-solid fa-envelope" style="color: var(--text-tertiary); font-size: 12px;"></i>
                                <span>siti@example.com</span>
                            </div>
                        </td>
                        <td>
                            <div style="display: flex; align-items: center; gap: 6px;">
                                <i class="fa-solid fa-phone" style="color: var(--text-tertiary); font-size: 12px;"></i>
                                <span>+62 814-5678-9012</span>
                            </div>
                        </td>
                        <td><span class="badge badge-warning"><i class="fa-solid fa-pen-fancy" style="margin-right: 4px;"></i>Author</span></td>
                        <td><span class="badge badge-warning"><i class="fa-solid fa-clock" style="margin-right: 4px;"></i>Pending</span></td>
                        <td>
                            <div style="font-size: 13px;">2024-03-10</div>
                            <div style="font-size: 11px; color: var(--text-tertiary);">02:45 PM</div>
                        </td>
                        <td>
                            <div style="display: flex; gap: 6px;">
                                <button class="btn btn-sm btn-info" title="View Details">
                                    <i class="fa-solid fa-eye"></i>
                                </button>
                                <button class="btn btn-sm btn-secondary" title="Edit">
                                    <i class="fa-solid fa-pen"></i>
                                </button>
                                <button class="btn btn-sm btn-danger" title="Delete">
                                    <i class="fa-solid fa-trash"></i>
                                </button>
                            </div>
                        </td>
                    </tr>
                    <tr>
                        <td><input type="checkbox" class="row-checkbox" style="cursor: pointer;"></td>
                        <td>4</td>
                        <td>
                            <div style="display: flex; align-items: center; gap: 12px;">
                                <img src="https://ui-avatars.com/api/?name=Andi+Wijaya&background=E81123&color=fff&size=40" style="border-radius: 50%; width: 40px; height: 40px;" alt="Avatar">
                                <div>
                                    <div style="font-weight: 600; margin-bottom: 2px;">Andi Wijaya</div>
                                    <div style="font-size: 12px; color: var(--text-tertiary);">@andi</div>
                                </div>
                            </div>
                        </td>
                        <td>
                            <div style="display: flex; align-items: center; gap: 6px;">
                                <i class="fa-solid fa-envelope" style="color: var(--text-tertiary); font-size: 12px;"></i>
                                <span>andi@example.com</span>
                            </div>
                        </td>
                        <td>
                            <div style="display: flex; align-items: center; gap: 6px;">
                                <i class="fa-solid fa-phone" style="color: var(--text-tertiary); font-size: 12px;"></i>
                                <span>+62 815-6789-0123</span>
                            </div>
                        </td>
                        <td><span class="badge badge-secondary"><i class="fa-solid fa-user" style="margin-right: 4px;"></i>User</span></td>
                        <td><span class="badge badge-danger"><i class="fa-solid fa-times-circle" style="margin-right: 4px;"></i>Inactive</span></td>
                        <td>
                            <div style="font-size: 13px;">2024-04-05</div>
                            <div style="font-size: 11px; color: var(--text-tertiary);">09:20 AM</div>
                        </td>
                        <td>
                            <div style="display: flex; gap: 6px;">
                                <button class="btn btn-sm btn-info" title="View Details">
                                    <i class="fa-solid fa-eye"></i>
                                </button>
                                <button class="btn btn-sm btn-secondary" title="Edit">
                                    <i class="fa-solid fa-pen"></i>
                                </button>
                                <button class="btn btn-sm btn-danger" title="Delete">
                                    <i class="fa-solid fa-trash"></i>
                                </button>
                            </div>
                        </td>
                    </tr>
                    <tr>
                        <td><input type="checkbox" class="row-checkbox" style="cursor: pointer;"></td>
                        <td>5</td>
                        <td>
                            <div style="display: flex; align-items: center; gap: 12px;">
                                <img src="https://ui-avatars.com/api/?name=Dewi+Lestari&background=0078D4&color=fff&size=40" style="border-radius: 50%; width: 40px; height: 40px;" alt="Avatar">
                                <div>
                                    <div style="font-weight: 600; margin-bottom: 2px;">Dewi Lestari</div>
                                    <div style="font-size: 12px; color: var(--text-tertiary);">@dewi</div>
                                </div>
                            </div>
                        </td>
                        <td>
                            <div style="display: flex; align-items: center; gap: 6px;">
                                <i class="fa-solid fa-envelope" style="color: var(--text-tertiary); font-size: 12px;"></i>
                                <span>dewi@example.com</span>
                            </div>
                        </td>
                        <td>
                            <div style="display: flex; align-items: center; gap: 6px;">
                                <i class="fa-solid fa-phone" style="color: var(--text-tertiary); font-size: 12px;"></i>
                                <span>+62 816-7890-1234</span>
                            </div>
                        </td>
                        <td><span class="badge badge-info"><i class="fa-solid fa-pen-to-square" style="margin-right: 4px;"></i>Editor</span></td>
                        <td><span class="badge badge-success"><i class="fa-solid fa-check-circle" style="margin-right: 4px;"></i>Active</span></td>
                        <td>
                            <div style="font-size: 13px;">2024-05-12</div>
                            <div style="font-size: 11px; color: var(--text-tertiary);">11:30 AM</div>
                        </td>
                        <td>
                            <div style="display: flex; gap: 6px;">
                                <button class="btn btn-sm btn-info" title="View Details">
                                    <i class="fa-solid fa-eye"></i>
                                </button>
                                <button class="btn btn-sm btn-secondary" title="Edit">
                                    <i class="fa-solid fa-pen"></i>
                                </button>
                                <button class="btn btn-sm btn-danger" title="Delete">
                                    <i class="fa-solid fa-trash"></i>
                                </button>
                            </div>
                        </td>
                    </tr>
                    <tr>
                        <td><input type="checkbox" class="row-checkbox" style="cursor: pointer;"></td>
                        <td>6</td>
                        <td>
                            <div style="display: flex; align-items: center; gap: 12px;">
                                <img src="https://ui-avatars.com/api/?name=Rudi+Hermawan&background=8764B8&color=fff&size=40" style="border-radius: 50%; width: 40px; height: 40px;" alt="Avatar">
                                <div>
                                    <div style="font-weight: 600; margin-bottom: 2px;">Rudi Hermawan</div>
                                    <div style="font-size: 12px; color: var(--text-tertiary);">@rudi</div>
                                </div>
                            </div>
                        </td>
                        <td>
                            <div style="display: flex; align-items: center; gap: 6px;">
                                <i class="fa-solid fa-envelope" style="color: var(--text-tertiary); font-size: 12px;"></i>
                                <span>rudi@example.com</span>
                            </div>
                        </td>
                        <td>
                            <div style="display: flex; align-items: center; gap: 6px;">
                                <i class="fa-solid fa-phone" style="color: var(--text-tertiary); font-size: 12px;"></i>
                                <span>+62 817-8901-2345</span>
                            </div>
                        </td>
                        <td><span class="badge badge-success"><i class="fa-solid fa-headset" style="margin-right: 4px;"></i>Support</span></td>
                        <td><span class="badge badge-success"><i class="fa-solid fa-check-circle" style="margin-right: 4px;"></i>Active</span></td>
                        <td>
                            <div style="font-size: 13px;">2024-06-18</div>
                            <div style="font-size: 11px; color: var(--text-tertiary);">03:15 PM</div>
                        </td>
                        <td>
                            <div style="display: flex; gap: 6px;">
                                <button class="btn btn-sm btn-info" title="View Details">
                                    <i class="fa-solid fa-eye"></i>
                                </button>
                                <button class="btn btn-sm btn-secondary" title="Edit">
                                    <i class="fa-solid fa-pen"></i>
                                </button>
                                <button class="btn btn-sm btn-danger" title="Delete">
                                    <i class="fa-solid fa-trash"></i>
                                </button>
                            </div>
                        </td>
                    </tr>
                    <tr>
                        <td><input type="checkbox" class="row-checkbox" style="cursor: pointer;"></td>
                        <td>7</td>
                        <td>
                            <div style="display: flex; align-items: center; gap: 12px;">
                                <img src="https://ui-avatars.com/api/?name=Maya+Sari&background=00B294&color=fff&size=40" style="border-radius: 50%; width: 40px; height: 40px;" alt="Avatar">
                                <div>
                                    <div style="font-weight: 600; margin-bottom: 2px;">Maya Sari</div>
                                    <div style="font-size: 12px; color: var(--text-tertiary);">@maya</div>
                                </div>
                            </div>
                        </td>
                        <td>
                            <div style="display: flex; align-items: center; gap: 6px;">
                                <i class="fa-solid fa-envelope" style="color: var(--text-tertiary); font-size: 12px;"></i>
                                <span>maya@example.com</span>
                            </div>
                        </td>
                        <td>
                            <div style="display: flex; align-items: center; gap: 6px;">
                                <i class="fa-solid fa-phone" style="color: var(--text-tertiary); font-size: 12px;"></i>
                                <span>+62 818-9012-3456</span>
                            </div>
                        </td>
                        <td><span class="badge badge-primary"><i class="fa-solid fa-chart-simple" style="margin-right: 4px;"></i>Analyst</span></td>
                        <td><span class="badge badge-success"><i class="fa-solid fa-check-circle" style="margin-right: 4px;"></i>Active</span></td>
                        <td>
                            <div style="font-size: 13px;">2024-07-22</div>
                            <div style="font-size: 11px; color: var(--text-tertiary);">01:45 PM</div>
                        </td>
                        <td>
                            <div style="display: flex; gap: 6px;">
                                <button class="btn btn-sm btn-info" title="View Details">
                                    <i class="fa-solid fa-eye"></i>
                                </button>
                                <button class="btn btn-sm btn-secondary" title="Edit">
                                    <i class="fa-solid fa-pen"></i>
                                </button>
                                <button class="btn btn-sm btn-danger" title="Delete">
                                    <i class="fa-solid fa-trash"></i>
                                </button>
                            </div>
                        </td>
                    </tr>
                </tbody>
            </table>
        </div>

        <div class="divider"></div>

        <!-- Pagination & Info -->
        <div style="display: flex; justify-content: space-between; align-items: center; flex-wrap: wrap; gap: 12px;">
            <div style="font-size: 13px; color: var(--text-tertiary);">
                <i class="fa-solid fa-circle-info" style="margin-right: 6px;"></i>
                Showing <strong>1</strong> to <strong>7</strong> of <strong>7</strong> entries
            </div>
            <div style="display: flex; gap: 4px;">
                <button class="btn btn-sm btn-secondary" disabled>
                    <i class="fa-solid fa-chevron-left" style="margin-right: 4px;"></i>
                    Previous
                </button>
                <button class="btn btn-sm btn-primary">1</button>
                <button class="btn btn-sm btn-secondary" disabled>
                    Next
                    <i class="fa-solid fa-chevron-right" style="margin-left: 4px;"></i>
                </button>
            </div>
        </div>
    </div>
</div>

<!-- Product Inventory Table -->
<div class="content-card" style="margin-top: 24px;">
    <div class="card-header">
        <div class="card-header-left">
            <div class="card-icon bg-success">
                <i class="fa-solid fa-box"></i>
            </div>
            <div>
                <h3>Product Inventory</h3>
                <p class="card-subtitle">Daftar produk dan stok tersedia</p>
            </div>
        </div>
        <button class="btn btn-success">
            <i class="fa-solid fa-plus"></i>
            <span>Add Product</span>
        </button>
    </div>
    <div class="card-body">
        <div class="table-responsive">
            <table class="table table-hover">
                <thead>
                    <tr>
                        <th style="width: 60px;">#</th>
                        <th>Product</th>
                        <th>SKU</th>
                        <th>Category</th>
                        <th>Price</th>
                        <th>Stock</th>
                        <th>Status</th>
                        <th style="width: 150px;">Actions</th>
                    </tr>
                </thead>
                <tbody>
                    <tr>
                        <td>1</td>
                        <td>
                            <div style="display: flex; align-items: center; gap: 12px;">
                                <div style="width: 50px; height: 50px; background: linear-gradient(135deg, #667eea 0%, #764ba2 100%); border-radius: 8px; display: flex; align-items: center; justify-content: center;">
                                    <i class="fa-solid fa-laptop" style="color: white; font-size: 20px;"></i>
                                </div>
                                <div>
                                    <div style="font-weight: 600; margin-bottom: 2px;">Laptop ASUS ROG</div>
                                    <div style="font-size: 12px; color: var(--text-tertiary);">Intel i7, 16GB RAM</div>
                                </div>
                            </div>
                        </td>
                        <td><code style="background: var(--bg-secondary); padding: 4px 8px; border-radius: 4px; font-size: 12px;">SKU-LPT-001</code></td>
                        <td><span class="badge badge-primary"><i class="fa-solid fa-laptop" style="margin-right: 4px;"></i>Electronics</span></td>
                        <td>
                            <div style="font-weight: 600; color: var(--accent);">Rp 18.500.000</div>
                        </td>
                        <td>
                            <div style="display: flex; align-items: center; gap: 8px;">
                                <div style="flex: 1; height: 6px; background: var(--bg-secondary); border-radius: 3px; overflow: hidden;">
                                    <div style="width: 75%; height: 100%; background: linear-gradient(90deg, var(--success), #3dd84a); border-radius: 3px;"></div>
                                </div>
                                <span style="font-size: 12px; font-weight: 600;">45</span>
                            </div>
                        </td>
                        <td><span class="badge badge-success"><i class="fa-solid fa-check-circle" style="margin-right: 4px;"></i>In Stock</span></td>
                        <td>
                            <div style="display: flex; gap: 6px;">
                                <button class="btn btn-sm btn-info" title="View Details">
                                    <i class="fa-solid fa-eye"></i>
                                </button>
                                <button class="btn btn-sm btn-secondary" title="Edit">
                                    <i class="fa-solid fa-pen"></i>
                                </button>
                                <button class="btn btn-sm btn-danger" title="Delete">
                                    <i class="fa-solid fa-trash"></i>
                                </button>
                            </div>
                        </td>
                    </tr>
                    <tr>
                        <td>2</td>
                        <td>
                            <div style="display: flex; align-items: center; gap: 12px;">
                                <div style="width: 50px; height: 50px; background: linear-gradient(135deg, #f093fb 0%, #f5576c 100%); border-radius: 8px; display: flex; align-items: center; justify-content: center;">
                                    <i class="fa-solid fa-mobile-screen" style="color: white; font-size: 20px;"></i>
                                </div>
                                <div>
                                    <div style="font-weight: 600; margin-bottom: 2px;">iPhone 15 Pro</div>
                                    <div style="font-size: 12px; color: var(--text-tertiary);">256GB, Titanium</div>
                                </div>
                            </div>
                        </td>
                        <td><code style="background: var(--bg-secondary); padding: 4px 8px; border-radius: 4px; font-size: 12px;">SKU-PHN-002</code></td>
                        <td><span class="badge badge-primary"><i class="fa-solid fa-mobile-screen" style="margin-right: 4px;"></i>Smartphone</span></td>
                        <td>
                            <div style="font-weight: 600; color: var(--accent);">Rp 21.000.000</div>
                        </td>
                        <td>
                            <div style="display: flex; align-items: center; gap: 8px;">
                                <div style="flex: 1; height: 6px; background: var(--bg-secondary); border-radius: 3px; overflow: hidden;">
                                    <div style="width: 20%; height: 100%; background: linear-gradient(90deg, var(--warning), #ffd43d); border-radius: 3px;"></div>
                                </div>
                                <span style="font-size: 12px; font-weight: 600;">8</span>
                            </div>
                        </td>
                        <td><span class="badge badge-warning"><i class="fa-solid fa-exclamation-triangle" style="margin-right: 4px;"></i>Low Stock</span></td>
                        <td>
                            <div style="display: flex; gap: 6px;">
                                <button class="btn btn-sm btn-info" title="View Details">
                                    <i class="fa-solid fa-eye"></i>
                                </button>
                                <button class="btn btn-sm btn-secondary" title="Edit">
                                    <i class="fa-solid fa-pen"></i>
                                </button>
                                <button class="btn btn-sm btn-danger" title="Delete">
                                    <i class="fa-solid fa-trash"></i>
                                </button>
                            </div>
                        </td>
                    </tr>
                    <tr>
                        <td>3</td>
                        <td>
                            <div style="display: flex; align-items: center; gap: 12px;">
                                <div style="width: 50px; height: 50px; background: linear-gradient(135deg, #4facfe 0%, #00f2fe 100%); border-radius: 8px; display: flex; align-items: center; justify-content: center;">
                                    <i class="fa-solid fa-headphones" style="color: white; font-size: 20px;"></i>
                                </div>
                                <div>
                                    <div style="font-weight: 600; margin-bottom: 2px;">Sony WH-1000XM5</div>
                                    <div style="font-size: 12px; color: var(--text-tertiary);">Wireless, Noise Cancelling</div>
                                </div>
                            </div>
                        </td>
                        <td><code style="background: var(--bg-secondary); padding: 4px 8px; border-radius: 4px; font-size: 12px;">SKU-HDP-003</code></td>
                        <td><span class="badge badge-info"><i class="fa-solid fa-headphones" style="margin-right: 4px;"></i>Accessories</span></td>
                        <td>
                            <div style="font-weight: 600; color: var(--accent);">Rp 5.999.000</div>
                        </td>
                        <td>
                            <div style="display: flex; align-items: center; gap: 8px;">
                                <div style="flex: 1; height: 6px; background: var(--bg-secondary); border-radius: 3px; overflow: hidden;">
                                    <div style="width: 90%; height: 100%; background: linear-gradient(90deg, var(--success), #3dd84a); border-radius: 3px;"></div>
                                </div>
                                <span style="font-size: 12px; font-weight: 600;">120</span>
                            </div>
                        </td>
                        <td><span class="badge badge-success"><i class="fa-solid fa-check-circle" style="margin-right: 4px;"></i>In Stock</span></td>
                        <td>
                            <div style="display: flex; gap: 6px;">
                                <button class="btn btn-sm btn-info" title="View Details">
                                    <i class="fa-solid fa-eye"></i>
                                </button>
                                <button class="btn btn-sm btn-secondary" title="Edit">
                                    <i class="fa-solid fa-pen"></i>
                                </button>
                                <button class="btn btn-sm btn-danger" title="Delete">
                                    <i class="fa-solid fa-trash"></i>
                                </button>
                            </div>
                        </td>
                    </tr>
                    <tr>
                        <td>4</td>
                        <td>
                            <div style="display: flex; align-items: center; gap: 12px;">
                                <div style="width: 50px; height: 50px; background: linear-gradient(135deg, #fa709a 0%, #fee140 100%); border-radius: 8px; display: flex; align-items: center; justify-content: center;">
                                    <i class="fa-solid fa-keyboard" style="color: white; font-size: 20px;"></i>
                                </div>
                                <div>
                                    <div style="font-weight: 600; margin-bottom: 2px;">Logitech MX Keys</div>
                                    <div style="font-size: 12px; color: var(--text-tertiary);">Wireless, Backlit</div>
                                </div>
                            </div>
                        </td>
                        <td><code style="background: var(--bg-secondary); padding: 4px 8px; border-radius: 4px; font-size: 12px;">SKU-KBD-004</code></td>
                        <td><span class="badge badge-info"><i class="fa-solid fa-keyboard" style="margin-right: 4px;"></i>Accessories</span></td>
                        <td>
                            <div style="font-weight: 600; color: var(--accent);">Rp 1.850.000</div>
                        </td>
                        <td>
                            <div style="display: flex; align-items: center; gap: 8px;">
                                <div style="flex: 1; height: 6px; background: var(--bg-secondary); border-radius: 3px; overflow: hidden;">
                                    <div style="width: 0%; height: 100%; background: linear-gradient(90deg, var(--danger), #ff6b6b); border-radius: 3px;"></div>
                                </div>
                                <span style="font-size: 12px; font-weight: 600;">0</span>
                            </div>
                        </td>
                        <td><span class="badge badge-danger"><i class="fa-solid fa-times-circle" style="margin-right: 4px;"></i>Out of Stock</span></td>
                        <td>
                            <div style="display: flex; gap: 6px;">
                                <button class="btn btn-sm btn-info" title="View Details">
                                    <i class="fa-solid fa-eye"></i>
                                </button>
                                <button class="btn btn-sm btn-secondary" title="Edit">
                                    <i class="fa-solid fa-pen"></i>
                                </button>
                                <button class="btn btn-sm btn-danger" title="Delete">
                                    <i class="fa-solid fa-trash"></i>
                                </button>
                            </div>
                        </td>
                    </tr>
                    <tr>
                        <td>5</td>
                        <td>
                            <div style="display: flex; align-items: center; gap: 12px;">
                                <div style="width: 50px; height: 50px; background: linear-gradient(135deg, #30cfd0 0%, #330867 100%); border-radius: 8px; display: flex; align-items: center; justify-content: center;">
                                    <i class="fa-solid fa-monitor" style="color: white; font-size: 20px;"></i>
                                </div>
                                <div>
                                    <div style="font-weight: 600; margin-bottom: 2px;">LG UltraWide 34"</div>
                                    <div style="font-size: 12px; color: var(--text-tertiary);">4K, HDR, 144Hz</div>
                                </div>
                            </div>
                        </td>
                        <td><code style="background: var(--bg-secondary); padding: 4px 8px; border-radius: 4px; font-size: 12px;">SKU-MNT-005</code></td>
                        <td><span class="badge badge-primary"><i class="fa-solid fa-monitor" style="margin-right: 4px;"></i>Monitor</span></td>
                        <td>
                            <div style="font-weight: 600; color: var(--accent);">Rp 12.500.000</div>
                        </td>
                        <td>
                            <div style="display: flex; align-items: center; gap: 8px;">
                                <div style="flex: 1; height: 6px; background: var(--bg-secondary); border-radius: 3px; overflow: hidden;">
                                    <div style="width: 50%; height: 100%; background: linear-gradient(90deg, var(--info), #3dd8f4); border-radius: 3px;"></div>
                                </div>
                                <span style="font-size: 12px; font-weight: 600;">25</span>
                            </div>
                        </td>
                        <td><span class="badge badge-success"><i class="fa-solid fa-check-circle" style="margin-right: 4px;"></i>In Stock</span></td>
                        <td>
                            <div style="display: flex; gap: 6px;">
                                <button class="btn btn-sm btn-info" title="View Details">
                                    <i class="fa-solid fa-eye"></i>
                                </button>
                                <button class="btn btn-sm btn-secondary" title="Edit">
                                    <i class="fa-solid fa-pen"></i>
                                </button>
                                <button class="btn btn-sm btn-danger" title="Delete">
                                    <i class="fa-solid fa-trash"></i>
                                </button>
                            </div>
                        </td>
                    </tr>
                </tbody>
            </table>
        </div>

        <div class="divider"></div>

        <!-- Pagination & Info -->
        <div style="display: flex; justify-content: space-between; align-items: center; flex-wrap: wrap; gap: 12px;">
            <div style="font-size: 13px; color: var(--text-tertiary);">
                <i class="fa-solid fa-circle-info" style="margin-right: 6px;"></i>
                Showing <strong>1</strong> to <strong>5</strong> of <strong>5</strong> entries
            </div>
            <div style="display: flex; gap: 4px;">
                <button class="btn btn-sm btn-secondary" disabled>
                    <i class="fa-solid fa-chevron-left" style="margin-right: 4px;"></i>
                    Previous
                </button>
                <button class="btn btn-sm btn-primary">1</button>
                <button class="btn btn-sm btn-secondary" disabled>
                    Next
                    <i class="fa-solid fa-chevron-right" style="margin-left: 4px;"></i>
                </button>
            </div>
        </div>
    </div>
</div>
@endsection

@push('scripts')
<script>
// Select All Checkbox Functionality
document.addEventListener('DOMContentLoaded', function() {
    const selectAll = document.getElementById('selectAll');
    const rowCheckboxes = document.querySelectorAll('.row-checkbox');
    
    // Select all checkboxes
    if (selectAll) {
        selectAll.addEventListener('change', function() {
            rowCheckboxes.forEach(checkbox => {
                checkbox.checked = this.checked;
            });
        });
    }
    
    // Update select all state when individual checkbox changes
    rowCheckboxes.forEach(checkbox => {
        checkbox.addEventListener('change', function() {
            const allChecked = Array.from(rowCheckboxes).every(cb => cb.checked);
            const someChecked = Array.from(rowCheckboxes).some(cb => cb.checked);
            
            if (selectAll) {
                selectAll.checked = allChecked;
                selectAll.indeterminate = someChecked && !allChecked;
            }
        });
    });
});
</script>
@endpush
