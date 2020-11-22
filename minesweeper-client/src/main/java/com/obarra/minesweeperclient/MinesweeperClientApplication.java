package com.obarra.minesweeperclient;

import org.springframework.boot.SpringApplication;
import org.springframework.boot.autoconfigure.SpringBootApplication;
import org.springframework.cloud.openfeign.EnableFeignClients;

@SpringBootApplication
@EnableFeignClients
public class MinesweeperClientApplication {

	public static void main(String[] args) {
		SpringApplication.run(MinesweeperClientApplication.class, args);
	}

}
