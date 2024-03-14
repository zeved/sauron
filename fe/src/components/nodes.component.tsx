import { MenuItem, Select } from "@mui/material";
import axios from "axios";
import { useEffect, useState } from "react";

const Nodes = ({ nodes, setNodes, selectedNode, setSelectedNode }: any) => {

  useEffect(() => {
    if (nodes.length === 0) {
      axios.get('http://localhost:8080/nodes', { headers: {"Access-Control-Allow-Origin": "*"}}).then(res => setNodes(res.data))
    }
  }, [nodes]);

  return (
    <>
      <Select label='Nodes' onChange={(e: any) => { console.log(e); setSelectedNode(e.target.value)}} value={selectedNode}>
        { nodes.length > 0 ? nodes.map((node:any) => <MenuItem key={node.ID} value={node.ID}>{node.ID}</MenuItem>) : <MenuItem>No nodes</MenuItem> }
      </Select>
    </>
  )
};

export default Nodes;