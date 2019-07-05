import React from "react";


class Region extends React.Component {
    constructor(props) {
        super(props);
        this.state = {
            locals: []
        };
    }

    componentDidMount() {
        const {region} = this.props.match.params;
        fetch(`http://localhost:3000/locals/regions/${region}`)
        .then(res => res.json())
        .then(data => {
            this.setState({
                locals: data
            });
        })
    }

    render() {
        return (
            <div className="Region">
                {this.state.locals.map( (local, index) => (
                    <div key={index}>
                        {local.name}
                    </div>
                ))}
            </div>
        );
    }
}

export default Region;