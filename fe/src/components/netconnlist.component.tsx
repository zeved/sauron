import { Grid, Table, TableBody, TableCell, TableHead, TableRow } from "@mui/material";
import { DataGrid, GridColDef } from "@mui/x-data-grid";

const NetConnList = ({netConnList}: any) => {
  const connections = JSON.parse(netConnList);

  const columns: GridColDef[] = [
    { field: 'fd', headerName: 'FD', width: 100 },
    { field: 'type', headerName: 'Type', width: 100 },
    { field: 'laddr', headerName: 'Local Address', width: 200 },
    { field: 'raddr', headerName: 'Remote Address', width: 200 },
    { field: 'state', headerName: 'State', width: 100 },
    { field: 'pid', headerName: 'PID', width: 100 },
    { field: 'uids', headerName: 'UIDs', width: 100 },
  ];

  const rows = connections.map((p: any) => {
    return {
      id: p.fd,
      fd: p.fd,
      type: p.type,
      laddr: p.localaddr.ip + ':' + p.localaddr.port,
      raddr: p.remoteaddr.ip + ':' + p.remoteaddr.port,
      state: p.status,
      uids: p.uids,
      pid: p.pid
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
        {/* <Table>
          <TableHead>
            <TableRow>
              { Object.keys(procs[0]).map((k:any) => <th key={k}>{k}</th>) }
            </TableRow>
          </TableHead>
          <TableBody>
          { procs.map((p: any) => <TableRow key={p.pid}>
              <TableCell>{p.pid}</TableCell>
              <TableCell>{p.ppid}</TableCell>
              <TableCell>{p.cmd.slice(0, 32)}</TableCell>
              <TableCell>{new Date(p.create_time).toLocaleString()}</TableCell>
              <TableCell>{p.status}</TableCell>
              <TableCell>{p.username}</TableCell>
              <TableCell>{p.cpu}</TableCell>
              <TableCell>{p.mem}</TableCell>
            </TableRow>) }
          </TableBody>
        </Table> */}
      </Grid>
    </>
  )
}

export default NetConnList;