package com.jogos.snake;

import java.awt.Rectangle;
import java.awt.Color;
import java.awt.Font;
import java.awt.FontMetrics;
import java.awt.Graphics2D;
import java.awt.event.ActionEvent;
import java.awt.event.ActionListener;

import javax.swing.JPanel;
import javax.swing.Timer;

//ACTIONLISTENER - serve para anima��es
public class Graphics extends JPanel implements ActionListener{
	
	private Timer t = new Timer(100, this);
	public String state;
	
	private Snake s;
	private Food f;
	private Game game;
	
	
	public Graphics(Game g) {
		t.start();
		state = "START";
		
		game = g;
		s = g.getPlayer();
		f =g.getFood();
		
		//add a keyListener
		this.addKeyListener(g);
		this.setFocusable(true);
		this.setFocusTraversalKeysEnabled(false);
	}
	
	public void paintComponent(java.awt.Graphics g) {
		super.paintComponent(g);
	
		Graphics2D g2d = (Graphics2D) g;
		
		g2d.setColor(Color.black);
		g2d.fillRect(0, 0, Game.width * Game.dimension + 5, Game.height * Game.dimension + 5);
		
		if(state == "START") {
			g2d.setColor(Color.white);
			g2d.drawString("Press Any Key", Game.width/2 * Game.dimension - 40, Game.height / 2 * Game.dimension - 20);
		}
		else if (state == "RUNNING") {
			
			g2d.setColor(Color.red);
			g2d.fillRect(f.getX() * Game.dimension, f.getY() * Game.dimension, Game.dimension, Game.dimension);
		
			g2d.setColor(Color.green); 
			for(Rectangle r : s.getBody()) {
				g2d.fill(r);
			}
		}
		else {
			g2d.setColor(Color.white);
			g2d.setFont(new Font("Consolas", Font.BOLD, 16));
			g2d.drawString("GAME OVER", Game.width/2 * Game.dimension - 40, Game.height / 2 * Game.dimension - 35);
			g2d.drawString("Your Score: " + (s.getBody().size() -3), Game.width/2 * Game.dimension - 55, Game.height / 2 * Game.dimension - 20);
		}	
	}
	
	
	@Override
	public void actionPerformed(ActionEvent e) {
		//repaint recalls paint component
		repaint();
		game.update();
	}

	
}
