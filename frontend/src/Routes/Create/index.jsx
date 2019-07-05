import React from "react";

class Create extends React.Component {
    constructor(props) {
        super(props);
        this.state = {
            name: "",
            day: "",
            location: "",
            region: "",
            organizers: ""
        }
    }

    handleValueChange(value, field) {
        let state = this.state;
        state[field] = value;
        this.setState(state);
    }

    render() {
        return(
            <div style={{display: "flex", flexDirection: "column", maxWidth: "200px", margin: "auto"}}>
                <input type="text" onChange={e => this.handleValueChange(e.target.value, "name")} value={this.state.name} placeholder="name"/>
                <input type="text" onChange={e => this.handleValueChange(e.target.value, "day")} name={this.state.day} placeholder="day"/>
                <input type="text" onChange={e => this.handleValueChange(e.target.value, "location")} name={this.state.location} placeholder="location"/>
                <input type="text" onChange={e => this.handleValueChange(e.target.value, "region")} name={this.state.region} placeholder="region"/>
                <input type="text" onChange={e => this.handleValueChange(e.target.value, "organizers")} name={this.state.organizers} placeholder="organizers"/>
                <input type="submit" />
            </div>
        );
    }
}

export default Create;