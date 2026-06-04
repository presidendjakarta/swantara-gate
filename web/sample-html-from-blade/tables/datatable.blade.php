@extends('layouts.app')

@section('title', 'DataTable Examples')

@section('content')
<div class="page-header">
    <div>
        <h1>DataTable Examples</h1>
        <p>Various DataTable configurations with DataTables.net plugin</p>
    </div>
</div>

<!-- 1. Basic Datatable -->
<div class="content-card" style="margin-bottom: 24px;">
    <div class="card-header">
        <div class="card-header-left">
            <div class="card-icon bg-primary">
                <i class="fa-solid fa-table"></i>
            </div>
            <div>
                <h3>1. Basic Datatable</h3>
                <p class="card-subtitle">Standard DataTable with search, sort, and pagination</p>
            </div>
        </div>
    </div>
    <div class="card-body no-padding">
        <div class="data-table-wrapper">
            <table id="basicTable" class="display" style="width:100%">
                <thead>
                    <tr>
                        <th>Name</th>
                        <th>Position</th>
                        <th>Office</th>
                        <th>Age</th>
                        <th>Start Date</th>
                        <th>Salary</th>
                    </tr>
                </thead>
                <tbody>
                    <tr>
                        <td>Tiger Nixon</td>
                        <td>System Architect</td>
                        <td>Edinburgh</td>
                        <td>61</td>
                        <td>2011/04/25</td>
                        <td>$320,800</td>
                    </tr>
                    <tr>
                        <td>Garrett Winters</td>
                        <td>Accountant</td>
                        <td>Tokyo</td>
                        <td>63</td>
                        <td>2011/07/25</td>
                        <td>$170,750</td>
                    </tr>
                    <tr>
                        <td>Ashton Cox</td>
                        <td>Junior Technical Author</td>
                        <td>San Francisco</td>
                        <td>66</td>
                        <td>2009/01/12</td>
                        <td>$86,000</td>
                    </tr>
                    <tr>
                        <td>Cedric Kelly</td>
                        <td>Senior Javascript Developer</td>
                        <td>Edinburgh</td>
                        <td>22</td>
                        <td>2012/03/29</td>
                        <td>$433,060</td>
                    </tr>
                    <tr>
                        <td>Airi Satou</td>
                        <td>Accountant</td>
                        <td>Tokyo</td>
                        <td>33</td>
                        <td>2008/11/28</td>
                        <td>$162,700</td>
                    </tr>
                    <tr>
                        <td>Brielle Williamson</td>
                        <td>Integration Specialist</td>
                        <td>New York</td>
                        <td>61</td>
                        <td>2012/12/02</td>
                        <td>$372,000</td>
                    </tr>
                    <tr>
                        <td>Herrod Chandler</td>
                        <td>Sales Assistant</td>
                        <td>San Francisco</td>
                        <td>59</td>
                        <td>2012/08/06</td>
                        <td>$137,500</td>
                    </tr>
                    <tr>
                        <td>Rhona Davidson</td>
                        <td>Integration Specialist</td>
                        <td>Tokyo</td>
                        <td>55</td>
                        <td>2010/10/14</td>
                        <td>$327,900</td>
                    </tr>
                </tbody>
            </table>
        </div>
    </div>
</div>

<!-- 4. Alternative Pagination Datatable -->
<div class="content-card" style="margin-bottom: 24px;">
    <div class="card-header">
        <div class="card-header-left">
            <div class="card-icon bg-info">
                <i class="fa-solid fa-forward"></i>
            </div>
            <div>
                <h3>4. Alternative Pagination Datatable</h3>
                <p class="card-subtitle">Full numbers pagination with first/last buttons</p>
            </div>
        </div>
    </div>
    <div class="card-body no-padding">
        <div class="data-table-wrapper">
            <table id="altPaginationTable" class="display" style="width:100%">
                <thead>
                    <tr>
                        <th>Name</th>
                        <th>Position</th>
                        <th>Office</th>
                        <th>Age</th>
                        <th>Start Date</th>
                        <th>Salary</th>
                    </tr>
                </thead>
                <tbody>
                    <tr><td>Tiger Nixon</td><td>System Architect</td><td>Edinburgh</td><td>61</td><td>2011/04/25</td><td>$320,800</td></tr>
                    <tr><td>Garrett Winters</td><td>Accountant</td><td>Tokyo</td><td>63</td><td>2011/07/25</td><td>$170,750</td></tr>
                    <tr><td>Ashton Cox</td><td>Junior Technical Author</td><td>San Francisco</td><td>66</td><td>2009/01/12</td><td>$86,000</td></tr>
                    <tr><td>Cedric Kelly</td><td>Senior Javascript Developer</td><td>Edinburgh</td><td>22</td><td>2012/03/29</td><td>$433,060</td></tr>
                    <tr><td>Airi Satou</td><td>Accountant</td><td>Tokyo</td><td>33</td><td>2008/11/28</td><td>$162,700</td></tr>
                    <tr><td>Brielle Williamson</td><td>Integration Specialist</td><td>New York</td><td>61</td><td>2012/12/02</td><td>$372,000</td></tr>
                    <tr><td>Herrod Chandler</td><td>Sales Assistant</td><td>San Francisco</td><td>59</td><td>2012/08/06</td><td>$137,500</td></tr>
                    <tr><td>Rhona Davidson</td><td>Integration Specialist</td><td>Tokyo</td><td>55</td><td>2010/10/14</td><td>$327,900</td></tr>
                    <tr><td>Colleen Hurst</td><td>Javascript Developer</td><td>San Francisco</td><td>39</td><td>2009/09/15</td><td>$205,500</td></tr>
                    <tr><td>Sonya Frost</td><td>Software Engineer</td><td>Edinburgh</td><td>23</td><td>2008/12/13</td><td>$103,600</td></tr>
                    <tr><td>Jena Gaines</td><td>Office Manager</td><td>London</td><td>30</td><td>2008/12/19</td><td>$90,560</td></tr>
                    <tr><td>Quinn Flynn</td><td>Support Lead</td><td>Edinburgh</td><td>22</td><td>2013/03/03</td><td>$342,000</td></tr>
                </tbody>
            </table>
        </div>
    </div>
