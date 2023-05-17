# wire graph

```sh
wire graph -injector initializeX -trim_prefixes github.com/qawatake/wire/_example/graph.
```

Generated graph.

```mermaid
flowchart BT;
	 0["context.Context"] -- "context.Context" --> 1["NewA"];
	 2["E"] -- "E" --> 3["F"];
	 2["E"] -- "E" --> 4["NewH"];
	 3["F"] -- "F" --> 5["NewX"];
	 6["H.G"] -- "G" --> 5["NewX"];
	 7["I2"] -- "I2" --> 8["NewD"];
	 1["NewA"] -- "I1" --> 9["NewB"];
	 1["NewA"] -- "I1" --> 4["NewH"];
	 9["NewB"] -- "B" --> 3["F"];
	 9["NewB"] -- "B" --> 5["NewX"];
	 8["NewD"] -- "D" --> 3["F"];
	 8["NewD"] -- "D" --> 5["NewX"];
	 4["NewH"] -- "H" --> 6["H.G"];
```

Raw output.

```
flowchart BT;
	 0["context.Context"] -- "context.Context" --> 1["NewA"];
	 2["E"] -- "E" --> 3["F"];
	 2["E"] -- "E" --> 4["NewH"];
	 3["F"] -- "F" --> 5["NewX"];
	 6["H.G"] -- "G" --> 5["NewX"];
	 7["I2"] -- "I2" --> 8["NewD"];
	 1["NewA"] -- "I1" --> 9["NewB"];
	 1["NewA"] -- "I1" --> 4["NewH"];
	 9["NewB"] -- "B" --> 3["F"];
	 9["NewB"] -- "B" --> 5["NewX"];
	 8["NewD"] -- "D" --> 3["F"];
	 8["NewD"] -- "D" --> 5["NewX"];
	 4["NewH"] -- "H" --> 6["H.G"];
```
