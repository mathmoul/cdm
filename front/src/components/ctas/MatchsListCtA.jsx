import React from 'react'

class MatchsListCtA extends React.Component {
    state = {
        list: [],
        loading: true
    };

    render() {
        const {list} = this.state;

        return (
            <div>
                {list.length === 0 && <h1>Pas de Matchs actuellement</h1>}
            </div>
        )
    }

}

export default MatchsListCtA