</div>

<!-- 5. Buttons Datatables -->
<div class="content-card" style="margin-bottom: 24px;">
    <div class="card-header">
        <div class="card-header-left">
            <div class="card-icon bg-danger">
                <i class="fa-solid fa-file-export"></i>
            </div>
            <div>
                <h3>5. Buttons Datatables</h3>
                <p class="card-subtitle">Export to Excel, PDF, CSV, Copy, and Print</p>
            </div>
        </div>
    </div>
    <div class="card-body no-padding">
        <div class="data-table-wrapper">
            <table id="buttonsTable" class="display" style="width:100%">
                <thead>
                    <tr>
                        <th>Name</th>
                        <th>Position</th>
                        <th>Office</th>
                        <th>Age</th>
                        <th>Start Date</th>
                        <th>Salary</th>
                    </tr>
                </thead>
                <tbody>
                    <tr><td>Tiger Nixon</td><td>System Architect</td><td>Edinburgh</td><td>61</td><td>2011/04/25</td><td>$320,800</td></tr>
                    <tr><td>Garrett Winters</td><td>Accountant</td><td>Tokyo</td><td>63</td><td>2011/07/25</td><td>$170,750</td></tr>
                    <tr><td>Ashton Cox</td><td>Junior Technical Author</td><td>San Francisco</td><td>66</td><td>2009/01/12</td><td>$86,000</td></tr>
                    <tr><td>Cedric Kelly</td><td>Senior Javascript Developer</td><td>Edinburgh</td><td>22</td><td>2012/03/29</td><td>$433,060</td></tr>
                    <tr><td>Airi Satou</td><td>Accountant</td><td>Tokyo</td><td>33</td><td>2008/11/28</td><td>$162,700</td></tr>
                    <tr><td>Brielle Williamson</td><td>Integration Specialist</td><td>New York</td><td>61</td><td>2012/12/02</td><td>$372,000</td></tr>
                    <tr><td>Herrod Chandler</td><td>Sales Assistant</td><td>San Francisco</td><td>59</td><td>2012/08/06</td><td>$137,500</td></tr>
                    <tr><td>Rhona Davidson</td><td>Integration Specialist</td><td>Tokyo</td><td>55</td><td>2010/10/14</td><td>$327,900</td></tr>
                    <tr><td>Colleen Hurst</td><td>Javascript Developer</td><td>San Francisco</td><td>39</td><td>2009/09/15</td><td>$205,500</td></tr>
                    <tr><td>Sonya Frost</td><td>Software Engineer</td><td>Edinburgh</td><td>23</td><td>2008/12/13</td><td>$103,600</td></tr>
                </tbody>
            </table>
        </div>
    </div>
</div>

