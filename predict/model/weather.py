import torch

import dataset
import modules
import repository

class Weather (torch.nn.Module):
    def __init__ (self,
        dropout=1e-1,
    ):
        self.dropout = torch.nn.Dropout(dropout)

        self.linear1 = torch.nn.Linear(18, 512)
        self.encoder = modules.AttentionEncoder(512)
        self.linear2 = torch.nn.Linear(8, 512)
        self.decoder = modules.AttentionEncoder(512)
        self.linear3 = torch.nn.Linear(512, 10)

    def forward (self, location, history):

        encode = self.linear1(history)
        encode = self.dropout(encode)
        encode = self.encoder(encode, encode, encode)

        decode = self.linear2(location)
        decode = self.dropout(decode)
        decode = self.decoder(encode, encode, decode)

        return self.linear3(decode)


    def train (self, device,
        batch_size=args.batch_size,
        learning_rate=args.learning_rate,
        learning_rate_decay=args.learning_rate_decay,
        learning_rate_decay_step_size=args.learning_rate_decay_step_size,
        momentum=args.momentum,
        weight_decay=args.weight_decay
    ):
        self.to(device)
        self.training = True
        for module in self.children(): module.train(self.training)

        # set up dataset
        dataset = dataset.Weather()

        training_set, validation_set = torch.utils.data.random_split(dataset, [len(dataset) - len(dataset) / 9, len(dataset) / 9])
        training_loader, validation_loader = torch.utils.data.DataLoader(training_set, batch_size=batch_size, shuffle=True), torch.utils.data.DataLoader(validation_set, batch_size=batch_size)

        # prepare optimizer
        criterion = torch.nn.MSELoss().to(device)
        optimizer = torch.optim.SGD(model.parameters(), learning_rate, momentum=momentum, weight_decay=weight_decay)
        scheduler = torch.optim.lr_scheduler.StepLR(optimizer, step_size=learning_rate_decay_step_size, gamma=learning_rate_decay)

        # remember losses
        losses = []
        best = 1.

        # run training
        for epoch in range(epochs):
            for batch, (location, history, target) in enumerate(training_loader):
                location = location.view(-1, 1, 8)
                target = target.view(-1, 1, 10)

                location = location.to(device=device, dtype=torch.bfloat16, non_blocking=True)
                history = history.to(device=device, dtype=torch.bfloat16, non_blocking=True)
                target = target.to(device=device, dtype=torch.bfloat16, non_blocking=True)

                output = model(location, history)
                loss = criterion(output, target)

                optimizer.zero_grad()
                loss.backward()
                optimizer.step()

            losses.append(0.)
            for batch, (location, history, target) in enumerate(validation_loader):
                with torch.no_grad():
                    location = location.to(device=device, dtype=torch.bfloat16, non_blocking=True)
                    history = history.to(device=device, dtype=torch.bfloat16, non_blocking=True)
                    target = target.to(device=device, dtype=torch.bfloat16, non_blocking=True)

                    output = model(location, history)
                    losses[-1] += criterion(output, target)

            print(f'epoch: {epoch} loss: {losses[-1]}')

            # save if best checkpoint
            if losses[-1] < best:
                repository.model.save({
                    "epoch": epoch,
                    "model_state_dict": model.model_state_dict(),
                    "loss": loss
                }, "model/weather.pt")
                best = losses[-1]

            scheduler.step()
