import React from "react";
import { Button, TextField } from "@material-ui/core";
import { Delete as DeleteRequest, GetAPILink } from "../_Utils/ApiConnector";
import "../global.css";

export class Delete extends React.Component {
  constructor(props) {
    super(props);
    this.state = {
      id: "",
      result: null,
    };
  }

  deleteButtonHandler(id) {
    GetAPILink().then((link) => {
      DeleteRequest(link + "/user/" + id).then((data) => {
        this.setState({
          result: data,
        });
      });
    });
  }

  render() {
    return (
      <div className="AppContent">
        <h1>Delete User</h1>
        <form>
          <TextField
            fullWidth
            type="number"
            id="user-id-input"
            label="User ID"
            onChange={(e) => this.setState({ id: e.target.value })}
            error={this.state.id === ""}
            value={this.state.id}
            helperText={this.state.id === "" ? "This cannot be empty!" : ""}
          />
          <Button
            variant="contained"
            color="primary"
            disabled={this.state.id === ""}
            onClick={() => this.deleteButtonHandler(this.state.id)}
          >
            Delete
          </Button>
        </form>
        {this.state.result !== null && <p>{this.state.result.msg}</p>}
      </div>
    );
  }
}
