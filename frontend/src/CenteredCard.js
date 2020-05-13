import React from "react";
import Card from "@material-ui/core/Card";
import makeStyles from "@material-ui/core/styles/makeStyles";

function CenteredCard(props) {
    const styles  = makeStyles((theme) => ({
       cardStyle: {
           width: "100%",
           height: "100vh",
           padding: "10%",
       }
    }))

    return (
        <>
            <Card style={styles.cardStyle}>
                {props.children}
            </Card>
        </>
    );
}

export default CenteredCard;

