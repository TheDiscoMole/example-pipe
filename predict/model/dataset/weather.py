import torch

import numpy as np
import pandas as pd

from google.cloud import bigquery

class Weather (torch.utils.data.Dataset):
    def __init__ (self):
        # connect to bigquery
        client = bigquery.Client()

        # load dataset
        data = client.query("SELECT * from weather")
        data = data.to_dataframe()

        # store dataset and tensor column selection
        self.data = data
        self.tensor_columns = [
            "latitude_sin", "latitude_cos",
            "longitude_sin", "longitude_cos",
            "week_of_the_year_sin", "week_of_the_year_cos",
            "day_of_the_week_sin", "day_of_the_week_cos",
            "temperature", "temperature_feels_like",
            "pressure", "humidity", "cloudiness", "visibility",
            "precipitation_probability", "rain_volume", "snow_volume",
            "wind_speed", "wind_angle", "wind_speed"
        ]

    # override len(data)
    def __len__ (self):
        return len(self.data)

    # override data[idx]
    def __getitem__ (self, idx):

        location = torch.tensor(self.data.iloc[idx][self.tensor_columns[:8]].values)
        history = self.locationHistory(forecast["latitude"], forecast["longitude"], 100, forecast["time"], forecast["time"] - 60*60*24*7)
        forecast = torch.tensor(self.data.iloc[idx][self.tensor_columns[8:]].values)

        return location, history, forecast

    # weather history for a location + radius
    def locationHistory (self, latitude, longitude, radius, time_start, time_stop):
        radius_earth = 6378

        latitude_offset = (radius / radius_earth) * (180 / np.pi)
        longitude_offset = (radius / radius_earth) * (180 / np.pi) / np.cos(latitude * np.pi / 180)

        latitude_filter = self.data["latitude"].between(latitude-latitude_offset, latitude+latitude_offset)
        longitude_filter = self.data["longitude"].between(longitude-longitude_offset, longitude+longitude_offset)

        history = self.data[latitude_filter and longitude_filter and self.data["time"].between(time_start, time_stop)]
        history = history[self.tensor_columns]

        return torch.Tensor(history.values)
