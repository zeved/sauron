import { Button, Chip, Grid, List, ListItem, Typography } from "@mui/material";
import axios from "axios";
import MemInfo from "./meminfo.component";
import HostInfo from "./hostinfo.component";
import ProcList from "./proclist.component";
import React from "react";
import NetConnList from "./netconnlist.component";

const NodeInfo = ({node, setNode}: any) => {

  const getNode = (nodeId: string) => {
    return axios.get(`http://localhost:8080/nodes/${nodeId}`, { headers: {"Access-Control-Allow-Origin": "*"}}).then(res => console.log(res.data))
  }

  const sendCommand = async(command: string) => {
    await axios.post('http://localhost:8080/nodes/cmd', {  nodeId: node.ID, command: command });
    await getNode(node.ID);
  }

  return (
    <>
      <h3>Node {node.ID}</h3>
      <Grid container spacing={2}>
          <Grid item xs={2}><Button color='success' variant='outlined' onClick={async () => await sendCommand('hostinfo')}>Refresh Host Info</Button></Grid>
          <Grid item xs={2}><Button color='success' variant='outlined' onClick={async () => await sendCommand('cpu')}>Refresh CPU Info</Button></Grid>
          <Grid item xs={2}><Button color='success' variant='outlined' onClick={async () => await sendCommand('mem')}>Refresh Memory Info</Button></Grid>
          <Grid item xs={2}><Button color='success' variant='outlined' onClick={async () => await sendCommand('ps')}>Refresh Process List</Button></Grid>
          <Grid item xs={2}><Button color='success' variant='outlined' onClick={async () => await sendCommand('netstat')}>Refresh Network Connections</Button></Grid>

      </Grid>
      <List>
        { Object.keys(node).map((k:any) => {
          let output = null;

          switch (k) {
            case 'MemInfo':
              output = <MemInfo memInfo={node[k]} />;
              break;
            case 'HostInfo':
              output = <HostInfo hostInfo={node[k]} />;
              break;
            case 'ProcessList':
              output = <ProcList procList={node[k]} />;
              break;
            case 'NetConnList':
              output = <NetConnList netConnList={node[k]} />;
              break;
            case 'LastResponse':
            case 'ID':
            case 'Topic':
              break;

            case 'LastHB':
              output = new Date(node[k] * 1000).toLocaleString();
              break;

            default: output = node[k];
          }

          return output && <ListItem key={k}>
            <Grid container spacing={2}>
            <Grid item xs={2}>
              <Typography variant="h6">{k}</Typography>
              {/* <Chip color="primary" label={k}>{k}</Chip> */}
            </Grid>
            <Grid item xs={10}>
              { React.isValidElement(output) ? output : <Chip variant='outlined' color='success' label={output} /> }
              {/* <Chip color='success'> label={output}</Chip>{output} */}
            </Grid>
            </Grid>
          </ListItem>;
         }) }
      </List>
    </>
  )
};

export default NodeInfo;