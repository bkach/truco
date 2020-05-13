import React from "react";
import List from "@material-ui/core/List";
import ListItem from "@material-ui/core/ListItem";
import ListItemText from "@material-ui/core/ListItemText";
import ListItemIcon from '@material-ui/core/ListItemIcon';
import AddIcon from '@material-ui/icons/Add';

class ItemList extends React.Component {
    constructor(props) {
        super(props);
        this.state = {
            items: props.items,
            onAdd: props.onAdd
        }
    }

    static getDerivedStateFromProps(props, state) {
        return {
            items: props.items,
            onAdd: props.onAdd
        }
    }

    render () {
        const listItems = this.state.items.map((item) =>
            <ListItem button key={item.id}>
                <ListItemText primary={item.text} />
            </ListItem>
        )

        return (
            <>
                <List component="nav">
                    {listItems}
                    <ListItem button key={"add button"}>
                        <ListItemIcon>
                            <AddIcon />
                        </ListItemIcon>
                        <ListItemText primary="Add Game" onClick={this.onAdd}/>
                    </ListItem>
                </List>
            </>
        );
    }
}

export default ItemList;
