import cli from "command-line-args";

export type ProjectorOptions = {
    args?: string[];
    pwd?: string; // projector --pwd ...
    config?: string; // projector --config ...
};

export default function getOptions(): ProjectorOptions {
    return cli([
        {name: 'args', type: String, defaultOption: true, multiple: true},
        {name: 'config', alias: 'c', type: String},
        {name: 'pwd', alias: 'p', type: String},
    ]) as ProjectorOptions;
};
