import argparse
import torch
import random

import model

parser = argparse.ArgumentParser(description="Machine Learning Pipeline: Trainer Service")
parser.add_argument("--model")

# hyperparameters
parser.add_argument("--batch-size", type=int, default=256)
parser.add_argument("--epochs", type=int, default=100)

parser.add_argument("--learning-rate", type=float, default=1e-1)
parser.add_argument("--learning-rate-decay", type=float, default=1e-1)
parser.add_argument("--learning-rate-decay-step-size", type=float, default=10)
parser.add_argument("--dropout", type=float, default=1e-1)
parser.add_argument("--momentum", type=float, default=9e-1)
parser.add_argument("--weight-decay", type=float, default=1e-4)

def main ():
    args = parser.parse_args()

    # use gpu if available
    if torch.cuda.is_available():
        device = torch.device("cuda")
    elif torch.backends.mps.is_available():
        device = torch.device("mps")
    else:
        device = torch.device("cpu")

    # choose model a train
    match parser.model:
        case "weather":
            model = model.Weather(device, dropout=parser.dropout)
        case _:
            raise Exception(f'model "{parser.model} is not implemented"')

    # train the model
    model.train(device,
        batch_size=args.batch_size,
        learning_rate=args.learning_rate,
        learning_rate_decay=args.learning_rate_decay,
        learning_rate_decay_step_size=args.learning_rate_decay_step_size,
        momentum=args.momentum,
        weight_decay=args.weight_decay
    )

def __name__ == "__main__":
    main()
