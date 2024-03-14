import { Grid } from "@mui/material";

const MemInfo = ({memInfo}: any) => {
  const mem = JSON.parse(memInfo);
  return (
    <>
      <Grid item xs={12}>
        &nbsp;Total: {Math.floor(mem?.total / 1024 / 1024)} MB | Available: {Math.floor(mem?.available / 1024 / 1024)} MB
      </Grid>
    </>
  )
}

export default MemInfo;