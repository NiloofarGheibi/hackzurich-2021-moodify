import React from "react";
import {VictoryChart, VictoryLine, VictoryScatter} from "victory";
import styled from '@emotion/styled';

const StyledPoint = styled.circle`
  fill: ${(props) => props.color};
`;

const colors = ["#A8E6CE", "#DCEDC2", "#FFD3B5", "#FFAAA6", "#FF8C94"];

const currentMonth = "Sep";

const ScatterPoint = ({ x, y, datum, min, max }) => {
    const i = React.useMemo(() => {
        return Math.floor(((datum.y - min) / (max - min)) * (colors.length - 1));
    }, [datum, min, max]);

    return <StyledPoint color={colors[i]} cx={x} cy={y} r={6} />;
};

export default function AdminView(){
    const data = [
        { x: "Jan", y: 43 },
        { x: "Feb", y: 44 },
        { x: "Mar", y: 47 },
        { x: "Apr", y: 51 },
        { x: "May", y: 57 },
        { x: "Jun", y: 62 },
        { x: "Jul", y: 67 },
        { x: "Aug", y: 68 },
        { x: "Sep", y: 63 },
        { x: "Oct", y: 54 },
        { x: "Nov", y: 47 },
        { x: "Dec", y: 42 }
    ];

    const temperatures = data.map(({ y }) => y);
    const min = Math.min(...temperatures);
    const max = Math.max(...temperatures);

    return (
        <VictoryChart>
            <VictoryLine data={data} />
            <VictoryScatter
                data={data}
                dataComponent={<ScatterPoint min={min} max={max} />}
            />
            <VictoryLine
                style={{
                    data: { stroke: "red", strokeWidth: 2 },
                    labels: { angle: -90, fill: "red", fontSize: 20 }
                }}
                // labels={["Important"]}
                // labelComponent={<VictoryLabel y={100}/>}
                x={() => "Sep"}
            />
        </VictoryChart>
    );
}