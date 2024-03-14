import { Grid } from "@mui/material";

const HostInfo = ({hostInfo}: any) => {
  const host = JSON.parse(hostInfo);
  return (
    <>
      <Grid item xs={12}>
        &nbsp;Hostname: {host?.hostname} | OS: {host?.os} | Arch: {host?.kernelArch} | Kernel: {host?.kernelVersion}
      </Grid>
    </>
  )
}

export default HostInfo;