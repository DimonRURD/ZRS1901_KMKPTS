import React, { Component } from "react";
import AuthService from "../services/auth.service";
import UserService from "../services/user.service";

export default class Profile extends Component {
  constructor(props) {
    super(props);

    this.state = {
      content: [],
    };
  }

  componentDidMount() {
    const currentUser = AuthService.getAccessToken();

    if (!currentUser) this.setState({ redirect: "/home" });
    this.setState({ currentUser: currentUser, userReady: true });

    UserService.getPackets().then((response) => {
      this.setState({
        content: response.data,
      });
    });
  }

  render() {
    const { content } = this.state;

    return (
      <div className="container">
        <div>
          <header className="jumbotron">
            <h3>Задетектированы данные пакеты:</h3>
          </header>
          <div
            style={{ display: "flex", flexDirection: "column", gap: "20px" }}
          >
            {content.map((item) => (
              <div
                style={{
                  backgroundColor: "#2795B9",
                  borderRadius: "10px",
                  padding: "10px",
                }}
              >
                <p>Timestamp: {item.timestamp}</p>
                <p>Длина пакета: {item.length}</p>
                <p>MAC адрес отправителя: {item.srcMAC}</p>
                <p>MAC адрес получателя: {item.dstMAC}</p>
                <p>IP адрес получателя: {item.srcIP}</p>
                <p>IP адрес получателя: {item.dstIP}</p>
                <p>Порт получателя: {item.srcPort}</p>
                <p>Порт получателя: {item.dstPort}</p>
                <p>Обнаруженные YARA правила: {item.matchedRules}</p>
              </div>
            ))}
          </div>
          <strong>Authorities:</strong>
        </div>
      </div>
    );
  }
}
