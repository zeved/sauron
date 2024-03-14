import { Grid } from "@mui/material";
import { DataGrid, GridColDef } from "@mui/x-data-grid";

const ProcList = ({procList}: any) => {
  const procs = JSON.parse(procList);

  const columns: GridColDef[] = [
    { field: 'pid', headerName: 'PID', width: 100 },
    { field: 'ppid', headerName: 'PPID', width: 100 },
    { field: 'username', headerName: 'Username', width: 100 },
    { field: 'cmd', headerName: 'Command', width: 500 },
    { field: 'cpu', headerName: 'CPU', width: 100 },
    { field: 'mem', headerName: 'Memory', width: 100 },
    { field: 'status', headerName: 'Status', width: 100 },
    { field: 'create_time', headerName: 'Create Time', width: 200 },
  ];

  const rows = procs.map((p: any) => {
    return {
      id: p.pid,
      pid: p.pid,
      ppid: p.ppid,
      username: p.username,
      cmd: p.cmd,
      cpu: p.cpu,
      mem: p.mem,
      status: p.status,
      create_time: new Date(p.create_time).toLocaleString()
    }
  });

  return (
    <>
      <Grid item xs={12}>
        <DataGrid columns={columns} rows={rows} initialState={{
          pagination: {
            paginationModel: {
              pageSize: 5,
            },
          },
        }}
        pageSizeOptions={[5]} autoHeight />
      </Grid>
    </>
  )
}

export default ProcList;