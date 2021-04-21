import React from "react";
import { Button, TextField } from "@material-ui/core";
import "../global.css";
import { Post } from "../_Utils/ApiCaller";

export class Upsert extends React.Component {
  constructor(props) {
    super(props);
    this.state = {
      id: "",
      name: "",
      secretKey: "",
      formInvalid: true,
      result: null,
    };
  }

  upsertButtonHandler() {
    const requestBody = {
      id: this.state.id === "" ? null : parseInt(this.state.id),
      name: this.state.name,
      secret_key: this.state.secretKey,
    };

    Post("http://localhost:5000/upsert", requestBody).then((data) => {
      this.setState({
        result: data,
      });
    });

    this.setState({
        id: "",
        name: "",
        secretKey: "",
    })
  }

  isInputInvalid() {
    if (
      this.state.id === "" &&
      this.state.name === "" &&
      this.state.secretKey === ""
    ) {
      this.setState({
        formInvalid: true,
      });
      return;
    }

    if (
      this.state.id === "" &&
      (this.state.name === "" || this.state.secretKey === "")
    ) {
      this.setState({
        formInvalid: true,
      });
      return;
    }

    if (
      this.state.id !== "" &&
      this.state.name === "" &&
      this.state.secretKey === ""
    ) {
      this.setState({
        formInvalid: true,
      });
      return;
    }

    this.setState({
      formInvalid: false,
    });
    return;
  }

  render() {
    return (
      <div className="AppContent">
        <h1>Upsert User</h1>
        <form>
          <TextField
            fullWidth
            type="number"
            id="user-id-input"
            label="User ID"
            onChange={(e) => {
              this.setState({ id: e.target.value }, this.isInputInvalid);
            }}
            value={this.state.id}
            error={this.state.formInvalid}
            helperText={"Skip this field to insert. Fill this field to update."}
          />
          <TextField
            fullWidth
            type="text"
            id="user-name-input"
            label="Name"
            onChange={(e) => {
              this.setState({ name: e.target.value }, this.isInputInvalid);
            }}
            value={this.state.name}
            error={this.state.formInvalid}
          />
          <TextField
            fullWidth
            type="text"
            id="user-secret-key-input"
            label="Secret Key"
            onChange={(e) => {
              this.setState({ secretKey: e.target.value }, this.isInputInvalid);
            }}
            value={this.state.secretKey}
            error={this.state.formInvalid}
          />
          <Button
            variant="contained"
            color="primary"
            disabled={this.state.formInvalid}
            onClick={() => this.upsertButtonHandler(this.state.id)}
          >
            Upsert
          </Button>
        </form>
        {this.state.result !== null && <p>{this.state.result.msg}</p>}
      </div>
    );
  }
}
