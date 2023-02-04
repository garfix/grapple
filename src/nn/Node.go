package nn

import "math"

// import "math"

type Node struct {
	bias    float64
	inputs  []float64
	output  float64
	weights []float64
}

// def __init__(self, bias):
//         self.bias = bias
//         self.weights = []

func createNode(bias float64) *Node {
	return &Node{
		bias:    bias,
		weights: []float64{},
	}
}

//     def calculate_output(self, inputs):
//         self.inputs = inputs
//         self.output = self.squash(self.calculate_total_net_input())
//         return self.output

func (node *Node) calculateOutput(inputs []float64) float64 {
	node.inputs = inputs
	node.output = node.squash(node.calculateTotalNetInput())
	return node.output
}

//     def calculate_total_net_input(self):
//         total = 0
//         for i in range(len(self.inputs)):
//             total += self.inputs[i] * self.weights[i]
//         return total + self.bias

func (node *Node) calculateTotalNetInput() float64 {
	total := 0.0
	for i := 0; i < len(node.inputs); i++ {
		total += node.inputs[i] * node.weights[i]
	}
	return total + node.bias
}

//     # Apply the logistic function to squash the output of the neuron
//     # The result is sometimes referred to as 'net' [2] or 'net' [1]
//     def squash(self, total_net_input):
//         return 1 / (1 + math.exp(-total_net_input))

func (node *Node) squash(totalNetInput float64) float64 {
	return 1.0 / (1.0 + math.Exp(-totalNetInput))
}

//     # Determine how much the neuron's total input has to change to move closer to the expected output
//     #
//     # Now that we have the partial derivative of the error with respect to the output (∂E/∂yⱼ) and
//     # the derivative of the output with respect to the total net input (dyⱼ/dzⱼ) we can calculate
//     # the partial derivative of the error with respect to the total net input.
//     # This value is also known as the delta (δ) [1]
//     # δ = ∂E/∂zⱼ = ∂E/∂yⱼ * dyⱼ/dzⱼ
//     #
//     def calculate_pd_error_wrt_total_net_input(self, target_output):
//         return self.calculate_pd_error_wrt_output(target_output) * self.calculate_pd_total_net_input_wrt_input();

func (node *Node) calculatePdErrorWrtTotalNetInput(targetOutput float64) float64 {
	return node.calculatePdErrorWrtOutput(targetOutput) * node.calculatePdTotalNetInputWrtInput()
}

//     # The error for each neuron is calculated by the Mean Square Error method:
//     def calculate_error(self, target_output):
//         return 0.5 * (target_output - self.output) ** 2

func (node *Node) calculateError(targetOutput float64) float64 {
	diff := targetOutput - node.output
	return 0.5 * diff * diff
}

//     # The partial derivate of the error with respect to actual output then is calculated by:
//     # = 2 * 0.5 * (target output - actual output) ^ (2 - 1) * -1
//     # = -(target output - actual output)
//     #
//     # The Wikipedia article on backpropagation [1] simplifies to the following, but most other learning material does not [2]
//     # = actual output - target output
//     #
//     # Alternative, you can use (target - output), but then need to add it during backpropagation [3]
//     #
//     # Note that the actual output of the output neuron is often written as yⱼ and target output as tⱼ so:
//     # = ∂E/∂yⱼ = -(tⱼ - yⱼ)
//     def calculate_pd_error_wrt_output(self, target_output):
//         return -(target_output - self.output)

func (node *Node) calculatePdErrorWrtOutput(targetOutput float64) float64 {
	return -(targetOutput - node.output)
}

//     # The total net input into the neuron is squashed using logistic function to calculate the neuron's output:
//     # yⱼ = φ = 1 / (1 + e^(-zⱼ))
//     # Note that where ⱼ represents the output of the neurons in whatever layer we're looking at and ᵢ represents the layer below it
//     #
//     # The derivative (not partial derivative since there is only one variable) of the output then is:
//     # dyⱼ/dzⱼ = yⱼ * (1 - yⱼ)
//     def calculate_pd_total_net_input_wrt_input(self):
//         return self.output * (1 - self.output)

func (node *Node) calculatePdTotalNetInputWrtInput() float64 {
	return node.output * (1.0 - node.output)
}

//     # The total net input is the weighted sum of all the inputs to the neuron and their respective weights:
//     # = zⱼ = netⱼ = x₁w₁ + x₂w₂ ...
//     #
//     # The partial derivative of the total net input with respective to a given weight (with everything else held constant) then is:
//     # = ∂zⱼ/∂wᵢ = some constant + 1 * xᵢw₁^(1-0) + some constant ... = xᵢ
//     def calculate_pd_total_net_input_wrt_weight(self, index):
//         return self.inputs[index]

func (node *Node) calculatePdTotalNetInputWrtWeight(index int) float64 {
	return node.inputs[index]
}
