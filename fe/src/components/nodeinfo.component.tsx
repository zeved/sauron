import { List, ListItem } from "@mui/material";

const NodeInfo = ({node}: any) => {

  return (
    <div>
      <List>
        { Object.keys(node).map((k:any) => <ListItem>{k} - {node[k]}</ListItem>) }
      </List>
    </div>
  )
};

export default NodeInfo;