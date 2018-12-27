import React from "react";


class menu extends React.Component {

    componentWillMount() {
		const tree = [
            {
                name: "menu1", 
                id: 1, 
                isOpen: true, 
                customComponent: InfinityMenu,
                children: [
                    {
                        name: "submenu1",
                        id: 1,
                        isOpen: true,
                        customComponent: InfinityMenu
                    
                    },
                    {
                        name: "submenu2",
                        id: 2,
                        isOpen: true,
                        customComponent: InfinityMenu
                        
                    },
                    {
                        name: "submenu1",
                        id: 1,
                        isOpen: true,
                        customComponent: InfinityMenu
                    
                    },
                ]
            }        
    ];

    this.setState({
        tree: tree
    });
}

onNodeMouseClick(event, tree) {
    this.setState({
        tree: tree
    });
}

render() {
    return (
        <Menu
            tree={this.state.tree}
            onNodeMouseClick={this.onNodeMouseClick.bind(this)}
            maxLeaves={2}
        />
    );
}
}
ReactDOM.render(<menu />, document.getElementById("example"));