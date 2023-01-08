package nn

import "math/rand"

type ThreeLayerNetwork struct {
	inputCount  int
	hiddenLayer *Layer
	outputLayer *Layer
}

const learningRate = 0.5

// class NeuralNetwork:
//     LEARNING_RATE = 0.5

//     def __init__(self, num_inputs, num_hidden, num_outputs, hidden_layer_weights = None, hidden_layer_bias = None, output_layer_weights = None, output_layer_bias = None):
//         self.num_inputs = num_inputs

//         self.hidden_layer = NeuronLayer(num_hidden, hidden_layer_bias)
//         self.output_layer = NeuronLayer(num_outputs, output_layer_bias)

//         self.init_weights_from_inputs_to_hidden_layer_neurons(hidden_layer_weights)
//         self.init_weights_from_hidden_layer_neurons_to_output_layer_neurons(output_layer_weights)

func CreateThreeLayerNetwork(inputCount int, hiddenCount int, outputCount int, hiddenBias float64, outputBias float64) *ThreeLayerNetwork {
	network := &ThreeLayerNetwork{
		inputCount:  inputCount,
		hiddenLayer: CreateLayer(hiddenCount, hiddenBias),
		outputLayer: CreateLayer(outputCount, outputBias),
	}

	network.initWeightsFromInputsToHiddenLayerNeurons()
	network.initWeightsFromHiddenLayerNeuronsToOutputLayerNeurons()

	return network
}

//     def init_weights_from_inputs_to_hidden_layer_neurons(self, hidden_layer_weights):
//         weight_num = 0
//         for h in range(len(self.hidden_layer.neurons)):
//             for i in range(self.num_inputs):
//                 if not hidden_layer_weights:
//                     self.hidden_layer.neurons[h].weights.append(random.random())
//                 else:
//                     self.hidden_layer.neurons[h].weights.append(hidden_layer_weights[weight_num])
//                 weight_num += 1

func (network *ThreeLayerNetwork) initWeightsFromInputsToHiddenLayerNeurons() {
	for h := 0; h < len(network.hiddenLayer.nodes); h++ {
		for i := 0; i < network.inputCount; i++ {
			network.hiddenLayer.nodes[h].weights = append(network.hiddenLayer.nodes[h].weights, rand.Float64())
		}
	}
}

//     def init_weights_from_hidden_layer_neurons_to_output_layer_neurons(self, output_layer_weights):
//         weight_num = 0
//         for o in range(len(self.output_layer.neurons)):
//             for h in range(len(self.hidden_layer.neurons)):
//                 if not output_layer_weights:
//                     self.output_layer.neurons[o].weights.append(random.random())
//                 else:
//                     self.output_layer.neurons[o].weights.append(output_layer_weights[weight_num])
//                 weight_num += 1

func (network *ThreeLayerNetwork) initWeightsFromHiddenLayerNeuronsToOutputLayerNeurons() {
	for o := 0; o < len(network.outputLayer.nodes); o++ {
		for h := 0; h < len(network.hiddenLayer.nodes); h++ {
			network.outputLayer.nodes[o].weights = append(network.outputLayer.nodes[o].weights, rand.Float64())
		}
	}
}

//     def inspect(self):
//         print('------')
//         print('* Inputs: {}'.format(self.num_inputs))
//         print('------')
//         print('Hidden Layer')
//         self.hidden_layer.inspect()
//         print('------')
//         print('* Output Layer')
//         self.output_layer.inspect()
//         print('------')

//     def feed_forward(self, inputs):
//         hidden_layer_outputs = self.hidden_layer.feed_forward(inputs)
//         return self.output_layer.feed_forward(hidden_layer_outputs)

func (network *ThreeLayerNetwork) feedForward(inputs []float64) []float64 {
	hiddenLayerOutputs := network.hiddenLayer.feedForward(inputs)
	return network.outputLayer.feedForward(hiddenLayerOutputs)
}

//     # Uses online learning, ie updating the weights after each training case
//     def train(self, training_inputs, training_outputs):

