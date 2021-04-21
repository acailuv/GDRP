import React from "react";

export class UserRow extends React.Component {
  render() {
    return (
      <tr>
        <td>{this.props.id}</td>
        <td>{this.props.name}</td>
        <td>{this.props.secretKey}</td>
      </tr>
    );
  }
}
