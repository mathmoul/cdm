import React from 'react'
import {Button, Header, Icon, Image, Menu, Segment, Sidebar} from 'semantic-ui-react'

class SidebarNavigation extends React.Component {
    state = {
        visible: false
    };
    toggleVisibility = () => this.setState({visible: !this.state.visible});

    render() {
        const {visible} = this.state;
        return (
            <div>
                <Button onClick={this.toggleVisibility}> Test</Button>
                <Sidebar.Pushable as={Segment}>
                    <Sidebar as={Menu} animation='slide out' width='thin' visible={visible} icon='labeled' vertical
                             inverted>
                        <Menu.Item name='home'>
                            <Icon name='home'/>
                            Home
                        </Menu.Item>
                        <Menu.Item name='gamepad'>
                            <Icon name='gamepad'/>
                            Games
                        </Menu.Item>
                        <Menu.Item name='camera'>
                            <Icon name='camera'/>
                            Channels
                        </Menu.Item>
                    </Sidebar>
                    <Sidebar.Pusher>
                        <Segment basic>
                            <Header as='h3'>Application Content</Header>
                            <Image src='/assets/images/wireframe/paragraph.png'/>
                        </Segment>
                    </Sidebar.Pusher>
                </Sidebar.Pushable>
            </div>
        )
    }
}

export default SidebarNavigation