<!-- 6. Ajax ServerSide Datatable -->
<div class="content-card" style="margin-bottom: 24px;">
    <div class="card-header">
        <div class="card-header-left">
            <div class="card-icon bg-secondary">
                <i class="fa-solid fa-server"></i>
            </div>
            <div>
                <h3>6. Ajax ServerSide Datatable</h3>
                <p class="card-subtitle">Server-side processing with Ajax data source</p>
            </div>
        </div>
    </div>
    <div class="card-body">
        <div class="alert alert-info" style="margin-bottom: 20px; padding: 16px; background: var(--accent-light); border-left: 4px solid var(--accent); border-radius: var(--radius-md);">
            <i class="fa-solid fa-circle-info" style="color: var(--accent); margin-right: 8px;"></i>
            <strong>Note:</strong> This example shows the configuration for server-side processing. You need to create an API endpoint that returns JSON data in DataTables format.
        </div>
        <div class="data-table-wrapper">
            <table id="ajaxTable" class="display" style="width:100%">
                <thead>
                    <tr>
                        <th>ID</th>
                        <th>Name</th>
                        <th>Position</th>
                        <th>Office</th>
                        <th>Start Date</th>
                        <th>Salary</th>
                    </tr>
                </thead>
                <tbody>
                    <tr>
                        <td colspan="6" style="text-align: center; padding: 40px; color: var(--text-tertiary);">
                            <i class="fa-solid fa-circle-info" style="font-size: 48px; margin-bottom: 12px; display: block;"></i>
                            Loading data from server...
                        </td>
                    </tr>
                </tbody>
            </table>
        </div>
        <div class="divider" style="margin: 20px 0;"></div>
        <div style="background: var(--bg-secondary); padding: 16px; border-radius: var(--radius-lg);">
            <h4 style="margin-bottom: 12px; font-size: 14px;">Example API Response Format:</h4>
            <pre style="background: var(--surface); padding: 16px; border-radius: var(--radius-md); overflow-x: auto; font-size: 12px;"><code>{
    "draw": 1,
    "recordsTotal": 57,
    "recordsFiltered": 57,
    "data": [
        {
            "id": 1,
            "name": "Tiger Nixon",
            "position": "System Architect",
            "office": "Edinburgh",
            "start_date": "2011/04/25",
            "salary": "$320,800"
        }
    ]
}</code></pre>
        </div>
    </div>
</div>
@endsection

@push('scripts')
<script>
$(function(){
    if($.fn.DataTable) {
        
        // 1. Basic Datatable
        $('#basicTable').DataTable({
            pageLength: 5,
            lengthMenu: [5, 10, 25, 50],
            language: {
                search: "Search:",
                lengthMenu: "Show _MENU_ entries",
                info: "Showing _START_ to _END_ of _TOTAL_ entries",
                infoEmpty: "No data available",
                zeroRecords: "No data found",
                paginate: {
                    previous: "Previous",
                    next: "Next"
                }
            }
        });
        
        // 3. Alternative Pagination Datatable
        $('#altPaginationTable').DataTable({
            pagingType: "full_numbers",
            pageLength: 5,
            lengthMenu: [5, 10, 25],
            language: {
                search: "Search:",
                lengthMenu: "Show _MENU_ entries",
                info: "Showing _START_ to _END_ of _TOTAL_ entries",
                infoEmpty: "No data available",
                zeroRecords: "No data found",
                paginate: {
                    first: "First",
                    previous: "Previous",
                    next: "Next",
                    last: "Last"
                }
            }
        });
        
        // 5. Buttons Datatables
        $('#buttonsTable').DataTable({
            dom: '<"button-row"B><"top-row"lf>rtip',
            pageLength: 10,
            buttons: [
                { extend: 'copy', text: '<i class="fa-solid fa-copy"></i> Copy' },
                { extend: 'csv', text: '<i class="fa-solid fa-file-csv"></i> CSV' },
                { extend: 'excel', text: '<i class="fa-solid fa-file-excel"></i> Excel' },
                { extend: 'pdf', text: '<i class="fa-solid fa-file-pdf"></i> PDF', orientation: 'landscape' },
                { extend: 'print', text: '<i class="fa-solid fa-print"></i> Print' }
            ],
            language: {
                search: "Search:",
                lengthMenu: "Show _MENU_ entries",
                info: "Showing _START_ to _END_ of _TOTAL_ entries",
                infoEmpty: "No data available",
                zeroRecords: "No data found"
            }
        });
        
        // 6. Ajax ServerSide Datatable (Example Configuration)
        $('#ajaxTable').DataTable({
            processing: true,
            serverSide: true,
            ajax: {
                url: "/api/employees",
                type: "GET",
                error: function() {
                    $('#ajaxTable tbody').html('<tr><td colspan="6" style="text-align: center; padding: 40px; color: var(--danger);"><i class="fa-solid fa-triangle-exclamation" style="font-size: 48px; margin-bottom: 12px; display: block;"></i>Error loading data. Please configure the API endpoint.</td></tr>');
                }
            },
            columns: [
                { data: "id" },
                { data: "name" },
                { data: "position" },
                { data: "office" },
                { data: "start_date" },
                { data: "salary" }
            ],
            pageLength: 10,
            language: {
                search: "Search:",
                lengthMenu: "Show _MENU_ entries",
                processing: "<i class='fa-solid fa-spinner fa-spin'></i> Processing...",
                info: "Showing _START_ to _END_ of _TOTAL_ entries",
                infoEmpty: "No data available",
                zeroRecords: "No matching records found"
            }
        });
        
    }
});
</script>
@endpush
