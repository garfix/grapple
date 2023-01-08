package nn

type Layer struct {
	nodes []*Node
}

// class NeuronLayer:
//     def __init__(self, num_neurons, bias):

//         # Every neuron in a layer shares the same bias
//         self.bias = bias if bias else random.random()

//         self.neurons = []
//         for i in range(num_neurons):
//             self.neurons.append(Neuron(self.bias))

func CreateLayer(nodeCount int, bias float64) *Layer {

	nodes := []*Node{}

	for i := 0; i < nodeCount; i++ {
		nodes = append(nodes, createNode(bias))
	}

	return &Layer{
		nodes: nodes,
	}
}

//     def inspect(self):
//         print('Neurons:', len(self.neurons))
//         for n in range(len(self.neurons)):
//             print(' Neuron', n)
//             for w in range(len(self.neurons[n].weights)):
//                 print('  Weight:', self.neurons[n].weights[w])
//             print('  Bias:', self.bias)

//     def feed_forward(self, inputs):
//         outputs = []
//         for neuron in self.neurons:
//             outputs.append(neuron.calculate_output(inputs))
//         return outputs

func (layer *Layer) feedForward(inputs []float64) []float64 {
	outputs := []float64{}
	for _, node := range layer.nodes {
		outputs = append(outputs, node.calculateOutput(inputs))
	}
	return outputs
}

//     def get_outputs(self):
//         outputs = []
//         for neuron in self.neurons:
//             outputs.append(neuron.output)
//         return outputs

func (layer *Layer) getOutputs() []float64 {
	outputs := []float64{}
	for _, node := range layer.nodes {
		outputs = append(outputs, node.output)
	}
	return outputs
}