func (network *ThreeLayerNetwork) Train(trainingInputs []float64, trainingOutputs []float64) {

	//         self.feed_forward(training_inputs)

	network.feedForward(trainingInputs)

	//         # 1. Output neuron deltas
	//         pd_errors_wrt_output_neuron_total_net_input = [0] * len(self.output_layer.neurons)
	//         for o in range(len(self.output_layer.neurons)):

	//             # ∂E/∂zⱼ
	//             pd_errors_wrt_output_neuron_total_net_input[o] = self.output_layer.neurons[o].calculate_pd_error_wrt_total_net_input(training_outputs[o])

	pdErrorsWrtOutputNeuronTotalNetInput := []float64{}
	for o := 0; o < len(network.outputLayer.nodes); o++ {
		pdErrorsWrtOutputNeuronTotalNetInput = append(pdErrorsWrtOutputNeuronTotalNetInput, network.outputLayer.nodes[o].calculatePdErrorWrtTotalNetInput(trainingOutputs[o]))
	}

	//         # 2. Hidden neuron deltas
	//         pd_errors_wrt_hidden_neuron_total_net_input = [0] * len(self.hidden_layer.neurons)
	//         for h in range(len(self.hidden_layer.neurons)):

	//             # We need to calculate the derivative of the error with respect to the output of each hidden layer neuron
	//             # dE/dyⱼ = Σ ∂E/∂zⱼ * ∂z/∂yⱼ = Σ ∂E/∂zⱼ * wᵢⱼ
	//             d_error_wrt_hidden_neuron_output = 0
	//             for o in range(len(self.output_layer.neurons)):
	//                 d_error_wrt_hidden_neuron_output += pd_errors_wrt_output_neuron_total_net_input[o] * self.output_layer.neurons[o].weights[h]

	//             # ∂E/∂zⱼ = dE/dyⱼ * ∂zⱼ/∂
	//             pd_errors_wrt_hidden_neuron_total_net_input[h] = d_error_wrt_hidden_neuron_output * self.hidden_layer.neurons[h].calculate_pd_total_net_input_wrt_input()

	pdErrorsWrtHiddenNeuronTotalNetInput := []float64{}
	for h := 0; h < len(network.hiddenLayer.nodes); h++ {
		dErrorWrtHiddenNeuronOutput := 0.0
		for o := 0; o < len(network.outputLayer.nodes); o++ {
			dErrorWrtHiddenNeuronOutput += pdErrorsWrtOutputNeuronTotalNetInput[o] * network.outputLayer.nodes[o].weights[h]
		}
		pdErrorsWrtHiddenNeuronTotalNetInput = append(pdErrorsWrtHiddenNeuronTotalNetInput, dErrorWrtHiddenNeuronOutput*network.hiddenLayer.nodes[h].calculatePdTotalNetInputWrtInput())
	}

	//         # 3. Update output neuron weights
	//         for o in range(len(self.output_layer.neurons)):
	//             for w_ho in range(len(self.output_layer.neurons[o].weights)):

	//                 # ∂Eⱼ/∂wᵢⱼ = ∂E/∂zⱼ * ∂zⱼ/∂wᵢⱼ
	//                 pd_error_wrt_weight = pd_errors_wrt_output_neuron_total_net_input[o] * self.output_layer.neurons[o].calculate_pd_total_net_input_wrt_weight(w_ho)

	//                 # Δw = α * ∂Eⱼ/∂wᵢ
	//                 self.output_layer.neurons[o].weights[w_ho] -= self.LEARNING_RATE * pd_error_wrt_weight

	for o := 0; o < len(network.outputLayer.nodes); o++ {
		for w_ho := 0; w_ho < len(network.outputLayer.nodes[o].weights); w_ho++ {
			pdErrorWrtWeight := pdErrorsWrtOutputNeuronTotalNetInput[o] * network.outputLayer.nodes[o].calculatePdTotalNetInputWrtWeight(w_ho)
			network.outputLayer.nodes[o].weights[w_ho] -= learningRate * pdErrorWrtWeight
		}
	}

	//         # 4. Update hidden neuron weights
	//         for h in range(len(self.hidden_layer.neurons)):
	//             for w_ih in range(len(self.hidden_layer.neurons[h].weights)):

	//                 # ∂Eⱼ/∂wᵢ = ∂E/∂zⱼ * ∂zⱼ/∂wᵢ
	//                 pd_error_wrt_weight = pd_errors_wrt_hidden_neuron_total_net_input[h] * self.hidden_layer.neurons[h].calculate_pd_total_net_input_wrt_weight(w_ih)

	//                 # Δw = α * ∂Eⱼ/∂wᵢ
	//                 self.hidden_layer.neurons[h].weights[w_ih] -= self.LEARNING_RATE * pd_error_wrt_weight

	for h := 0; h < len(network.hiddenLayer.nodes); h++ {
		for w_ih := 0; w_ih < len(network.hiddenLayer.nodes[h].weights); w_ih++ {
			pdErrorWrtWeight := pdErrorsWrtHiddenNeuronTotalNetInput[h] * network.hiddenLayer.nodes[h].calculatePdTotalNetInputWrtWeight(w_ih)
			network.hiddenLayer.nodes[h].weights[w_ih] -= learningRate * pdErrorWrtWeight
		}
	}

}

//     def calculate_total_error(self, training_sets):
//         total_error = 0
//         for t in range(len(training_sets)):
//             training_inputs, training_outputs = training_sets[t]
//             self.feed_forward(training_inputs)
//             for o in range(len(training_outputs)):
//                 total_error += self.output_layer.neurons[o].calculate_error(training_outputs[o])
//         return total_error

func (network *ThreeLayerNetwork) CalculateTotalError(trainingSets [][][]float64) float64 {
	totalError := 0.0
	for t := 0; t < len(trainingSets); t++ {
		trainingInputs := trainingSets[t][0]
		trainingOutputs := trainingSets[t][1]
		network.feedForward(trainingInputs)
		for o := 0; o < len(trainingOutputs); o++ {
			totalError += network.outputLayer.nodes[o].calculateError(trainingOutputs[o])
		}
	}
	return totalError
}

func (network *ThreeLayerNetwork) GetHiddenLayerWeights(nodeIndex int) []float64 {
	weights := []float64{}

	for i := 0; i < len(network.hiddenLayer.nodes); i++ {
		weights = append(weights, network.hiddenLayer.nodes[i].weights[nodeIndex])
	}

	return weights
}
