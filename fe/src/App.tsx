import React, { useEffect, useState } from 'react';
import './App.css';
import Nodes from './components/nodes.component';
import NodeInfo from './components/nodeinfo.component';

const App = () => {
  const [nodes, setNodes] = useState([]);
  const [selectedNode, setSelectedNode] = useState('');

  useEffect(() => {
    if (selectedNode !== '') {
      console.log(selectedNode);
    }
  }, [selectedNode]);

  return (
    <div className="App">
      <Nodes nodes={nodes} setNodes={setNodes} selectedNode={selectedNode} setSelectedNode={setSelectedNode}/>
      { selectedNode !== '' ? <NodeInfo node={nodes.find((node: any) => node.ID === selectedNode)} /> : 'no selected node'}
    </div>
  );
}

export default App;
